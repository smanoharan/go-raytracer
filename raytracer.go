// raytracer.go: Contains ray tracing methods.

package main

import (
	"math/rand"
	"image"
)

// the options controlling the behaviour of the RayTracer
type RayTracerOptions struct {
	maxDepth, samplingFactor, numShadowRays int
}

// The main 'class' which performs the ray tracing
type RayTracer struct {
	width, height                     int
	halfWidth, halfHeight, tanX, tanY entry
	basisU, basisV, basisW, eyePos    Vec3
	options                           *RayTracerOptions
}

// create a new ray tracer using the given view-window and options
func NewRayTracer(view *Camera, options *RayTracerOptions) *RayTracer {

	// compute half-dimensions and angles:
	halfWidth, halfHeight := entry(view.width)/TWO, entry(view.height)/TWO
	tanY := tan(view.fovY / TWO)
	tanX := tanY * (halfWidth / halfHeight)

	// compute eye-basis vectors:
	bW := view.pos.minus(&view.lookAt).direction()
	bU := view.up.cross(bW).direction()
	bV := bW.cross(bU)

	return &RayTracer{
		view.width, view.height,
		halfWidth, halfHeight, tanX, tanY,
		*bU, *bV, *bW, view.pos,
		options,
	}
}

// find the closest shape which intersects the ray
func findClosestIntersection(ray *Ray, scene []*Shape) (hit bool, inter *Intersection, closest *Shape) {

	// by default: no intersection
	hit, inter, closest = false, nil, nil

	// iterate through each shape:
	for _, shape := range scene {
		// check if this shape hits the ray at a closer point than previous least.
		if h, i := (*shape).Intersect(ray); h && (!hit || i.dist < inter.dist) {
			hit, inter, closest = true, i, shape
		}
	}

	return
}

// build the ray travelling from the eye to the i,j point in the image
func (r *RayTracer) buildRayFromEyeToImage(i, j entry, eye *Vec3) *Ray {
	// formulas from reference calculations
	alpha := r.tanX * ((j / r.halfWidth) - ONE)
	beta := r.tanY * (ONE - (i / r.halfHeight))
	dir := r.basisU.scale(alpha).plus(r.basisV.scale(beta)).minus(&r.basisW).direction()
	return &Ray{*eye, *dir}
}

// scale a color in [0,1] to the [0,255] range.
func (e entry) toRGB() uint8 {
	if e >= ONE {
		return uint8(255)
	}
	if e <= ZERO {
		return uint8(0)
	}
	return uint8(e * entry(256))
}

// reflect a ray about normal
func reflect(dir, normal *Vec3) *Vec3 {
	return dir.minus(normal.scale(TWO * normal.dot(dir)))
}

// generates a small random number in range (-alpha, alpha) where alpha = 0.001
func smallRand(sc float64) entry {
	return entry(rand.Float64()-0.5) * entry(sc)
}

func randVec() *Vec3 {
	return &Vec3{smallRand(0.25), smallRand(0.25), smallRand(0.25)}
}

// Compute the color of the current ray by tracing it into the scene
func (r *RayTracer) findColor(ray *Ray, scene []*Shape, lights []*Light, curDepth int) *Vec3 {

	// check if the ray hits any objects:
	if hit, inter, closest := findClosestIntersection(ray, scene); hit {

		// apply material of the closest shape
		material := (*closest).GetMaterial()
		color := material.ambient.plus(&material.emission)

		// apply each light that is visible from the intersection point
		for _, light := range lights {

			// find offset to light
			lightOffset := (*light).OffsetFrom(&inter.point)
			distToLight := lightOffset.magnitude()

			// enable soft-shadowing by tracing multiple shadow rays
			numRays := r.options.numShadowRays
			rayWeight := ONE / entry(numRays)
			for j := 0; j < numRays; j++ {
				shadowRayDir := lightOffset.plus(randVec()).direction()
				shadowRay := &Ray{
					*inter.point.plus(shadowRayDir.scale(entry(0.001))), // push ray towards light
					*shadowRayDir,
				}

				// check if shadowRay hits any objects in the scene:
				if h, i, _ := findClosestIntersection(shadowRay, scene); (!h) || i.dist >= distToLight {
					extraColor := BlinnPhongShader(light, shadowRayDir, &i.normal, ray, material, distToLight)
					color = color.plus(extraColor.scale(rayWeight))
				}
			}
		}

		// recursively trace rays
		if curDepth < r.options.maxDepth {
			refRayDir := reflect(&ray.direction, &inter.normal)

			// if this is the primary ray, perform some blurring
			numRays := 1
			if curDepth == 0 {
				numRays = 4 // TODO move into options
			}
			refRayWeight := ONE / entry(numRays)

			for i := 0; i < numRays; i++ {

				// build reflected ray
				refRay := &Ray{
					*inter.point.plus(refRayDir.scale(entry(0.001))), // to avoid self-collision
					*refRayDir.plus(inter.normal.scale(smallRand(0.001))).direction(),
				}

				// trace the reflected ray // TODO early stop if extraColor is small
				extraColor := material.specular.scale(refRayWeight).times(r.findColor(refRay, scene, lights, curDepth+1))
				color = color.plus(extraColor)
			}
		}
		return color
	}

	// no intersections:
	return &ZERO_V3
}

func (r *RayTracer) Draw(scene []*Shape, lights []*Light) *image.RGBA {
	
	img := NewOutputImage(r.width,r.height)

	// iterate through the image
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			cx, cy := entry(x), entry(y)
			ray := r.buildRayFromEyeToImage(cy, cx, &r.eyePos)
			Set(img, x, y, r.findColor(ray, scene, lights, 0))
		}
	}

	return img
}

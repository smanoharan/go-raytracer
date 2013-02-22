//shape.go: Contains definitions of 3D primitives of the objects in the scene.

package main

// Intersection holds the results of an intersection test
type Intersection struct {
	point, normal Vec3  // Point of intersection and the normal
	dist          entry // distance from ray-origin to intersection point
}

// A Shape is a primitive in 3D space. 
type Shape interface {
	GetMaterial() *Material
	Intersect(ray *Ray) (bool, *Intersection)
}

// Sphere implementation of Shape
type Sphere struct {
	// transform, inverse(trans), transpose(transInv):
	trans, transInv, transInvTr Mat4
	mat                         *Material
}

func rotate(axis *Vec3, angle entry) *Mat3 {
	x, y, z := axis[cX], axis[cY], axis[cZ]

	// using Rodrequiz formula
	cosT, sinT := cos(angle), sin(angle)
	p1 := &Mat3{x * x, x * y, x * z, x * y, y * y, y * z, x * z, y * z, z * z}
	p2 := &Mat3{0, -z, y, z, 0, -x, -y, x, 0}
	return IDENTITY_M3.scale(cosT).plus(p1.scale(ONE - cosT)).plus(p2.scale(sinT))
}

func transform(scale, pos, rotAxis *Vec3, angle entry) *Mat4 {
	sx, sy, sz := scale[cX], scale[cY], scale[cZ]
	rot := rotate(rotAxis, angle)
	return &Mat4{
		rot[0] * sx, rot[1] * sx, rot[2] * sx, pos[cX],
		rot[3] * sy, rot[4] * sy, rot[5] * sy, pos[cY],
		rot[6] * sz, rot[7] * sz, rot[8] * sz, pos[cZ],
		0, 0, 0, 1,
	}
}

// NewSphere creates a Sphere at a given point, with a given radius and material
func NewSphere(radius entry, center *Vec3, mat *Material) *Sphere {
	radiusVec := &Vec3{radius, radius, radius}
	return NewEllipsoid(radiusVec, center, mat)
}

// NewEllipsoid creates a Ellipsoid (scaled sphere) at a point
func NewEllipsoid(radius, center *Vec3, mat *Material) *Sphere {
	return NewRotatedEllipsoid(radius, center, &X_V3, ZERO, mat)
}

// NewRotatedEllipsoid creates an Ellipsoid with a rotation applied
// Parameters: radius-{x,y,z} ; center-{x,y,z} ; rotation-axis-{x,y,z}, rotation-angle (degrees)
func NewRotatedEllipsoid(radius, center, rot *Vec3, angle entry, mat *Material) *Sphere {
	trans := transform(radius, center, rot, angle)
	transInv := trans.inverse()
	transInvTr := transInv.transpose()
	return &Sphere{*trans, *transInv, *transInvTr, mat}
}

// GetMaterial returns the material of the surface of the sphere.
func (s *Sphere) GetMaterial() *Material {
	return s.mat
}

func toV4(v3 *Vec3, w entry) *Vec4 {
	return &Vec4{v3[cX], v3[cY], v3[cZ], w}
}

func toV3(v4 *Vec4) *Vec3 {
	return &Vec3{v4[cX], v4[cY], v4[cZ]}
}

// Intersect checks if the ray intersects the sphere.
func (s *Sphere) Intersect(ray *Ray) (hit bool, res *Intersection) {
	// by default: no intersection:
	hit, res = false, nil

	// transform ray by the sphere's inverse transform,
	// which allows comparison against a unit sphere.
	invStart := s.transInv.timesVec(toV4(&(ray.start), ONE))
	invStart[cW] = ZERO // correcting for translation.
	invDir := s.transInv.timesVec(toV4(&(ray.direction), ZERO))

	// ray-sphere intersection:
	// a quadratic ax^2 + bx + c = 0
	a := invDir.dot(invDir)
	b := TWO * invDir.dot(invStart)
	c := invStart.dot(invStart) - ONE // since s.radius^2 is ONE

	// check det >= 0
	if det := b*b - FOUR*a*c; det >= 0 {

		// compute roots:
		sqrtDet, twoA := sqrt(det), TWO*a
		x1 := (-b + sqrtDet) / twoA
		x2 := (-b - sqrtDet) / twoA

		// check that at least one root is positive:
		if (x1 >= 0) || (x2 >= 0) {

			// if both roots are positive, pick least root (i.e. closest intersection)
			if (x1 < 0) || (x2 > 0 && x2 < x1) {
				x1 = x2
			}

			// compute point of intersection (in transformed space)
			invPt := invStart.plus(invDir.scale(x1))

			invPt[cW] = ZERO // correcting for translation.
			normal := toV3(s.transInvTr.timesVec(invPt)).direction()

			invPt[cW] = ONE                     // correcting for translation.
			pt := toV3(s.trans.timesVec(invPt)) // convert back into normal co-ords
			dist := pt.distanceTo(&(ray.start))

			hit, res = true, &Intersection{*pt, *normal, dist}
		}

	}
	return
}

// TODO: Concrete implementations of shapes: Triangle and Quad

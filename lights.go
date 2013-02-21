// lights.go: Contains implementation of lighting systems.

package main

// A Material stores lighting properties of an object.
type Material struct {
	ambient, emission, diffuse, specular Vec3
	shininess                            entry
}

// A Light is a source of light in the scene.
type Light interface {
	OffsetFrom(point *Vec3) *Vec3   // get un-normalized direction to light from point
	AttenuationAt(dist entry) entry // get attentuation factor at the distance
	GetColor() *Vec3                // get color of the light
}

// A Shader determines the color of a point in the scene using the Lights and Materials.
type Shader func(light *Light, lightDir, normal *Vec3, ray *Ray, mat *Material, dist entry) *Vec3

// Implementations of Lights: Point and Directional
type PointLight struct {
	// color, light-position, attenuation coeff
	color, position, atten Vec3
}

func (p *PointLight) GetColor() *Vec3 {
	return &(p.color)
}

func (p *PointLight) OffsetFrom(point *Vec3) *Vec3 {
	return p.position.minus(point)
}

func (p *PointLight) AttenuationAt(dist entry) entry {
	// attenuation for point lights is dependent of the distance (d)
	// attenuation factor is (a + b*d + c*d^2) where a,b,c are the attenuation coeff's.
	return p.atten[cX] + p.atten[cY]*dist + p.atten[cZ]*dist*dist
}

type DirectionalLight struct {
	color, direction Vec3
}

func (d *DirectionalLight) GetColor() *Vec3 {
	return &(d.color)
}

func (d *DirectionalLight) OffsetFrom(point *Vec3) *Vec3 {
	return &(d.direction) // direction is const for these lights
}

func (d *DirectionalLight) AttenuationAt(dist entry) entry {
	return ONE // no attenuation for directional lights.
}

// An implementation of a Shader: Blinn-Phong Lighting model
func BlinnPhongShader(lightPtr *Light, lightDir, normal *Vec3, ray *Ray, mat *Material, dist entry) *Vec3 {
	// resolve light pointer:
	light := *lightPtr

	// compute the halfway vector between eye direction and light direction:
	halfVec := lightDir.minus(&(ray.direction)).direction()

	// diffuse factor is normal dot light direction (if > 0)
	diffuseColor := &ZERO_V3
	if diffuse := normal.dot(lightDir); diffuse > 0 {
		diffuseColor = mat.diffuse.scale(diffuse)
	}

	// specular factor is (normal dot halfvec)^shininess (if > 0)
	specularColor := &ZERO_V3
	if specular := normal.dot(halfVec); specular > 0 {
		specularColor = mat.specular.scale(specular.pow(mat.shininess))
	}

	// scale down by attenuation
	return light.GetColor().scale(ONE / light.AttenuationAt(dist)).times(diffuseColor.plus(specularColor))
}

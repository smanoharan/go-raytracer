// lights.go: Contains implementation of lighting systems.

package main

// A Material stores lighting properties of an object.
type Material struct {
	Ambient, Emission, Diffuse, Specular Vec3
	Shininess entry
}

// A Light is a source of light in the scene.
type Light interface {
	OffsetFrom(point *Vec3) *Vec3 // get un-normalized direction to light from point
	AttenuationAt(dist entry) entry // get attentuation factor at the distance
}

// A Shader determines the color of a point in the scene using the Lights and Materials.
type Shader func(light *Light, lightDir, normal *Vec3, ray *Ray, mat *Material, dist entry) *Vec3

// TODO: Concrete impl of Lights: Point and Directional
// TODO: Concrete impl of Shader: Blinn-Phong

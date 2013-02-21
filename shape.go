//shape.go: Contains definitions of 3D primitives of the objects in the scene.

package main

// Intersection holds the results of an intersection test
type Intersection struct {
	Point, Normal Vec3 // Point of intersection and the normal
	DistToRay entry // distance from ray-origin to intersection point
}

// A Shape is a primitive in 3D space. 
type Shape interface {
	GetMaterial() *Material
	Intersect(ray *Ray) (bool, *Intersection)
}

// TODO: Concrete implementations of shapes: Sphere, Triangle and Quad

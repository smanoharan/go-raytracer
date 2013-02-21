// contains tests for shape.go

package main

import (
	"fmt"
	"testing"
)

func TestIntersectionForUnitSphere(t *testing.T) {
	s := NewSphere(ONE, &Material{}, &IDENTITY_M4) // unit sphere, centered at origin
	msg := "Ray-Sphere intersection "

	// case 0: a ray from origin passing through x-axis:
	ray := &Ray{ZERO_V3, X_V3}
	exp := &Intersection{X_V3, X_V3, ONE}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"0")

	// case 1: a ray from outside the sphere passing through the sphere via y-axis:
	ray = &Ray{*Y_V3.scale(-FOUR), Y_V3}
	exp = &Intersection{*Y_V3.scale(-ONE), *Y_V3.scale(-ONE), entry(3)}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"1")

	// case 2: a ray just passing through the sphere at (0,1,0):
	ray = &Ray{Vec3{0, 1, -2}, Z_V3}
	exp = &Intersection{Y_V3, Y_V3, TWO}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"2")

	// case 3: a ray in dir (1,1,1) missing the sphere
	dir := &Vec3{1, 1, 1}
	ray = &Ray{Vec3{2, 1, 2}, *(dir.direction())}
	assertIntersectionEquals(t, s, ray, false, nil, msg+"3")

	// case 4: a ray in dir (-1,3,-5) hitting the sphere at (0, 0.6, 0.8):
	dir = &Vec3{-1, 3, -5}
	ray = &Ray{Vec3{1.7, -4.5, 9.3}, *(dir.direction())}
	hit := Vec3{0, 0.6, 0.8}
	exp = &Intersection{hit, hit, entry(1.7) * sqrt(entry(35))}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"4")
}

// Same as TestIntersectionForUnitSphere, except the sphere is scaled by (0.5, 1.25, 2.5)
func TestIntersectionForScaledSphere(t *testing.T) {
	sc := &Vec3{0.5, 1.25, 2.5}
	msg := "Scaled Ray-Sphere intersection "

	// unit sphere, centered at origin
	mat1 := &Mat4{0.5, 0, 0, 0, 0, 1.25, 0, 0, 0, 0, 2.5, 0, 0, 0, 0, 1}
	s1 := NewSphere(ONE, &Material{}, mat1)

	// sphere of radius 0.5, centered at origin
	mat2 := &Mat4{1, 0, 0, 0, 0, 2.5, 0, 0, 0, 0, 5, 0, 0, 0, 0, 1}
	s2 := NewSphere(entry(0.5), &Material{}, mat2)

	// case 0: a ray from origin passing through x-axis:
	ray := &Ray{ZERO_V3, X_V3}
	exp := &Intersection{*X_V3.scale(sc[cX]), X_V3, ONE * sc[cX]}
	assertIntersectionEquals(t, s1, ray, true, exp, msg+"0.1")
	assertIntersectionEquals(t, s2, ray, true, exp, msg+"0.2")

	// case 1: a ray from outside the sphere passing through the sphere via y-axis:
	ray = &Ray{*Y_V3.scale(-FOUR), Y_V3}
	exp = &Intersection{*Y_V3.scale(-ONE).scale(sc[cY]), *Y_V3.scale(-ONE), entry(2.75)}
	assertIntersectionEquals(t, s1, ray, true, exp, msg+"1.1")
	assertIntersectionEquals(t, s2, ray, true, exp, msg+"1.2")

	// case 2: a ray just passing through the sphere at (0,1,0):
	ray = &Ray{Vec3{0, 1.25, -4}, Z_V3}
	exp = &Intersection{*Y_V3.scale(sc[cY]), Y_V3, FOUR}
	assertIntersectionEquals(t, s1, ray, true, exp, msg+"2.1")
	assertIntersectionEquals(t, s2, ray, true, exp, msg+"2.2")

	// case 3: a ray in dir (1,1,1) missing the sphere
	dir := &Vec3{1, 1, 1}
	ray = &Ray{Vec3{2, 1, 2}, *(dir.direction())}
	assertIntersectionEquals(t, s1, ray, false, nil, msg+"3.1")
	assertIntersectionEquals(t, s2, ray, false, nil, msg+"3.2")

	// case 4: skipped.
}

// same cases as TestIntersectionForUnitSphere, except for translation by (-10, 0.44, -2.5)
func TestIntersectionForTranslatedUnitSphere(t *testing.T) {
	tr := &Vec3{-10, 0.44, -2.5}
	trMat := Mat4{1, 0, 0, -10, 0, 1, 0, 0.44, 0, 0, 1, -2.5, 0, 0, 0, 1}
	s := NewSphere(ONE, &Material{}, &trMat) // unit sphere, centered at origin
	msg := "Translated Ray-Sphere intersection "

	// case 0: a ray from origin passing through x-axis:
	ray := &Ray{*ZERO_V3.plus(tr), X_V3}
	exp := &Intersection{*X_V3.plus(tr), X_V3, ONE}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"0")

	// case 1: a ray from outside the sphere passing through the sphere via y-axis:
	ray = &Ray{*Y_V3.scale(-FOUR).plus(tr), Y_V3}
	exp = &Intersection{*Y_V3.scale(-ONE).plus(tr), *Y_V3.scale(-ONE), entry(3)}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"1")

	// case 2: a ray just passing through the sphere at (0,1,0):
	src := &Vec3{0, 1, -2}
	ray = &Ray{*src.plus(tr), Z_V3}
	exp = &Intersection{*Y_V3.plus(tr), Y_V3, TWO}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"2")

	// case 3: a ray in dir (1,1,1) missing the sphere
	src = &Vec3{2, 1, 2}
	dir := &Vec3{1, 1, 1}
	ray = &Ray{*src.plus(tr), *(dir.direction())}
	assertIntersectionEquals(t, s, ray, false, nil, msg+"3")

	// case 4: a ray in dir (-1,3,-5) hitting the sphere at (0, 0.6, 0.8):
	src = &Vec3{1.7, -4.5, 9.3}
	dir = &Vec3{-1, 3, -5}
	ray = &Ray{*src.plus(tr), *(dir.direction())}
	hit := Vec3{0, 0.6, 0.8}
	exp = &Intersection{*hit.plus(tr), hit, entry(1.7) * sqrt(entry(35))}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"4")
}

// tests for a sphere which is translated to (-3,1,2) and scaled by (2,0.5,3)
func TestIntersectionForTranslatedScaledSphere(t *testing.T) {
	mat := &Mat4{0.5, 0, 0, -3, 0, 0.125, 0, 1, 0, 0, 0.75, 2, 0, 0, 0, 1}
	s := NewSphere(FOUR, &Material{}, mat)
	msg := "Translated Scaled Ray-Sphere intersection "

	// case 0: a ray which is missing the sphere
	ray := &Ray{ZERO_V3, *X_V3.scale(entry(3))}
	assertIntersectionEquals(t, s, ray, false, nil, msg+"0")

	// case 1: a ray which hits the sphere
	ray = &Ray{Vec3{-3, 1, -5}, Z_V3}
	exp := &Intersection{Vec3{-3, 1, -1}, *Z_V3.scale(-ONE), FOUR}
	assertIntersectionEquals(t, s, ray, true, exp, msg+"1")
}

func isIntersectionResultEqual(exp, act *Intersection) bool {
	return (!exp.dist.neq(act.dist)) &&
		isMatEqual(exp.point[:], act.point[:], V3LEN) &&
		isMatEqual(exp.normal[:], act.normal[:], V3LEN)
}

func assertIntersectionEquals(t *testing.T, sphere *Sphere, ray *Ray, expHit bool, expInter *Intersection, msg string) {
	hit, res := sphere.Intersect(ray)
	passed := assert(t, hit == expHit, msg+fmt.Sprint(": Expected Hit: ", expHit))
	if passed && expHit {
		assert(t, isIntersectionResultEqual(expInter, res), msg+fmt.Sprint(":\n\t\tExp: ", *expInter, "\n\t\tAct: ", *res))
	}
}

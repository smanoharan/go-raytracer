// Test file for matrix.go
package main

import (
	"fmt"
	"testing"
)

// Test Cases:
// 	- addition tests
//		- Vec3, Vec4
//		- Mat3, Mat4
//  - multiplication tests
//		- scalar (Vec3, Vec4, Mat3, Mat4)
//		- dot (Vec3, Vec4), cross (Vec3) 
//		- Mat3*Mat3, Mat4*Mat4 
//		- Vec3*Mat3, Mat3*Vec3, Vec4*Mat4, Vec4*Mat4 

var (
	// some example matrices and vectors, for testing
	v31 Vec3 = Vec3{4, 7, 112}
	v32 Vec3 = Vec3{1, 2, -11}
	v41 Vec4 = Vec4{-10, 5, 112, 18}
	v42 Vec4 = Vec4{144, 3, -10, 22}
	m31 Mat3 = Mat3{11, 12, 13, 21, 22, 23, 31, 32, 33}
	m32 Mat3 = Mat3{44, 45, 46, 33, 30, 27, 51, 42, 33}
	m41 Mat4 = Mat4{11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 33, 34, 41, 42, 43, 44}
	m42 Mat4 = Mat4{44, 45, 46, 47, 33, 30, 27, 24, 51, 42, 33, 24, 11, 21, 13, 41}
)

func TestAdditionOf3DVectors(t *testing.T) {
	expV := Vec3{5, 9, 101} // v31+v32

	// check both addition directions:
	assertEquals(t, expV, *v31.plus(&v32), "Vec3 addition 1+2")
	assertEquals(t, expV, *v32.plus(&v31), "Vec3 addition 2+1")
}

func TestAdditionOf4DVectors(t *testing.T) {
	expV := Vec4{134, 8, 102, 40} // v41+v42

	// check both addition directions:
	assertEquals(t, expV, *v41.plus(&v42), "Vec4 addition 1+2")
	assertEquals(t, expV, *v42.plus(&v41), "Vec4 addition 2+1")
}

func TestAdditionOf3x3Matrices(t *testing.T) {
	expM := Mat3{55, 57, 59, 54, 52, 50, 82, 74, 66} // m31+m32

	// check both addition directions
	assertEquals(t, expM, *m31.plus(&m32), "Mat3 addition 1+2")
	assertEquals(t, expM, *m32.plus(&m31), "Mat3 addition 2+1")
}

func TestAdditionOf4x4Matrices(t *testing.T) {
	expM := Mat4{55, 57, 59, 61, 54, 52, 50, 48, 82, 74, 66, 58, 52, 63, 56, 85} // m41+m42

	// check both addition directions
	assertEquals(t, expM, *m41.plus(&m42), "Mat4 addition 1+2")
	assertEquals(t, expM, *m42.plus(&m41), "Mat4 addition 2+1")
}

func TestScalarMultiplicationOf3DVectors(t *testing.T) {
	assertEquals(t, Vec3{8, 14, 224}, *v31.scale(2), "Vec3 mult by 2")
	assertEquals(t, Vec3{-1, -1.75, -28}, *v31.scale(-0.25), "Vec3 mult by -0.25")
}

func TestScalarMultiplicationOf4DVectors(t *testing.T) {
	assertEquals(t, Vec4{-30, 15, 336, 54}, *v41.scale(3), "Vec4 mult by 3")
	assertEquals(t, Vec4{5, -2.5, -56, -9}, *v41.scale(-0.5), "Vec4 mult by -0.5")
}

func TestScalarMultiplicationOf3x3Matrices(t *testing.T) {
	exp1m := Mat3{44, 48, 52, 84, 88, 92, 124, 128, 132}       // m31*4
	exp2m := Mat3{-11, -12, -13, -21, -22, -23, -31, -32, -33} // m31 * -1.0
	assertEquals(t, exp1m, *m31.scale(4), "Mat3 mult by 4")
	assertEquals(t, exp2m, *m31.scale(-1.0), "Mat3 mult by -1.0")
}

func TestScalarMultiplicationOf4x4Matrices(t *testing.T) {
	exp1m := Mat4{110, 120, 130, 140, 210, 220, 230, 240, 310, 320, 330, 340, 410, 420, 430, 440}      // m41*10
	exp2m := Mat4{2.75, 3, 3.25, 3.5, 5.25, 5.50, 5.75, 6, 7.75, 8, 8.25, 8.5, 10.25, 10.5, 10.75, 11} // m41*0.25
	assertEquals(t, exp1m, *m41.scale(10), "Mat3 mult by 10")
	assertEquals(t, exp2m, *m41.scale(0.25), "Mat3 mult by 0.25")
}

func TestDotProductOf3DVectors(t *testing.T) {
	expV := Vec3{4, 14, -1232} // v31 . v32

	// check both dot-prod directions
	assertEquals(t, expV, *v31.dot(&v32), "Vec3 dot 1.2")
	assertEquals(t, expV, *v32.dot(&v31), "Vec3 dot 2.1")
}

func TestDotProduct4DVectors(t *testing.T) {
	expV := Vec4{-1440, 15, -1120, 396} // v41 . v42

	// check both dot-prod directions
	assertEquals(t, expV, *v41.dot(&v42), "Vec4 dot 1.2")
	assertEquals(t, expV, *v42.dot(&v41), "Vec4 dot 2.1")
}

func TestCrossProduct(t *testing.T) {
	assertEquals(t, Vec3{-301, 156, 1}, *v31.cross(&v32), "Vec3 cross 1x2")
	assertEquals(t, Vec3{301, -156, -1}, *v32.cross(&v31), "Vec3 cross 2x1")

	// extra test: X-axis cross Y-axis should be Z-axis
	xAx, yAx, zAx := Vec3{1, 0, 0}, Vec3{0, 1, 0}, Vec3{0, 0, 1}
	assertEquals(t, zAx, *xAx.cross(&yAx), "Vec3 cross XxY")
}

func TestMultiplicationOf3x3Matrices(t *testing.T) {
	exp1m := Mat3{1543, 1401, 1259, 2823, 2571, 2319, 4103, 3741, 3379} // m31 x m32
	exp2m := Mat3{2855, 2990, 3125, 1830, 1920, 2010, 2466, 2592, 2718} // m32 x m31
	assertEquals(t, exp1m, *m31.times(&m32), "Mat3 mult 1x2")
	assertEquals(t, exp2m, *m32.times(&m31), "Mat3 mult 2x1")
}

func TestMultiplicationOf4x4Matrices(t *testing.T) {
	exp1m := Mat4{1697, 1695, 1441, 1691, 3087, 3075, 2631, 3051, 4477, 4455, 3821, 4411, 5867, 5835, 5011, 5771} // m41 x m42
	exp2m := Mat4{4782, 4964, 5146, 5328, 2814, 2928, 3042, 3156, 3450, 3600, 3750, 3900, 2646, 2732, 2818, 2904} // m42 x m41
	assertEquals(t, exp1m, *m41.times(&m42), "Mat4 mult 1x2")
	assertEquals(t, exp2m, *m42.times(&m41), "Mat4 mult 2x1")
}

func TestMultiplicationOf3VectorAndMatrix(t *testing.T) {
	exp1v := Vec3{1584, 2814, 4044} // m31 * v31
	exp2v := Vec3{3663, 3786, 3909} // v31 * m31
	assertEquals(t, exp1v, *m31.timesVec(&v31), "Vec3 mult Mat3")
	assertEquals(t, exp2v, *v31.timesMat(&m31), "Mat3 mult Vec3")
}

func TestMultiplicationOf4VectorAndMatrix(t *testing.T) {
	exp1v := Vec4{1658, 2908, 4158, 5408} // v41 * m41
	exp2v := Vec4{4205, 4330, 4455, 4580} // m41 * v41
	assertEquals(t, exp1v, *m41.timesVec(&v41), "Vec4 mult Mat4")
	assertEquals(t, exp2v, *v41.timesMat(&m41), "Mat4 mult Vec4")
}

// Helper functions (for checking equality, with error messages)
// Each function returns whether the assert passed

func assert(t *testing.T, cond bool, msg string) bool {
	if !cond {
		t.Errorf(msg)
	}
	return cond
}

// use ONLY where equality is already defined in GO 
func assertEquals(t *testing.T, exp, act interface{}, msg string) bool {
	return assert(t, exp == act, msg+fmt.Sprint(":\n\t\tExp: ", exp, "\n\t\tAct: ", act))
}

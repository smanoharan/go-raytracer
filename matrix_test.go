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
//		- TODO

var (
	v31, v32 Vec3 = Vec3{4, 7, 112}, Vec3{1, 2, -11}
	v41, v42 Vec4 = Vec4{-10, 5, 112, 18}, Vec4{144, 3, -10, 22}
	m31      Mat3 = Mat3{11, 12, 13, 21, 22, 23, 31, 32, 33}
	m32      Mat3 = Mat3{44, 45, 46, 33, 30, 27, 51, 42, 33}
	m41      Mat4 = Mat4{11, 12, 13, 14, 21, 22, 23, 24, 31, 32, 33, 34, 41, 42, 43, 44}
	m42      Mat4 = Mat4{44, 45, 46, 47, 33, 30, 27, 24, 51, 42, 33, 24, 11, 21, 13, 41}
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

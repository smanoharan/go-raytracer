package main

import "math"

// Matrix.go: contains common operations for matrices and vectors.
// Only required types for the raytracer are 3D and 4D vectors and matrices.

// define the length of each data type
const (
	M3LEN = 9
	M4LEN = 16
	V3LEN = 3
	V4LEN = 4
)

type (
	entry float64      // each element of a matrix/vector
	Vec3  [V3LEN]entry // 3D vector
	Vec4  [V4LEN]entry // 4D vector
	Mat3  [M3LEN]entry // 3x3 matrix
	Mat4  [M4LEN]entry // 4x4 matrix
)

// define shorthands:
const (
	cX, cY, cZ, cW = 0, 1, 2, 3 // vector ordinates
 	ZERO, ONE, TWO, FOUR entry = entry(0), entry(1), entry(2), entry(4) // numbers
)

// identity and zero of each type:
var (
	IDENTITY_M3 Mat3 = Mat3{1,0,0,0,1,0,0,0,1}
	ZERO_M3 Mat3 = Mat3{0,0,0,0,0,0,0,0,0}

	IDENTITY_M4 Mat4 = Mat4{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}
	ZERO_M4 Mat4 = Mat4{0,0,0,0,0,0,0,0, 0,0,0,0, 0,0,0,0}

	X_V3, Y_V3, Z_V3 Vec3 = Vec3{1,0,0}, Vec3{0,1,0}, Vec3{0,0,1}
	ZERO_V3 Vec3 = Vec3{0,0,0}

	X_V4, Y_V4, Z_V4, W_V4 Vec4 = Vec4{1,0,0,0}, Vec4{0,1,0,0}, Vec4{0,0,1,0}, Vec4{0,0,0,1}
	ZERO_V4 Vec4 = Vec4{0,0,0,0}
)

// addition: add slices m1 and m2, each with n entries. result is placed in m3. 
func add(m1, m2, m3 []entry, n int) {
	for i := 0; i < n; i++ {
		m3[i] = m1[i] + m2[i]
	}
}

// dot product: 3-vectors
func (m *Vec3) dot(n *Vec3) entry {
	return m[cX]*n[cX] + m[cY]*n[cY] + m[cZ]*n[cZ]
}

// dot product: 4-vectors
func (m *Vec4) dot(n *Vec4) entry {
	return m[cX]*n[cX] + m[cY]*n[cY] + m[cZ]*n[cZ] + m[cW]*n[cW]
}

// elementwise product: 3-vectors
func (m *Vec3) times(n *Vec3) *Vec3 {
	return &Vec3{ m[cX]*n[cX], m[cY]*n[cY], m[cZ]*n[cZ] }
}

// elementwise product: 4-vectors
func (m *Vec4) times(n *Vec4) *Vec4 {
	return &Vec4{ m[cX]*n[cX], m[cY]*n[cY], m[cZ]*n[cZ], m[cW]*n[cW] }
}

// cross product: (only defined for) 3-vectors
func (m *Vec3) cross(n *Vec3) *Vec3 {
	return &Vec3{
		m[cY]*n[cZ] - n[cY]*m[cZ],
		m[cZ]*n[cX] - n[cZ]*m[cX],
		m[cX]*n[cY] - n[cX]*m[cY],
	}
}

// scalar product: matrix (m1) with n entries by a scalar s, result m2
func multScalar(m1, m2 []entry, n int, s entry) {
	for i := 0; i < n; i++ {
		m2[i] = m1[i] * s
	}
}

// multiplication: axn matrix (m1)  by nxb matrix (m2) and 
// place the result into axb matrix (m3)
func mult(m1, m2, m3 []entry, a, n, b int) {

	// iterate through rows of m1
	for row := 0; row < a; row++ {
		ro := row * n // row-offset

		// iterate through cols of m2
		for col := 0; col < b; col++ {

			// compute m1[row] dot-product m2[col]
			sum := entry(0)
			for i := 0; i < n; i++ {
				sum += m1[ro+i] * m2[i*b+col]
			}
			m3[row*b+col] = sum
		}
	}
}

// wrap around math.Sqrt
func sqrt(e entry) entry {
	return entry(math.Sqrt(float64(e)))
}

// wrap around math.Pow
func (x entry) pow(y entry) entry {
	return entry(math.Pow(float64(x), float64(y)))
}

// distanceTo: 3-vectors
func (v *Vec3) distanceTo(u *Vec3) entry {
	dx, dy, dz := v[cX]-u[cX], v[cY]-u[cY], v[cZ]-u[cZ]
	return sqrt(dx*dx + dy*dy + dz*dz)
}

// distanceTo: 4-vectors
func (v *Vec4) distanceTo(u *Vec4) entry {
	dx, dy, dz, dw := v[cX]-u[cX], v[cY]-u[cY], v[cZ]-u[cZ], v[cW]-u[cW]
	return sqrt(dx*dx + dy*dy + dz*dz + dw*dw)
}

// length: 3-vectors
func (v *Vec3) magnitude() entry {
	return sqrt(v.dot(v))
}

// length: 4-vectors
func (v *Vec4) magnitude() entry {
	return sqrt(v.dot(v))
}

// normalized direction: 3-vectors
func (v *Vec3) direction() *Vec3 {
	return v.scale(1 / v.magnitude())
}

// normalized direction: 4-vectors
func (v *Vec4) direction() *Vec4 {
	return v.scale(1 / v.magnitude())
}

// transpose the nxn matrix in m1 and place the result into m2
func transpose(m1, m2 []entry, n int) {
	// iterate through rows and cols:
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			m2[row*n+col] = m1[col*n+row]
		}
	}
}

// determinant of a 2x2 sub-matrix of a nxn matrix
// args: slice containing a nxn matrix, n, (row,col) indices
func det2(m []entry, n, r1, r2, c1, c2 int) entry {
	return (m[r1*n+c1] * m[r2*n+c2]) - (m[r1*n+c2] * m[r2*n+c1])
}

// determinant: 3x3 matrix
func (m *Mat3) determinant() entry {
	ms, n := m[:], V3LEN

	// approach: expand coefficients across the first row
	return 0 + // for formatting
		m[0]*det2(ms, n, 1, 2, 1, 2) +
		m[1]*det2(ms, n, 1, 2, 2, 0) +
		m[2]*det2(ms, n, 1, 2, 0, 1)
}

// determinant: 4x4 matrix
func (m *Mat4) determinant() entry {

	// precompute required 2-determinants:
	// each such 2-det is a 2x2 matrix using only 
	//	columns 2 and 3 (with any 2 rows of 4)
	ms, n := m[:], V4LEN
	d01 := det2(ms, n, 2, 3, 0, 1)
	d02 := det2(ms, n, 2, 3, 0, 2)
	d03 := det2(ms, n, 2, 3, 0, 3)
	d12 := det2(ms, n, 2, 3, 1, 2)
	d13 := det2(ms, n, 2, 3, 1, 3)
	d23 := det2(ms, n, 2, 3, 2, 3)

	// approach: 
	//	expand coefficents across the first row, 
	//	to get four 3x3 matrices.
	//  for each 3x3 matrix, 
	//		expand coefficients across the first row again.
	return 0 + // for formatting
		m[0]*(+m[n+1]*d23-m[n+2]*d13+m[n+3]*d12) +
		m[1]*(-m[n+0]*d23+m[n+2]*d03-m[n+3]*d02) +
		m[2]*(+m[n+0]*d13-m[n+1]*d03+m[n+3]*d01) +
		m[3]*(-m[n+0]*d12+m[n+1]*d02-m[n+2]*d01)
}

// inverse: 3x3 matrices
// only defined if matrix is invertible (i.e. det != 0)
func (m *Mat3) inverse() *Mat3 {
	sc := entry(1) / m.determinant()
	r0, r1, r2 := 0, V3LEN, 2*V3LEN
	return &Mat3{ // formula adapted from the GLM library
		+sc * (m[r1+1]*m[r2+2] - m[r2+1]*m[r1+2]),
		-sc * (m[r0+1]*m[r2+2] - m[r2+1]*m[r0+2]),
		+sc * (m[r0+1]*m[r1+2] - m[r1+1]*m[r0+2]),

		-sc * (m[r1+0]*m[r2+2] - m[r2+0]*m[r1+2]),
		+sc * (m[r0+0]*m[r2+2] - m[r2+0]*m[r0+2]),
		-sc * (m[r0+0]*m[r1+2] - m[r1+0]*m[r0+2]),

		+sc * (m[r1+0]*m[r2+1] - m[r2+0]*m[r1+1]),
		-sc * (m[r0+0]*m[r2+1] - m[r2+0]*m[r0+1]),
		+sc * (m[r0+0]*m[r1+1] - m[r1+0]*m[r0+1]),
	}
}

// inverse: 4x4 matrices
// only defined if matrix is invertible (i.e. det != 0)
// formula adapted from the GLM library
func (m *Mat4) inverse() *Mat4 {
	sc := entry(1) / m.determinant()
	r0, r1, r2, r3 := 0, V4LEN, 2*V4LEN, 3*V4LEN

	// precompute coefficients
	c00 := m[r2+2]*m[r3+3] - m[r3+2]*m[r2+3]
	c02 := m[r1+2]*m[r3+3] - m[r3+2]*m[r1+3]
	c03 := m[r1+2]*m[r2+3] - m[r2+2]*m[r1+3]

	c04 := m[r2+1]*m[r3+3] - m[r3+1]*m[r2+3]
	c06 := m[r1+1]*m[r3+3] - m[r3+1]*m[r1+3]
	c07 := m[r1+1]*m[r2+3] - m[r2+1]*m[r1+3]

	c08 := m[r2+1]*m[r3+2] - m[r3+1]*m[r2+2]
	c10 := m[r1+1]*m[r3+2] - m[r3+1]*m[r1+2]
	c11 := m[r1+1]*m[r2+2] - m[r2+1]*m[r1+2]

	c12 := m[r2+0]*m[r3+3] - m[r3+0]*m[r2+3]
	c14 := m[r1+0]*m[r3+3] - m[r3+0]*m[r1+3]
	c15 := m[r1+0]*m[r2+3] - m[r2+0]*m[r1+3]

	c16 := m[r2+0]*m[r3+2] - m[r3+0]*m[r2+2]
	c18 := m[r1+0]*m[r3+2] - m[r3+0]*m[r1+2]
	c19 := m[r1+0]*m[r2+2] - m[r2+0]*m[r1+2]

	c20 := m[r2+0]*m[r3+1] - m[r3+0]*m[r2+1]
	c22 := m[r1+0]*m[r3+1] - m[r3+0]*m[r1+1]
	c23 := m[r1+0]*m[r2+1] - m[r2+0]*m[r1+1]

	return &Mat4{
		+sc * (m[r1+1]*c00 - m[r1+2]*c04 + m[r1+3]*c08),
		-sc * (m[r0+1]*c00 - m[r0+2]*c04 + m[r0+3]*c08),
		+sc * (m[r0+1]*c02 - m[r0+2]*c06 + m[r0+3]*c10),
		-sc * (m[r0+1]*c03 - m[r0+2]*c07 + m[r0+3]*c11),

		-sc * (m[r1+0]*c00 - m[r1+2]*c12 + m[r1+3]*c16),
		+sc * (m[r0+0]*c00 - m[r0+2]*c12 + m[r0+3]*c16),
		-sc * (m[r0+0]*c02 - m[r0+2]*c14 + m[r0+3]*c18),
		+sc * (m[r0+0]*c03 - m[r0+2]*c15 + m[r0+3]*c19),

		+sc * (m[r1+0]*c04 - m[r1+1]*c12 + m[r1+3]*c20),
		-sc * (m[r0+0]*c04 - m[r0+1]*c12 + m[r0+3]*c20),
		+sc * (m[r0+0]*c06 - m[r0+1]*c14 + m[r0+3]*c22),
		-sc * (m[r0+0]*c07 - m[r0+1]*c15 + m[r0+3]*c23),

		-sc * (m[r1+0]*c08 - m[r1+1]*c16 + m[r1+2]*c20),
		+sc * (m[r0+0]*c08 - m[r0+1]*c16 + m[r0+2]*c20),
		-sc * (m[r0+0]*c10 - m[r0+1]*c18 + m[r0+2]*c22),
		+sc * (m[r0+0]*c11 - m[r0+1]*c19 + m[r0+2]*c23),
	}
}

// Wrappers for each of the types:

// addition: 3x3 matrices
func (m *Mat3) plus(n *Mat3) *Mat3 {
	res := new(Mat3)
	add(m[:], n[:], res[:], M3LEN)
	return res
}

// addition: 4x4 matrices
func (m *Mat4) plus(n *Mat4) *Mat4 {
	res := new(Mat4)
	add(m[:], n[:], res[:], M4LEN)
	return res
}

// addition: 3-vectors
func (m *Vec3) plus(n *Vec3) *Vec3 {
	res := new(Vec3)
	add(m[:], n[:], res[:], V3LEN)
	return res
}

// addition: 4-vectors
func (m *Vec4) plus(n *Vec4) *Vec4 {
	res := new(Vec4)
	add(m[:], n[:], res[:], V4LEN)
	return res
}

// subtraction: 3-vectors
func (m *Vec3) minus(n *Vec3) *Vec3 {
	return m.plus(n.scale(-ONE))
}

// subtraction: 4-vectors
func (m *Vec4) minus(n *Vec4) *Vec4 {
	return m.plus(n.scale(-ONE))
}

// scalar product for 3-vectors
func (m *Vec3) scale(s entry) *Vec3 {
	res := new(Vec3)
	multScalar(m[:], res[:], V3LEN, s)
	return res
}

// scalar product for 4-vectors
func (m *Vec4) scale(s entry) *Vec4 {
	res := new(Vec4)
	multScalar(m[:], res[:], V4LEN, s)
	return res
}

// scalar product for 3x3 matrices
func (m *Mat3) scale(s entry) *Mat3 {
	res := new(Mat3)
	multScalar(m[:], res[:], M3LEN, s)
	return res
}

// scalar product for 4x4 matrices
func (m *Mat4) scale(s entry) *Mat4 {
	res := new(Mat4)
	multScalar(m[:], res[:], M4LEN, s)
	return res
}

// multiplication: 3x3 matrices
func (m *Mat3) times(n *Mat3) *Mat3 {
	res := new(Mat3)
	mult(m[:], n[:], res[:], V3LEN, V3LEN, V3LEN)
	return res
}

// multiplication: 4x4 matrices
func (m *Mat4) times(n *Mat4) *Mat4 {
	res := new(Mat4)
	mult(m[:], n[:], res[:], V4LEN, V4LEN, V4LEN)
	return res
}

// multiplication: 3-vec * 3x3 mat
func (v *Vec3) timesMat(m *Mat3) *Vec3 {
	res := new(Vec3)
	mult(v[:], m[:], res[:], 1, V3LEN, V3LEN)
	return res
}

// multiplication: 3x3 mat * 3-vec
func (m *Mat3) timesVec(v *Vec3) *Vec3 {
	res := new(Vec3)
	mult(m[:], v[:], res[:], V3LEN, V3LEN, 1)
	return res
}

// multiplication: 4-vec * 4x4 mat
func (v *Vec4) timesMat(m *Mat4) *Vec4 {
	res := new(Vec4)
	mult(v[:], m[:], res[:], 1, V4LEN, V4LEN)
	return res
}

// multiplication: 4x4 mat * 4-vec
func (m *Mat4) timesVec(v *Vec4) *Vec4 {
	res := new(Vec4)
	mult(m[:], v[:], res[:], V4LEN, V4LEN, 1)
	return res
}

// transpose: 3-vectors
func (m *Mat3) transpose() *Mat3 {
	res := new(Mat3)
	transpose(m[:], res[:], V3LEN)
	return res
}

// transpose: 4-vectors
func (m *Mat4) transpose() *Mat4 {
	res := new(Mat4)
	transpose(m[:], res[:], V4LEN)
	return res
}

func main() {
	// TODO
	print("\n")
}

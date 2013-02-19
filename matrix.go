package main

// Matrix.go: contains common operations for matrices and vectors.
// Only required types for the raytracer are 3D and 4D vectors and matrices.

// define the length of each data type
const (
	M3LEN = 9
	M4LEN = 16
	V3LEN = 3
	V4LEN = 4
)

// define names for Vector ordinates
const (
	cX = 0
	cY = 1
	cZ = 2
	cW = 3
)

type (
	entry float64      // each element of a matrix/vector
	Mat3  [M3LEN]entry // 3x3 matrix
	Mat4  [M4LEN]entry // 4x4 matrix
	Vec3  [V3LEN]entry // 3D vector
	Vec4  [V4LEN]entry // 4D vector
)

// addition: add slices m1 and m2, each with n entries. result is placed in m3. 
func add(m1, m2, m3 []entry, n int) {
	for i := 0; i < n; i++ {
		m3[i] = m1[i] + m2[i]
	}
}

// dot product: 3-vectors
func (m *Vec3) dot(n *Vec3) *Vec3 {
	return &Vec3{
		m[cX] * n[cX],
		m[cY] * n[cY],
		m[cZ] * n[cZ],
	}
}

// dot product: 4-vectors
func (m *Vec4) dot(n *Vec4) *Vec4 {
	return &Vec4{
		m[cX] * n[cX],
		m[cY] * n[cY],
		m[cZ] * n[cZ],
		m[cW] * n[cW],
	}
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

func main() {
	// TODO
	print("\n")
}

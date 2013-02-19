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

type (
	entry float64      // each element of a matrix/vector
	Mat3  [M3LEN]entry // 3x3 matrix
	Mat4  [M4LEN]entry // 4x4 matrix
	Vec3  [V3LEN]entry // 3D vector
	Vec4  [V4LEN]entry // 4D vector
)

// addition: 3x3 matrices
func (m *Mat3) plus(n *Mat3) *Mat3 {
	res := new(Mat3)
	for i := 0; i < M3LEN; i++ {
		res[i] = m[i] + n[i]
	}
	return res
}

// addition: 4x4 matrices
func (m *Mat4) plus(n *Mat4) *Mat4 {
	res := new(Mat4)
	for i := 0; i < M4LEN; i++ {
		res[i] = m[i] + n[i]
	}
	return res
}

// addition: 3-vectors
func (m *Vec3) plus(n *Vec3) *Vec3 {
	res := new(Vec3)
	for i := 0; i < V3LEN; i++ {
		res[i] = m[i] + n[i]
	}
	return res
}

// addition: 4-vectors
func (m *Vec4) plus(n *Vec4) *Vec4 {
	res := new(Vec4)
	for i := 0; i < V4LEN; i++ {
		res[i] = m[i] + n[i]
	}
	return res
}

// multiplication: 3x3 matrices
func (m *Mat3) times(n *Mat3) *Mat3 {
	res := new(Mat3)

	for row := 0; row < V3LEN; row++ {

		ri := row * V3LEN                        // row index
		mr1, mr2, mr3 := m[ri], m[ri+1], m[ri+2] // get row from m

		for col := 0; col < V3LEN; col++ {

			nc1, nc2, nc3 := n[col], n[col+V3LEN], n[col+2*V3LEN] // get col from n
			res[row*V3LEN+col] = mr1*nc1 + mr2*nc2 + mr3*nc3
		}
	}

	return res
}

// multiplication: 4x4 matrices
func (m *Mat4) times(n *Mat4) *Mat4 {
	res := new(Mat4)

	for row := 0; row < V3LEN; row++ {

		ri := row * V3LEN                        // row index
		mr1, mr2, mr3 := m[ri], m[ri+1], m[ri+2] // get row from m

		for col := 0; col < V3LEN; col++ {

			nc1, nc2, nc3 := n[col], n[col+V3LEN], n[col+2*V3LEN] // get col from n
			res[row*V3LEN+col] = mr1*nc1 + mr2*nc2 + mr3*nc3
		}
	}

	return res
}

func add(a, b entry) entry {
	return a + b
}

func main() {
	// TODO fix
	print(add(entry(1), entry(2)))
	print("\n")
}

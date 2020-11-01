package pattern

// Vertex some stuff here ...
type Vertex struct {
	x int
	y int
}

// NewVertex creates a vertex
func NewVertex(tx, ty int) *Vertex {
	return &Vertex{x: tx, y: ty}
}

// Add
func (v Vertex) Add(x, y int) {
	v.x = v.x + x
	v.y = v.y + y
}

// Sum a , b
func Sum(a int, b int) int {
	return a + b
}

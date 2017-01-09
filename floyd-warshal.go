package fw

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type G struct {
	NVertices int
	Adjacency [][]int
}

// NewRandomG creates a new graph with $occupancy% of the vertices
// adjacent to each other.
//
// occupancy should be between 0 and 1.
func NewRandomG(nVertices int, occupancy float64) (*G, error) {
	if occupancy < 0 || occupancy > 1 {
		return nil, fmt.Errorf("occupancy must be between 0 and 1, not %d", occupancy)
	}

	G := &G{
		NVertices: nVertices,
		Adjacency: make([][]int, nVertices),
	}

	for i := 0; i < nVertices; i++ {
		G.Adjacency[i] = make([]int, nVertices)
	}

	max := big.NewInt(int64(nVertices))
	goal := int(float64(nVertices) * occupancy)
	for n := 0; n < goal && goal != 0; {
		bigI, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, fmt.Errorf("i = rand.Int(max: %v): %s", max, err)
		}
		i := int(bigI.Uint64())
		bigJ, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, fmt.Errorf("j = rand.Int(max: %v): %s", max, err)
		}
		j := int(bigJ.Uint64())
		if G.Adjacency[i][j] == 0 {
			G.Adjacency[i][j] = 1
			n++
		}
	}

	return G, nil
}

func copy2D(src [][]int) (dst [][]int) {
	dst = make([][]int, len(src))
	for k := 0; k < len(src); k++ {
		dst[k] = make([]int, len(src[k]))
		copy(dst[k], src[k])
	}
	return dst
}

// Floyd Warshall dynamic programming algorithm
func (g *G) AllPairsShortestPaths() [][]int {
	// initialize the shortest paths with a copy of the adjacency list
	prev := copy2D(g.Adjacency)
	curr := copy2D(g.Adjacency)

	nVertices := g.NVertices
	for k := 0; k < nVertices; k++ {
		// order of i,j iteration makes big difference
		for i := 0; i < nVertices; i++ {
			for j := 0; j < nVertices; j++ {
				a := prev[i][k] + prev[k][j]
				b := prev[i][j]
				if a < b {
					curr[i][j] = a
				} else {
					curr[i][j] = b
				}
			}
		}
		prev, curr = curr, prev
	}

	return prev
}

package fw

import (
	"testing"
)

const connectedness = 0.8

var result [][]int

var (
	g10, g20, g50, g100, g500, g1000 *G
)

func init() {
	var err error

	g10, err = NewRandomG(10, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}

	g20, err = NewRandomG(20, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}

	g50, err = NewRandomG(50, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}

	g100, err = NewRandomG(100, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}

	g500, err = NewRandomG(500, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}

	g1000, err = NewRandomG(1000, connectedness)
	if err != nil {
		panic("NewRandomG failed")
	}
}

func BenchmarkFW10(b *testing.B) {
	g := g10
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

func BenchmarkFW20(b *testing.B) {
	g := g20
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

func BenchmarkFW50(b *testing.B) {
	g := g50
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

func BenchmarkFW100(b *testing.B) {
	g := g100
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

func BenchmarkFW500(b *testing.B) {
	g := g500
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

func BenchmarkFW1000(b *testing.B) {
	g := g1000
	for i := 0; i < b.N; i++ {
		result = g.AllPairsShortestPaths()
	}
}

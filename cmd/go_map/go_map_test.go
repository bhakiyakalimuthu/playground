package main

import "testing"

var (
	m1 map[uint32]struct{}
	m2 map[uint64]*x
	m3 map[uint32]*x
)

type x struct {
	val int
}

func BenchmarkMaps1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m1 = make(map[uint32]struct{}, 1e6)
	}

}
func BenchmarkMaps2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m2 = make(map[uint64]*x, 1e6)
	}

}
func BenchmarkMaps3(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m3 = make(map[uint32]*x, 1e6)
	}

}

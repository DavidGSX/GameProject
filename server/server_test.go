package main

import (
	"testing"
)

func TestModifydb(t *testing.T) {
	//for i := 0; i < 3000; i++ {
	Modifydb(1002, "money", 1)
	//}
}

func BenchmarkModifydb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Modifydb(1003, "money", 1)
	}
}

func BenchmarkModifydbParallel(b *testing.B) {
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				Modifydb(1004, "money", 1)
			}
		},
	)
}

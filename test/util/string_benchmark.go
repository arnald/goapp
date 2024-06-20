package util

import (
	"goapp/pkg/util"
	"testing"
)

func BenchmarkRandHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.RandHexString(16)
	}
}

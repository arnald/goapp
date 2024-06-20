package util

import (
	"goapp/pkg/util"
	"strings"
	"testing"
)

func TestHandlesZeroLengthInputGracefully(t *testing.T) {
	result := util.RandHexString(0)
	if len(result) != 0 {
		t.Errorf("Expected length 0, but got %d", len(result))
	}
}

func TestRandHexStringValidHexCharacters(t *testing.T) {
	result := util.RandHexString(10)
	for _, char := range result {
		if !strings.Contains("0123456789abcdef", string(char)) {
			t.Errorf("Expected valid hex character, but got %s", string(char))
		}
	}
}

func TestHandlesTypicalLengths(t *testing.T) {
	lengths := []int{8, 16, 32}
	for _, length := range lengths {
		result := util.RandHexString(length)
		if len(result) != length {
			t.Errorf("Expected length %d, but got %d", length, len(result))
		}
	}
}

func TestMemoryUsageForLargeInputs(t *testing.T) {
	length := 1000000
	result := util.RandHexString(length)
	if len(result) != length {
		t.Errorf("Expected length %d, but got %d", length, len(result))
	}
}

func TestRandHexStringDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function panicked: %v", r)
		}
	}()
	_ = util.RandHexString(10)
}

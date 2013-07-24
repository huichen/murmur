package gomurmur

import (
	"testing"
)

func TestMurmur3(t *testing.T) {
	text1 := []byte("The quick brown fox jumps over the lazy dog")
	expectedHash1 := uint32(0x78e69e27)
	actualHash1 := Murmur3(text1)
	if expectedHash1 != actualHash1 {
		t.Errorf("Expected hash value: %d, got %d", expectedHash1, actualHash1)
	}

	text2 := []byte("The quick brown fox jumps over the lazy cog")
	expectedHash2 := uint32(0xd5ece287)
	actualHash2 := Murmur3(text2)
	if expectedHash2 != actualHash2 {
		t.Errorf("Expected hash value: %d, got %d", expectedHash2, actualHash2)
	}
}

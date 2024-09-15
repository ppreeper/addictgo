package main

import "testing"

func TestEncodePassword(t *testing.T) {
	encodePassword("hello")
}

func BenchmarkEncodePassword(b *testing.B) {
	encodePassword("hello")
}

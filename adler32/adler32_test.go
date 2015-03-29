package adler32_test

import (
	rollsum "github.com/chmduquesne/rollinghash/adler32"
	"hash/adler32"
	"testing"
)

func TestRolling(t *testing.T) {
	s := []byte("The brown fox jumps over the lazy dog")

	// window len
	n := 16

	vanilla := adler32.New()
	rolling := rollsum.New()

	// Load the window
	rolling.Write(s[0:n])

	// Roll it and compare the result with full re-calculus every time
	for i := 1; i < len(s)-n; i++ {

		vanilla.Reset()
		vanilla.Write(s[i : i+n])

		err := rolling.Roll(s[i-1], s[i+n-1])
		if err != nil {
			t.Fatal(err)
		}

		if vanilla.Sum32() != rolling.Sum32() {
			t.Fatal("%v: expected %x, got %x",
				s[i:i+n], vanilla.Sum32(), rolling.Sum32())
		}

	}
}

func TestUninitialized(t *testing.T) {
	s := []byte("The brown fox jumps over the lazy dog")
	hash := rollsum.New()
	err := hash.Roll(s[0], s[1])

	if err == nil {
		t.Fatal("Rolling with an uninitialized window should trigger an error")
	}
}

package anagram_test

import (
	"FizzBuzz/anagram"
	"fmt"
	"strings"
	"testing"
)

// NOTE: The tests call IsAnagram(a, b string) which you should implement in the same package.
// The tests assume the behavior described in the problem statement: case-insensitive,
// ignore non-alphanumeric characters, treat runes as-is (no Unicode normalization).

func TestIsAnagram_Basic(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want bool
	}{
		{"simple true", "listen", "silent", true},
		{"case insens true", "Apple", "papel", true},
		{"spaces and punctuation", "A gentleman", "Elegant man", true},
		{"famous phrase", "William Shakespeare", "I am a weakish speller", true},
		{"ignore punctuation", "Hello", "Olelh!", true},
		{"numbers", "12345", "54321", true},
		{"unicode kanji", "東京", "京東", true},
		{"repeated chars", "aabbcc", "abcabc", true},
		{"different lengths", "abc", "ab", false},
		{"not anagram", "hello", "world", false},
		{"mixed nonalpha", "foo!!", "ofo", true},
		{"empty strings", "", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := anagram.IsAnagram(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("IsAnagram(%q, %q) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// Additional tests to exercise edge cases and case normalization
func TestIsAnagram_EdgeCases(t *testing.T) {
	// long string made by repeating a pattern
	pat := strings.Repeat("aA1! ", 1000) // should normalize to same counts for both
	a := pat + "bbb"
	b := strings.Repeat("1aA ", 1000) + "bbb"
	if !anagram.IsAnagram(a, b) {
		t.Fatalf("Expected long repeated strings to be anagrams")
	}

	// strings with different counts of a rune
	if anagram.IsAnagram("aabb", "abbb") {
		t.Fatalf("Expected aabb and abbb to NOT be anagrams")
	}
}

// Example for godoc / fmt.Example
func ExampleIsAnagram() {
	fmt.Println(anagram.IsAnagram("A gentleman", "Elegant man"))
	fmt.Println(anagram.IsAnagram("Hello", "World"))
	// Output:
	// true
	// false
}

// Benchmark to check performance for larger inputs
func BenchmarkIsAnagram(b *testing.B) {
	// create two large anagram strings
	const n = 200000
	base := strings.Repeat("abcd1234! ", n/10) // includes some ignored chars
	// Build a and b with same counts but different ordering
	a := base + strings.Repeat("x", n/10)
	bs := []rune(base + strings.Repeat("x", n/10))
	// shuffle bs pseudo-randomly but deterministically (simple rotate)
	rot := len(bs) / 3
	copy(bs, append(bs[rot:], bs[:rot]...))
	bStr := string(bs)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = anagram.IsAnagram(a, bStr)
	}
}

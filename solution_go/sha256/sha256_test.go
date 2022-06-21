package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/mdwhatcott/sha256_project/solution_go/assert"
)

func TestSHA256(t *testing.T) {
	assertHash(t, []byte("Hello, world!"))
	assertHash(t, []byte("Hello, worlds!"))
	assertHash(t, []byte("asdfqwerzxcv"))
}

func assertHash(t *testing.T, input []byte) {
	hash := sha256.New()
	hash.Write(input)
	sum := hash.Sum(nil)
	expected := hex.EncodeToString(sum)
	actual := hex.EncodeToString(Hash(input))
	assert.That(t, actual).Equals(expected)
	t.Logf("Hash: %s Input: %s", actual, string(input))
}

func BenchmarkStdlib(b *testing.B) {
	input := []byte(strings.Repeat(fmt.Sprint(b.N), 100))
	hash := sha256.New()
	var result []byte
	for i := 0; i < b.N; i++ {
		hash.Reset()
		hash.Write(input)
		result = hash.Sum(nil)
	}
	b.Log(result)
}

func BenchmarkProject(b *testing.B) {
	input := []byte(strings.Repeat(fmt.Sprint(b.N), 100))
	var result []byte
	for i := 0; i < b.N; i++ {
		result = Hash(input)
	}
	b.Log(result)
}

package ShortRequestID

import (
	"crypto/rand"
	"time"
)

// generateRandomBytes generates a safe random bytes data
func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fallback := time.Now().Nanosecond()
		for i := range b {
			b[i] = byte(fallback >> (i * 8) & 0xFF)
		}
	}
	return b
}

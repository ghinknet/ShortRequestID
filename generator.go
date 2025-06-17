package ShortRequestID

import (
	"encoding/base32"
	"time"
)

// GenerateRequestID generates a BASE32 based Request ID
func GenerateRequestID() string {
	timestamp := time.Now().UnixMilli()

	randomSuffix := generateRandomBytes(2)

	data := make([]byte, 10)

	for i := 7; i >= 0; i-- {
		data[i] = byte(timestamp & 0xFF)
		timestamp >>= 8
	}

	copy(data[8:], randomSuffix)

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	encoded := encoder.EncodeToString(data)

	return encoded
}

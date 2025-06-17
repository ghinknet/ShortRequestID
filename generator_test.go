package ShortRequestID

import "testing"

func TestGenerateRequestID(t *testing.T) {
	if len(GenerateRequestID()) != 16 {
		t.Errorf("GenerateRequestID() should generate 16 characters")
	}
}

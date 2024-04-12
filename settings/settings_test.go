package settings

import (
	"testing"
)

func TestEngineExists(t *testing.T) {
	if DefaultSettings.Host != "127.0.0.1" && DefaultSettings.Port != 8080 {
		t.Errorf("Unexpected default configuration for HOST and PORT provided")
	}
}

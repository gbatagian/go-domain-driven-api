package settings

import (
	"os"
	"testing"
)

func TestEngineExists(t *testing.T) {
	if DefaultSettings.Host != os.Getenv("API_HOST") && DefaultSettings.Port != os.Getenv("API_PORT") {
		t.Errorf("Unexpected default configuration for HOST and PORT provided")
	}
}

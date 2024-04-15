package settings

import (
	"os"
	"testing"
)

func TestDefaultSettings(t *testing.T) {
	if DefaultSettings.Host != "0.0.0.0" && DefaultSettings.Port != "8080" {
		t.Errorf("Unexpected default configuration for HOST and PORT provided: %s:%s", DefaultSettings.Host, DefaultSettings.Port)
	}
}

func TestEnvSettings(t *testing.T) {
	os.Setenv("API_HOST", "192.168.1.54")
	os.Setenv("API_PORT", "8888")
	EnvSettings.Host = getApiHost()
	EnvSettings.Port = getApiPort()

	if EnvSettings.Host != "192.168.1.54" && EnvSettings.Port != "8888" {
		t.Errorf("Unexpected env configuration for HOST and PORT provided: %s:%s", EnvSettings.Host, EnvSettings.Port)
	}
}

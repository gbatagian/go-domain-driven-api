package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngineExists(t *testing.T) {

	assert.NotNil(t, ENGINE)

}

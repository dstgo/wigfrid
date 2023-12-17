package assets

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteBanner(t *testing.T) {
	// nil data
	err := WriteBanner(os.Stdout, nil)
	assert.Nil(t, err)
}

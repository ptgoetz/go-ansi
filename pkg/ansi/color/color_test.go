package color

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestColor_HexString(t *testing.T) {
	assert.Assert(t, Black.HexString() == "#000000")
}

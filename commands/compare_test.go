package commands

import (
	"testing"

	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"
)

func TestParseRange(t *testing.T) {
	s := "1.0..2.0"
	assert.Equal(t, "1.0...2.0", parseCompareRange(s))

	s = "1.0...2.0"
	assert.Equal(t, "1.0...2.0", parseCompareRange(s))
}

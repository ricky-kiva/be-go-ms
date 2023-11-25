package main

// download by `go get github.com/stretchr/testify/assert`
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run by `go`
func TestMathExpressionValue(t *testing.T) {
	var expression = 8 * 5
	assert.Equal(t, expression, 40, "Values are not the same")
}

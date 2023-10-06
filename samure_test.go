package samure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamure(t *testing.T) {
	ctx, err := CreateContextWithBackend(nil, &RawBackend{})
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()
}

package samure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamure(t *testing.T) {
	raw, err := InitRawBackend()
	if !assert.Nil(t, err) {
		return
	}
	ctx, err := CreateContextWithBackend(nil, raw)
	if !assert.Nil(t, err) {
		return
	}
	defer ctx.Destroy()
}

package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	bridge, err := New(nil, "", Config{})
	assert.Nil(t, bridge)
	assert.Error(t, err)
}

func TestNewValid(t *testing.T) {
	Register(new(fakeFactory), "fake")
	// Note: the following is valid for New() since it does not
	// actually connect to docker.
	bridge, err := New(nil, "fake://", Config{})

	assert.NotNil(t, bridge)
	assert.NoError(t, err)
}

func Test_getServiceNameFromImage(t *testing.T) {
	tests := []struct {
		name    string
		image   string
		service string
	}{
		{
			name:    "with tag",
			image:   "123456789012.dkr.ecr.eu-west-1.amazonaws.com/foo/bar:v1.2.3",
			service: "bar",
		},
		{
			name:    "with digest",
			image:   "123456789012.dkr.ecr.eu-west-1.amazonaws.com/foo/bar@sha256:abcc3e12da56ecf5756b48502baa2a98d09e69c5d033dc37ce93de77dc6cb123",
			service: "bar",
		},
		{
			name:    "with both",
			image:   "123456789012.dkr.ecr.eu-west-1.amazonaws.com/foo/bar:v1.2.3@sha256:abcc3e12da56ecf5756b48502baa2a98d09e69c5d033dc37ce93de77dc6cb123",
			service: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.service, getServiceNameFromImage(tt.image), "getServiceNameFromImage(%v)", tt.image)
		})
	}
}

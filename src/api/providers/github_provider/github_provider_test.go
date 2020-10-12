package github_provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("abc123")
	assert.EqualValues(t, header, "token abc123")
}

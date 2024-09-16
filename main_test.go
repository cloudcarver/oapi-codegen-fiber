package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "Get", toCamelCase("GET"))
}

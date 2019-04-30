package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrichTemplates(t *testing.T) {
	enrichedStr := EnrichTemplateStr("hello {{git_user}}")
	assert.Equal(t, "hello user", enrichedStr, "Expecting `hello user`")
}

package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileParser_thenReturnContent(t *testing.T) {
	assertions := assert.New(t)
	parser := TableServiceContentParser{"local/path"}

	contentMap := parser.Parse()
	testContent := contentMap["test"]

	assertions.NotNil(testContent, "TestContent must not be null")
}

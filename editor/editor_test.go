package editor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	EXAMPLE_MARKDOWN_CONTENT = `cool markdown
<!--START_SECTION:chatgptme-->
this will be replaced by chatgpt with some random cool text
<!--END_SECTION:chatgptme-->`
)

func TestWriteMarkdown(t *testing.T) {
	filePath := "test.md"
	err := WriteMarkdown(filePath, EXAMPLE_MARKDOWN_CONTENT)
	assert.NoError(t, err)
}

func TestReplaceSection(t *testing.T) {
	filePath := "test.md"
	err := WriteMarkdown(filePath, EXAMPLE_MARKDOWN_CONTENT)
	markdownTxt, err := ReadMarkdown(filePath)
	newTxt := FindAndReplace(markdownTxt, "hello")
	WriteMarkdown(filePath, newTxt)
	assert.NoError(t, err)
	assert.Equal(t, newTxt, `cool markdown
	hello`)
}

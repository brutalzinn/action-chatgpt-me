package commit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	err := SetConfig(GitConfig{
		Username: "my name custom",
		Email:    "blabla@noreply.com",
	})
	assert.NoError(t, err)
}

func TestCommit(t *testing.T) {
	err := Commit("test commmit")
	assert.NoError(t, err)
}
func TestAdd(t *testing.T) {
	err := Add()
	assert.NoError(t, err)
}

func TestPush(t *testing.T) {
	err := Push()
	assert.NoError(t, err)
}

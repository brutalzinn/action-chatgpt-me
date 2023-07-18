package integrations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWakatimeUserStats(t *testing.T) {
	userStats, err := GetUserStats()
	assert.NoError(t, err)
	assert.NotNil(t, userStats)
}

func TestWakaTimeUserMetrics(t *testing.T) {
	userStats, err := GetUserStats()
	assert.NoError(t, err)
	assert.NotNil(t, userStats)
	messages := "This is the time that i recognize that i need some vocation."
	for _, item := range userStats.Data.Languages {
		messages += item.Name + " " + item.Time
	}
}

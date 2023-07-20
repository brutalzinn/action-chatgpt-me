package integrations

import (
	"fmt"
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
	messages := "My cool profile with wakatime metrics"
	for _, item := range userStats.Data.Languages {
		messages += fmt.Sprintf("%s studies hour with %s\n", item.Time, item.Name)
	}
	t.Log(messages)
}

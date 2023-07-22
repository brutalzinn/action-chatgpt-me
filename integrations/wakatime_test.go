package integrations

import (
	"fmt"
	"strings"
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
	var text strings.Builder
	text.WriteString("My cool profile with wakatime metrics")
	for _, item := range userStats.Data.OperationalSystem[:3] {
		text.WriteString(fmt.Sprintf("- %s", item.Name))
	}
	t.Log(text)
}

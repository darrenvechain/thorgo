package thorest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_FeesHistory(t *testing.T) {
	best, err := thorClient.BestBlock()
	assert.NoError(t, err)
	history, err := thorClient.FeesHistory(best.Number, 100)
	assert.NoError(t, err)
	assert.NotNil(t, history)
}

func TestClient_FeesPriority(t *testing.T) {
	priority, err := thorClient.FeesPriority()
	assert.NoError(t, err)
	assert.NotNil(t, priority)
}

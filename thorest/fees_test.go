package thorest_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/stretchr/testify/assert"
)

func TestClient_FeesHistory(t *testing.T) {
	history, err := thorClient.FeesHistory(thorest.RevisionBest(), 100, []float64{})
	assert.NoError(t, err)
	assert.NotNil(t, history)
}

func TestClient_FeesPriority(t *testing.T) {
	priority, err := thorClient.FeesPriority()
	assert.NoError(t, err)
	assert.NotNil(t, priority)
}

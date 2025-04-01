package thorest_test

import (
	"os"
	"testing"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/stretchr/testify/assert"
)

var (
	galactica   = os.Getenv("IS_GALACTICA")
	isGalactica = galactica == "true"
)

func TestClient_FeesHistory(t *testing.T) {
	if !isGalactica {
		t.Skip("Skipping test as it is only available in Galactica")
	}

	history, err := thorClient.FeesHistory(thorest.RevisionBest(), 100)
	assert.NoError(t, err)
	assert.NotNil(t, history)
}

func TestClient_FeesPriority(t *testing.T) {
	if !isGalactica {
		t.Skip("Skipping test as it is only available in Galactica")
	}
	priority, err := thorClient.FeesPriority()
	assert.NoError(t, err)
	assert.NotNil(t, priority)
}

package thorest_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/thorest"
)

var thorClient *thorest.Client

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	m.Run()
}

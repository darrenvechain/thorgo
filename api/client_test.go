package api_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/api"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
)

var thorClient *api.Client

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	m.Run()
}

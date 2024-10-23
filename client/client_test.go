package client_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
)

var thorClient *client.Client

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	m.Run()
}

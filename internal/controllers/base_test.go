package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/internal/util"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store models.Store) *Server {
	config := util.Config{}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

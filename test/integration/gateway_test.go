package integration

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MiralSNk/go-micro-blog/internal/gateway/router"
	"github.com/MiralSNk/go-micro-blog/internal/logger"
	"github.com/go-jose/go-jose/v4/testutils/require"
	"github.com/stretchr/testify/assert"
)

func TestGatewayHTTP(t *testing.T) {
	log := logger.New("test", "prod")
	r := router.New(log)
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// читаем тело
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello!", string(body))
}

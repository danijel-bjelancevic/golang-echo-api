package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	ctx, rec, handler := setUp("GET", "admin/healthcheck", nil, t)
	is := require.New(t)

	if !assert.NoError(t, handler.Healthcheck(ctx)) {
		is.FailNow("failed during execution of healthcheck request")
	}

	is.Equal(http.StatusOK, rec.Code)

	compareWithGolden(is, t.Name(), rec.Body.String())
}

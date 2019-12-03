package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestLivetest(t *testing.T) {
	ctx, rec, handler := setUp("GET", "admin/livetest", nil, t)
	is := require.New(t)

	if !assert.NoError(t, handler.Livetest(ctx)) {
		is.FailNow("failed during execution of livetest request")
	}

	is.Equal(http.StatusOK, rec.Code)

	compareWithGolden(is, t.Name(), rec.Body.String())
}

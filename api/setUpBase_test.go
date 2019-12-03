package api

import (
	"flag"
	"github.com/danijel-bjelancevic/golang-echo-api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// SetUp takes http method, path and request body and sets up everything to avoid duplication
func setUp(method string, path string, body io.Reader, t *testing.T) (echo.Context, *httptest.ResponseRecorder, *Handler) {
	e := Handler{}.NewRouter()
	req := httptest.NewRequest(method, "/", body)
	req.Header.Set(mdw.RequestId, "11111111-1111-1111-1111-111111111111")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath(path)

	logger := logrus.New()
	h := &Handler{
		Logger: logger,
	}

	return ctx, rec, h
}

var update = flag.Bool("update", false, "update golden files")

func compareWithGolden(is *require.Assertions, testName, actual string) {
	if *update {
		is.NoError(ioutil.WriteFile(filepath.Join("testdata", testName+".golden"), []byte(actual), 0644))
	}

	golden, err := ioutil.ReadFile(filepath.Join("testdata", testName+".golden"))
	is.NoError(err)

	is.EqualValues(string(golden), actual)
}

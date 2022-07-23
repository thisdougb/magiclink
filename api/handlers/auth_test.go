//go:build dev || test

package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

// we can't really test the html output here
var TestItems = []struct {
	comment      string
	MagicLinkID  string
	httpMethod   string
	expectStatus int
}{
	{
		comment:      "valid request",
		MagicLinkID:  "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		httpMethod:   "GET",
		expectStatus: 302,
	},
	{
		comment:      "invalid http method",
		MagicLinkID:  "MBm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		httpMethod:   "POST",
		expectStatus: 405,
	},
	{
		comment:      "invalid id",
		MagicLinkID:  "Bm7vHhhfa9nE5gaiWyXEbvdyRjYf1XJWKj3UJsMGhfkGl36AXCWwYdPTSWWPbpb",
		httpMethod:   "GET",
		expectStatus: 400,
	},
	{
		comment:      "datastore error",
		MagicLinkID:  "sVm4ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t",
		httpMethod:   "GET",
		expectStatus: 500,
	},
	{
		comment:      "magic link not found",
		MagicLinkID:  "MBm7ECyEaec1HYBI9yP8nqLPMP1f8PXSar2O1ZN5HzyNn1WCr5Zx7JuInMUB8o8t",
		httpMethod:   "GET",
		expectStatus: 302,
	},
}

func TestAuthWeb(t *testing.T) {

	// create our mock service
	r := auth.NewMockRepository()
	authService := auth.NewService(r)

	// inject mock service
	env := &Env{AuthService: authService}

	for _, item := range TestItems {

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(env.Auth)

		httpURL := fmt.Sprintf("http://localhost%s%s%s", env.GetURLPrefix(), "/auth/", item.MagicLinkID)

		req, err := http.NewRequest(item.httpMethod, httpURL, nil)
		if err != nil {
			assert.Fail(t, item.comment)
		}

		handler.ServeHTTP(rr, req)

		assert.Equal(t, item.expectStatus, rr.Code, item.comment)
	}
}

package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/thisdougb/magiclink/pkg/usecase/owner"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// we can't really test the html output here
var SessionOwnerTestItems = []struct {
	comment        string
	httpURL        string
	httpMethod     string
	bodyData       string
	expectBodyData string
	expectStatus   int
}{
	{
		comment:        "valid request",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"12345","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   200,
		expectBodyData: `{"Owner":"valid@session.owner"}`,
	},
	{
		comment:        "invalid method",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "GET",
		bodyData:       `{"token":"user@domain.com","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   405,
		expectBodyData: "Method Not Allowed\n",
	},
	{
		comment:        "invalid json input",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"user@domain.com","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"`,
		expectStatus:   400,
		expectBodyData: "Bad Request\n",
	},
	{
		comment:        "invalid access token",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"11111","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   401,
		expectBodyData: "Unauthorized\n",
	},
	{
		comment:        "session not found",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"12345","session":"999991lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   404,
		expectBodyData: "Not Found\n",
	},
}

func TestSessionOwnerRequest(t *testing.T) {

	os.Setenv("MAGICLINK_SESSION_OWNER_ACCESS_TOKENS", "12345")
	defer os.Unsetenv("MAGICLINK_SESSION_OWNER_ACCESS_TOKENS")

	// create our mock service
	r := owner.NewMockRepository()
	ownerService := owner.NewService(r)

	// inject mock service
	env := &Env{OwnerService: ownerService}

	for _, item := range SessionOwnerTestItems {

		// httptest lets us interrogate the http response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(env.Owner)

		bodyReader := bytes.NewReader([]byte(item.bodyData))

		req, err := http.NewRequest(item.httpMethod, item.httpURL, bodyReader)
		if err != nil {
			assert.Fail(t, item.comment)
		}
		req.Header.Add("Content-Type", "encoding/json")

		handler.ServeHTTP(rr, req)

		assert.Equal(t, item.expectStatus, rr.Code, item.comment)
		assert.Equal(t, item.expectBodyData, rr.Body.String(), item.comment)
	}
}

// we can't really test the html output here
var SessionMultipleTokenItems = []struct {
	comment        string
	httpURL        string
	httpMethod     string
	bodyData       string
	expectBodyData string
	expectStatus   int
}{
	{
		comment:        "valid request first token",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"54321","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   200,
		expectBodyData: `{"Owner":"valid@session.owner"}`,
	},
	{
		comment:        "valid request second token",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"12345","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   200,
		expectBodyData: `{"Owner":"valid@session.owner"}`,
	},
	{
		comment:        "invalid access token",
		httpURL:        "http://localhost/sessionowner/",
		httpMethod:     "POST",
		bodyData:       `{"token":"99999","session":"1tleL1lgn0UDADpa1UhEcmga6x5j8YkFNRvhCAZNysxLQzzlmKgTP5wFvgdPfgPn"}`,
		expectStatus:   401,
		expectBodyData: "Unauthorized\n",
	},
}

func TestMultipleAccessTokens(t *testing.T) {

	// set multiple access tokens the service will accept,
	// simulates rolling the token
	os.Setenv("MAGICLINK_SESSION_OWNER_ACCESS_TOKENS", "54321,12345")
	defer os.Unsetenv("MAGICLINK_SESSION_OWNER_ACCESS_TOKENS")

	// create our mock service
	r := owner.NewMockRepository()
	ownerService := owner.NewService(r)

	// inject mock service
	env := &Env{OwnerService: ownerService}

	for _, item := range SessionMultipleTokenItems {

		// httptest lets us interrogate the http response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(env.Owner)

		bodyReader := bytes.NewReader([]byte(item.bodyData))

		req, err := http.NewRequest(item.httpMethod, item.httpURL, bodyReader)
		if err != nil {
			assert.Fail(t, item.comment)
		}
		req.Header.Add("Content-Type", "encoding/json")

		handler.ServeHTTP(rr, req)

		assert.Equal(t, item.expectStatus, rr.Code, item.comment)
		assert.Equal(t, item.expectBodyData, rr.Body.String(), item.comment)
	}
}

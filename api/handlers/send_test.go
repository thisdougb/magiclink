package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/thisdougb/magiclink/pkg/usecase/send"
	"net/http"
	"net/http/httptest"
	"testing"
)

// we can't really test the html output here
var RequestLinkTestItems = []struct {
	comment      string
	httpURL      string
	httpMethod   string
	bodyData     string
	expectStatus int
}{
	{
		comment:      "valid request",
		httpURL:      "http://localhost/requestlink/",
		httpMethod:   "POST",
		bodyData:     `{"email":"user@domain.com"}`,
		expectStatus: 200,
	},
	{
		comment:      "invalid http method",
		httpURL:      "http://localhost/requestlink/",
		httpMethod:   "GET",
		bodyData:     `{"email":"user@domain.com"}`,
		expectStatus: 405,
	},
	{
		comment:      "malformed json body",
		httpURL:      "http://localhost/requestlink/",
		httpMethod:   "POST",
		bodyData:     `{"email":"user@domain.com`,
		expectStatus: 400,
	},
	{
		comment:      "trigger datastore error",
		httpURL:      "http://localhost/requestlink/",
		httpMethod:   "POST",
		bodyData:     `{"email":"fail@datastore.error"}`,
		expectStatus: 500,
	},
}

func TestMagicLinkRequestWeb(t *testing.T) {

	// create our mock service
	r := send.NewMockRepository()
	sendService := send.NewService(r)

	// inject mock service
	env := &Env{SendService: sendService}

	for _, item := range RequestLinkTestItems {

		// httptest lets us interrogate the http response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(env.Send)

		bodyReader := bytes.NewReader([]byte(item.bodyData))

		req, err := http.NewRequest(item.httpMethod, item.httpURL, bodyReader)
		if err != nil {
			assert.Fail(t, item.comment)
		}
		req.Header.Add("Content-Type", "encoding/json")

		handler.ServeHTTP(rr, req)

		assert.Equal(t, item.expectStatus, rr.Code, item.comment)
	}
}

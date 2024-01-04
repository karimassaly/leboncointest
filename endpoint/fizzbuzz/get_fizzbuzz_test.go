package fizzbuzz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFizzBuzz(t *testing.T) {
	for _, tt := range []struct {
		name                          string
		Int1, Int2, Limit, Str1, Str2 string
		Body                          string
	}{
		{
			name:  "int1 is not a number",
			Int1:  "abc",
			Int2:  "5",
			Limit: "15",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"error\":\"invalid parameters: int1, int2, and limit must be integers\"}",
		},
		{
			name:  "int2 is not a number",
			Int1:  "3",
			Int2:  "xyz",
			Limit: "15",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"error\":\"invalid parameters: int1, int2, and limit must be integers\"}",
		},
		{
			name:  "limit is not a number",
			Int1:  "3",
			Int2:  "5",
			Limit: "-----////",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"error\":\"invalid parameters: int1, int2, and limit must be integers\"}",
		},
		{
			name:  "limit is 0",
			Int1:  "3",
			Int2:  "5",
			Limit: "0",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"error\":\"invalid parameters: int1, int2, and limit must be greater than 0\"}",
		},
		{
			name:  "success long limit",
			Int1:  "3",
			Int2:  "5",
			Limit: "200",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"result\":\"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz,101,fizz,103,104,fizzbuzz,106,107,fizz,109,buzz,fizz,112,113,fizz,buzz,116,fizz,118,119,fizzbuzz,121,122,fizz,124,buzz,fizz,127,128,fizz,buzz,131,fizz,133,134,fizzbuzz,136,137,fizz,139,buzz,fizz,142,143,fizz,buzz,146,fizz,148,149,fizzbuzz,151,152,fizz,154,buzz,fizz,157,158,fizz,buzz,161,fizz,163,164,fizzbuzz,166,167,fizz,169,buzz,fizz,172,173,fizz,buzz,176,fizz,178,179,fizzbuzz,181,182,fizz,184,buzz,fizz,187,188,fizz,buzz,191,fizz,193,194,fizzbuzz,196,197,fizz,199,buzz\"}",
		},
		{
			name:  "success",
			Int1:  "3",
			Int2:  "5",
			Limit: "15",
			Str1:  "fizz",
			Str2:  "buzz",
			Body:  "{\"result\":\"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz\"}",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.GET("/fizzbuzz", GetEndpoint())

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
			q := req.URL.Query()
			q.Add("int1", tt.Int1)
			q.Add("int2", tt.Int2)
			q.Add("limit", tt.Limit)
			q.Add("str1", tt.Str1)
			q.Add("str2", tt.Str2)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.Body, w.Body.String())
		})
	}
}

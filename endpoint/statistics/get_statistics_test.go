package statistics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock "leboncointest/pkg/mock/storage"
	"leboncointest/storage"
	endpointstore "leboncointest/storage/endpointstore/sql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockEndpointstore = &mock.MockEndpointRepository{
	GetAllFunc: func() ([]endpointstore.EndpointStatistics, error) {
		return []endpointstore.EndpointStatistics{
			{
				Route:  "/fizzbuzz",
				Method: "GET",
				Count:  5,
			},
			{
				Route:  "/satistics",
				Method: "GET",
				Count:  3,
			},
		}, nil
	},
	CountUpFunc: func(es endpointstore.EndpointStatistics) error {
		return nil
	},
}

func TestGetStatistics(t *testing.T) {
	for _, tt := range []struct {
		name string
		Body string
	}{
		{
			name: "success",
			Body: "{\"statistics\":[{\"ID\":0,\"Route\":\"/fizzbuzz\",\"Method\":\"GET\",\"Count\":5},{\"ID\":0,\"Route\":\"/satistics\",\"Method\":\"GET\",\"Count\":3}]}",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var sr = storage.Repositories{
				EndpointStatistics: mockEndpointstore,
			}

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.GET("/statistics", GetStatistics(&sr))

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/statistics", nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.Body, w.Body.String())
		})
	}
}

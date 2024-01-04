package mock

import (
	"errors"

	endpointstore "leboncointest/storage/endpointstore/sql"
)

type MockEndpointRepository struct {
	GetAllFunc  func() ([]endpointstore.EndpointStatistics, error)
	CountUpFunc func(es endpointstore.EndpointStatistics) error
}

func (m *MockEndpointRepository) GetAll() ([]endpointstore.EndpointStatistics, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, errors.New("GetAll not implemented")
}

func (m *MockEndpointRepository) CountUp(es endpointstore.EndpointStatistics) error {
	if m.CountUpFunc != nil {
		return m.CountUpFunc(es)
	}
	return errors.New("CountUp not implemented")
}

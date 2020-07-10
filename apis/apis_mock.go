package apis

import "github.com/stretchr/testify/mock"

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Get(url string) ([]byte, error) {
	ret := m.Called(url)
	return ret.Get(0).([]byte), ret.Error(1)
}

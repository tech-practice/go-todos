package apis

type MockHttpClient struct{}

func (m *MockHttpClient) Get(url string) ([]byte, error) {
	return MockGet(url)
}

var MockGet func(string) ([]byte, error)

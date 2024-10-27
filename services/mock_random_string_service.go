package services

type MockRandomStringService struct{}

func (m *MockRandomStringService) GetRandomString(length int) (string, error) {
	return "MOCKSTRING", nil
}

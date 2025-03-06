package kms

// MockKMS provides a mock implementation of KMSProvider
type MockKMS struct{}

// Init simulates initializing KMS client
func (m *MockKMS) Init() error {
	return nil
}

// GetAwsSecretValue simulates getting a secret value
func (m *MockKMS) GetAwsSecretValue(key string) (string, error) {
	// If not found, return a fixed string
	return "fake-password", nil
}

// GetAwsSecretData returns all currently saved mock secrets
func (m *MockKMS) GetAwsSecretData() map[string]error {
	// Copy or return directly
	return nil
}

package kms

// KMSProvider defines the interface for KMS operations
type KMSProvider interface {
	Init() error
	GetAwsSecretValue(key string) (string, error)
	GetAwsSecretData() map[string]error
}

// Global provider variable - defaults to MockKMS
var Provider KMSProvider = &MockKMS{}

// Init initializes the KMS provider
func Init() error {
	return Provider.Init()
}

// GetAwsSecretValue retrieves a secret value by key
func GetAwsSecretValue(key string) (string, error) {
	return Provider.GetAwsSecretValue(key)
}

// GetAwsSecretData returns all secrets data
func GetAwsSecretData() map[string]error {
	return Provider.GetAwsSecretData()
}

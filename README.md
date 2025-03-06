# KMS Interface

This is a KMS interface library that provides both interface definitions and default mock implementations for KMS operations.

## Usage

### Basic Usage

```go
import "github.com/okx/kms-interface/kms"

// Use default mock implementation
func main() {
    // Initialize KMS (no-op in mock mode)
    kms.Init()
    
    // Get a secret value (returns mock data)
    value, err := kms.GetAwsSecretValue("my-secret")
    if err != nil {
        // Handle error
    }
    
    // Use the value
    fmt.Println(value) // Will print "fake-password"
}
```

### Implementing Real KMS Provider

To use a real KMS implementation in internal environments, create a bridge package:

```go
// In your internal bridge package
package kmsbridge

import (
    "github.com/okx/kms-interface/kms"
    realkms "gitlab.okg.com/okcoin-commons/ok-kms-go-client/kms"
)

// RealKMS connects to the actual KMS implementation
type RealKMS struct{}

func (r *RealKMS) Init() error {
    return realkms.Init()
}

func (r *RealKMS) GetAwsSecretValue(key string) (string, error) {
    return realkms.GetAwsSecretValue(key)
}

func (r *RealKMS) GetAwsSecretData() map[string]error {
    return realkms.GetAwsSecretData()
}

// SetupRealKMS registers the real implementation
func SetupRealKMS() {
    kms.Provider = &RealKMS{}
}
```

Then in your internal environment initialization code:

```go
import "your-internal-repo/kmsbridge"

func init() {
    kmsbridge.SetupRealKMS()
}
```

## Benefits

1. Public code doesn't reference internal paths
2. Same behavior in external and internal environments, just different implementations
3. Minimal code changes required
4. Easily extensible to support other KMS implementations 
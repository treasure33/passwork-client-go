# passwork-client-go
REST Client for the Password Manager Passwork written in Go.

The client can currently perform CRUD operations on passwords, folders and vaults.

## Example usage

```go
package main

import "github.com/treasure33/passwork-client-go"

func main() {
	host := "https://my-passwork-instance.com/api/v1"
	apiKey := "my-secret-api-key"
	timeout := time.Second * 30

	// Create a new client and log in
	client := passwork.NewClient(host, apiKey, timeout)
	client.Login()

	// Create a vault
	vaultRequest := VaultAddRequest{
		Name:         "example-vault",
		IsPrivate:    true,
		PasswordHash: "example-hash",
		Salt:         "example-salt",
		MpCrypted:    "example-mp",
	}
	vaultResponse, _ := client.AddVault(vaultRequest)

	// Create a password
	passwordRequest := PasswordRequest{
		Name:            "example-password",
		VaultId:         vaultResponse.Data,
		Login:           "example-login",
		CryptedPassword: "ZXhhbXBsZS1wYXNzd29yZAo=", // Password must be base64 encoded
		Description:     "example-description",
		Url:             "https://example.com",
		Color:           1,
		Tags:            []string{"example", "tag"},
	}
	client.AddPassword(passwordRequest)

	// Logout
	client.Logout()
}

```

## Running tests

### Option 1: Using .env file (recommended)

Create a `.env` file in the project root:

```bash
cp .env.example .env
```

Edit `.env` with your credentials:

```
PASSWORK_API_KEY=your-api-key
PASSWORK_HOST=https://your-passwork-instance.com/api/v1
PASSWORK_VAULT_ID=your-vault-id
```

Run tests:

```bash
go test
# or with verbose output
go test -v -cover
```

### Option 2: Using environment variables

```bash
export PASSWORK_API_KEY="api-key"
export PASSWORK_HOST="https://my-passwork-instance/api/v1"
export PASSWORK_VAULT_ID="vault-id"

go test -v -cover
```
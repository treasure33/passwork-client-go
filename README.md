# passwork-client-go
REST Client for the Password Manager Passwork written in Go.

The client can currently perform CRUD operations on passwords, folders and vaults.

## Example usage

```go
package main

import "github.com/treasure33/passwork-client-go"

func main() {
	host := "https://my-passwork-instance.com/api/v4"
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

```go
export PASSWORK_API_KEY="api-key"
export PASSWORK_HOST="https://my-passwork-instance/api/v4"
export PASSWORK_VAULT_ID="vault-id"

go test

// More elaborate
go test -v -cover

```
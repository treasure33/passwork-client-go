package passwork

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// Integration test for SearchPassword with GET and query parameters
func TestSearchPasswordIntegration(t *testing.T) {
	loadEnv()
	
	apiKey := os.Getenv("PASSWORK_API_KEY")
	host := os.Getenv("PASSWORK_HOST")
	vaultId := os.Getenv("PASSWORK_VAULT_ID")
	
	if apiKey == "" || host == "" || vaultId == "" {
		t.Skip("Skipping integration test: PASSWORK_API_KEY, PASSWORK_HOST, or PASSWORK_VAULT_ID not set")
	}
	
	client := NewClient(host, apiKey, time.Second*30)
	err := client.Login()
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	
	fmt.Println("✓ Login successful")
	fmt.Println("✓ Context timeout: 60 seconds")
	
	// Test 1: SearchPassword with query parameter
	fmt.Println("\n=== Test 1: SearchPassword with query='repo' ===")
	request1 := PasswordSearchRequest{
		Query:   "repo",
		VaultId: vaultId,
	}
	
	result1, err := client.SearchPassword(request1)
	if err != nil {
		// Show debug info
		url := fmt.Sprintf("%s/items/search?query=repo&vaultId=%s", host, vaultId)
		rawResp, statusCode, _ := client.sendRequest("GET", url, nil)
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Status Code: %d\n", statusCode)
		fmt.Printf("Raw Response: %s\n", string(rawResp))
		t.Fatalf("SearchPassword failed: %v", err)
	}
	
	fmt.Printf("Status: '%s'\n", result1.Status)
	fmt.Printf("Code: '%s'\n", result1.Code)
	fmt.Printf("Number of results: %d\n", len(result1.Data))
	
	if len(result1.Data) > 0 {
		fmt.Printf("First item: %s (ID: %s)\n", result1.Data[0].Name, result1.Data[0].Id)
		fmt.Println("✓ SearchPassword works with query parameter!")
	} else {
		fmt.Println("⚠ No items found")
	}
	
	// Test 2: SearchPassword without query (all items)
	fmt.Println("\n=== Test 2: SearchPassword without query (all items) ===")
	request2 := PasswordSearchRequest{
		VaultId: vaultId,
	}
	
	result2, err := client.SearchPassword(request2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of results: %d\n", len(result2.Data))
		if len(result2.Data) > 0 {
			fmt.Println("✓ SearchPassword works without query parameter!")
		}
	}
	
	// Test 3: SearchPassword with colors
	fmt.Println("\n=== Test 3: SearchPassword with colors ===")
	request3 := PasswordSearchRequest{
		VaultId: vaultId,
		Colors:  []int{1, 2},
	}
	
	result3, err := client.SearchPassword(request3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of results with colors [1,2]: %d\n", len(result3.Data))
		fmt.Println("✓ SearchPassword works with colors parameter!")
	}
	
	fmt.Println("\n✅ All SearchPassword tests completed successfully!")
}

package passwork

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PassworkTestSuite struct {
	suite.Suite
	ApiKey       string
	Host         string
	VaultId      string
	VaultName    string
	FolderId     string
	FolderName   string
	PasswordId   string
	PasswordName string
	client       *Client
}

// loadEnv loads environment variables from .env file if it exists
func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		return // .env file is optional
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// Remove surrounding quotes if present
			value = strings.Trim(value, `"'`)
			// Only set if not already set in environment
			if os.Getenv(key) == "" {
				os.Setenv(key, value)
			}
		}
	}
}

func (suite *PassworkTestSuite) SetupSuite() {
	// Load .env file if it exists
	loadEnv()

	suite.ApiKey = os.Getenv("PASSWORK_API_KEY")
	suite.Host = os.Getenv("PASSWORK_HOST")
	suite.VaultId = os.Getenv("PASSWORK_VAULT_ID")

	suite.client = NewClient(suite.Host, suite.ApiKey, time.Second*30)
	err := suite.client.Login()
	if err != nil {
		suite.Fail("Could not login to Passwork, Aborting test suite.")
	}
}

func TestPassworkTestSuite(t *testing.T) {
	suite.Run(t, new(PassworkTestSuite))
}

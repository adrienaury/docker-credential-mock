package internal

import (
	"github.com/docker/docker-credential-helpers/credentials"
)

// Store credentials in a local file
type YAMLStorage struct{}

// Add adds new credentials to the storage.
func (h YAMLStorage) Add(creds *credentials.Credentials) error {
	return nil
}

// Delete removes credentials from storage.
func (h YAMLStorage) Delete(serverURL string) error {
	return nil
}

// Get returns the username and secret to use for a given registry server URL.
func (h YAMLStorage) Get(serverURL string) (string, string, error) {
	return "", "", nil
}

// List returns the stored URLs and corresponding usernames.
func (h YAMLStorage) List() (map[string]string, error) {
	return nil, nil
}

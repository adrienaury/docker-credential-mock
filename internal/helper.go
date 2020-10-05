package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docker/docker-credential-helpers/credentials"
	"gopkg.in/yaml.v3"
)

// Version of the YAML strcuture
const Version string = "v1"

type YAMLCredentialsStore struct {
	Version         string            `yaml:"version"`
	CredentialsList []YAMLCredentials `yaml:"credentials,omitempty"`
}

type YAMLCredentials struct {
	ServerURL string `yaml:"serverURL"`
	Username  string `yaml:"username"`
	Secret    string `yaml:"secret"`
}

// Store credentials in a local file.
type YAMLStorage struct{}

// Add adds new credentials to the storage.
func (h YAMLStorage) Add(creds *credentials.Credentials) error {
	store, err := readFile()
	if err != nil {
		return err
	}

	yml := YAMLCredentials{
		ServerURL: creds.ServerURL,
		Username:  creds.Username,
		Secret:    creds.Secret,
	}

	added := false
	newList := []YAMLCredentials{}
	for _, credential := range store.CredentialsList {
		if credential.ServerURL != creds.ServerURL {
			newList = append(newList, credential)
		} else {
			newList = append(newList, yml)
			added = true
		}
	}
	if !added {
		newList = append(newList, yml)
	}
	store.CredentialsList = newList

	err = writeFile(store)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes credentials from storage.
func (h YAMLStorage) Delete(serverURL string) error {
	store, err := readFile()
	if err != nil {
		return err
	}

	newList := []YAMLCredentials{}
	for _, credential := range store.CredentialsList {
		if credential.ServerURL != serverURL {
			newList = append(newList, credential)
		}
	}

	store.CredentialsList = newList

	err = writeFile(store)
	if err != nil {
		return err
	}

	return nil
}

// Get returns the username and secret to use for a given registry server URL.
func (h YAMLStorage) Get(serverURL string) (string, string, error) {
	store, err := readFile()
	if err != nil {
		return "", "", err
	}

	for _, credential := range store.CredentialsList {
		if credential.ServerURL == serverURL {
			return credential.Username, credential.Secret, nil
		}
	}
	return "", "", nil
}

// List returns the stored URLs and corresponding usernames.
func (h YAMLStorage) List() (map[string]string, error) {
	store, err := readFile()
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	for _, credential := range store.CredentialsList {
		result[credential.ServerURL] = credential.Username
	}

	return result, nil
}

func readFile() (*YAMLCredentialsStore, error) {
	store := &YAMLCredentialsStore{
		Version: Version,
	}

	if _, err := os.Stat("credentials.yaml"); os.IsNotExist(err) {
		return store, nil
	}

	dat, err := ioutil.ReadFile("credentials.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dat, store)
	if err != nil {
		return nil, err
	}

	if store.Version != Version {
		return nil, fmt.Errorf("invalid version in ./credentials.yaml (%s)", store.Version)
	}

	return store, nil
}

func writeFile(list *YAMLCredentialsStore) error {
	out, err := yaml.Marshal(list)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("credentials.yaml", out, 0600)
	if err != nil {
		return err
	}

	return nil
}

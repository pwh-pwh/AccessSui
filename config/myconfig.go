package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	configFileName = "mnemonic.txt"
)

// SaveMnemonic saves the mnemonic to a local file.
func SaveMnemonic(mnemonic string) error {
	filePath := filepath.Join(".", configFileName) // Save in the current directory
	err := ioutil.WriteFile(filePath, []byte(mnemonic), 0644)
	if err != nil {
		return fmt.Errorf("failed to save mnemonic: %w", err)
	}
	return nil
}

// LoadMnemonic loads the mnemonic from a local file.
func LoadMnemonic() (string, error) {
	filePath := filepath.Join(".", configFileName) // Load from the current directory
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // File does not exist, no mnemonic saved yet
		}
		return "", fmt.Errorf("failed to load mnemonic: %w", err)
	}
	return string(content), nil
}

// ClearMnemonic clears the mnemonic from the local file.
func ClearMnemonic() error {
	filePath := filepath.Join(".", configFileName) // Delete from the current directory
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, nothing to clear
		}
		return fmt.Errorf("failed to clear mnemonic: %w", err)
	}
	return nil
}

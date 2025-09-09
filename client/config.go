package client

import (
	"fmt"

	"github.com/pwh-pwh/AccessSui/config"
)

func GetSuiClient() (*SuiContractClient, error) {
	mnemonic, err := config.LoadMnemonic()
	if err != nil {
		return nil, fmt.Errorf("failed to load mnemonic: %w", err)
	}
	contractClient, err := NewSuiContractClient(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create sui contract client: %w", err)
	}
	return contractClient, nil
}

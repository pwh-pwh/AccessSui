package AccessSui

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/pwh-pwh/AccessSui/config"
)

type SuiContractClient struct {
	cli *sui.SuiClient
	signerAccount *signer.SuiSigner
	privateKey string
}

func NewSuiContractClient(mnemonic string) (*SuiContractClient, error) {
	cli := sui.NewSuiClient(constant.BvTestnetEndpoint)

	signerAccount, err := signer.NewSignertWithMnemonic(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer account: %w", err)
	}
	
	return &SuiContractClient{
		cli: cli,
		signerAccount: signerAccount,
		privateKey: signerAccount.PriKey,
	}, nil
}

// PublishContent calls the 'publish_content' Move function.
func (s *SuiContractClient) PublishContent(
	ctx context.Context,
	creator string,
	uri string,
	contentTitle string,
	contentHash string,
	price string, // u64
	gasObjectId string,
	gasBudget string,
) (*models.SuiTransactionBlockResponse, error) {
	
	// Arguments for the Move function
	arguments := []interface{}{
		creator,
		[]byte(uri),
		[]byte(contentTitle),
		[]byte(contentHash),
		price,
	}

	rsp, err := s.cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          s.signerAccount.Address,
		PackageObjectId: config.PACKAGE_ID,
		Module:          "content",
		Function:        "publish_content",
		TypeArguments:   []interface{}{},
		Arguments:       arguments,
		Gas:             &gasObjectId,
		GasBudget:       gasBudget,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare MoveCall: %w", err)
	}

	rsp2, err := s.cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      s.privateKey,
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sign and execute transaction: %w", err)
	}

	utils.PrettyPrint(rsp2) // For debugging
	return rsp2, nil
}

// AddAccessLevel calls the 'add_access_level' Move function.
func (s *SuiContractClient) AddAccessLevel(
	ctx context.Context,
	contentObjectId string,
	levelName string,
	levelValue string, // u8
	gasObjectId string,
	gasBudget string,
) (*models.SuiTransactionBlockResponse, error) {

	// Arguments for the Move function
	arguments := []interface{}{
		contentObjectId,
		[]byte(levelName),
		levelValue,
	}

	rsp, err := s.cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          s.signerAccount.Address,
		PackageObjectId: config.PACKAGE_ID,
		Module:          "content",
		Function:        "add_access_level",
		TypeArguments:   []interface{}{},
		Arguments:       arguments,
		Gas:             &gasObjectId,
		GasBudget:       gasBudget,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare MoveCall: %w", err)
	}

	rsp2, err := s.cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      s.privateKey,
		Options: models.SuiTransactionBlockOptions{
			ShowInput:          true,
			ShowRawInput:       true,
			ShowEffects:        true,
			ShowEvents:         true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sign and execute transaction: %w", err)
	}

	utils.PrettyPrint(rsp2) // For debugging
	return rsp2, nil
}

// GetContentObject fetches the Content object details.
func (s *SuiContractClient) GetContentObject(ctx context.Context, objectId string) (*models.SuiGetObjectResponse, error) {
	rsp, err := s.cli.SuiGetObject(ctx, models.SuiGetObjectRequest{
		ObjectId: objectId,
		Options: models.SuiObjectDataOptions{
			ShowContent:             true,
			ShowDisplay:             true,
			ShowType:                true,
			ShowBcs:                 true,
			ShowOwner:               true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get content object: %w", err)
	}
	return rsp, nil
}
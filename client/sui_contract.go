package AccessSui

import (
	"context"
	"crypto/ed25519"
	"fmt"

	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
	"github.com/pwh-pwh/AccessSui/config"
)

type SuiContractClient struct {
	cli           sui.ISuiAPI
	signerAccount *signer.Signer
	privateKey    ed25519.PrivateKey
}

func NewSuiContractClient(mnemonic string) (*SuiContractClient, error) {
	cli := sui.NewSuiClient(constant.BvTestnetEndpoint)
	signerAccount, err := signer.NewSignertWithMnemonic(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer account: %w", err)
	}

	return &SuiContractClient{
		cli:           cli,
		signerAccount: signerAccount,
		privateKey:    signerAccount.PriKey,
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
	return &rsp2, nil
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
	return &rsp2, nil
}

// BuyAccessToken calls the 'buy_access_token' Move function.
func (s *SuiContractClient) BuyAccessToken(
	ctx context.Context,
	contentObjectId string,
	paymentCoinId string,
	durationSeconds string, // u64
	gasObjectId string,
	gasBudget string,
) (*models.SuiTransactionBlockResponse, error) {

	// Arguments for the Move function
	arguments := []interface{}{
		contentObjectId,
		paymentCoinId,
		durationSeconds,
		"0x6", // Clock object shared ID
	}

	rsp, err := s.cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          s.signerAccount.Address,
		PackageObjectId: config.PACKAGE_ID,
		Module:          "token",
		Function:        "buy_access_token",
		TypeArguments:   []interface{}{},
		Arguments:       arguments,
		Gas:             &gasObjectId,
		GasBudget:       gasBudget,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare MoveCall for BuyAccessToken: %w", err)
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
		return nil, fmt.Errorf("failed to sign and execute transaction for BuyAccessToken: %w", err)
	}

	utils.PrettyPrint(rsp2) // For debugging
	return &rsp2, nil
}

// TransferAccessToken calls the 'transfer_access_token' Move function.
func (s *SuiContractClient) TransferAccessToken(
	ctx context.Context,
	accessTokenId string,
	recipient string,
	paymentCoinId string,
	contentObjectId string,
	gasObjectId string,
	gasBudget string,
) (*models.SuiTransactionBlockResponse, error) {

	// Arguments for the Move function
	arguments := []interface{}{
		accessTokenId,
		recipient,
		paymentCoinId,
		contentObjectId,
	}

	rsp, err := s.cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          s.signerAccount.Address,
		PackageObjectId: config.PACKAGE_ID,
		Module:          "token",
		Function:        "transfer_access_token",
		TypeArguments:   []interface{}{},
		Arguments:       arguments,
		Gas:             &gasObjectId,
		GasBudget:       gasBudget,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare MoveCall for TransferAccessToken: %w", err)
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
		return nil, fmt.Errorf("failed to sign and execute transaction for TransferAccessToken: %w", err)
	}

	utils.PrettyPrint(rsp2) // For debugging
	return &rsp2, nil
}

// RevokeAccessToken calls the 'revoke_access_token' Move function.
func (s *SuiContractClient) RevokeAccessToken(
	ctx context.Context,
	accessTokenId string,
	contentObjectId string,
	gasObjectId string,
	gasBudget string,
) (*models.SuiTransactionBlockResponse, error) {

	// Arguments for the Move function
	arguments := []interface{}{
		accessTokenId,
		contentObjectId,
	}

	rsp, err := s.cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          s.signerAccount.Address,
		PackageObjectId: config.PACKAGE_ID,
		Module:          "token",
		Function:        "revoke_access_token",
		TypeArguments:   []interface{}{},
		Arguments:       arguments,
		Gas:             &gasObjectId,
		GasBudget:       gasBudget,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare MoveCall for RevokeAccessToken: %w", err)
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
		return nil, fmt.Errorf("failed to sign and execute transaction for RevokeAccessToken: %w", err)
	}

	utils.PrettyPrint(rsp2) // For debugging
	return &rsp2, nil
}

// GetAccessTokenObject fetches the AccessToken object details.
func (s *SuiContractClient) GetAccessTokenObject(ctx context.Context, objectId string) (*models.SuiObjectResponse, error) {
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
		return nil, fmt.Errorf("failed to get access token object: %w", err)
	}
	return &rsp, nil
}

// QueryEvents queries Sui events based on a filter.
func (s *SuiContractClient) QueryEvents(
	ctx context.Context,
	queryFilter models.SuiEventFilter,
	limit uint64,
	descendingOrder bool,
) (*models.PaginatedEventsResponse, error) {
	rsp, err := s.cli.SuiXQueryEvents(ctx, models.SuiXQueryEventsRequest{
		SuiEventFilter:  queryFilter,
		Limit:           limit,
		DescendingOrder: descendingOrder,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	return &rsp, nil
}

// SubscribeEvents subscribes to a stream of Sui events via WebSocket.
// The received events will be sent to the provided receiveMsgCh.
func (s *SuiContractClient) SubscribeEvents(
	ctx context.Context,
	queryFilter models.SuiEventFilter,
	receiveMsgCh chan models.SuiEventResponse,
) error {
	wsCli := sui.NewSuiWebsocketClient(constant.WssBvTestnetEndpoint) // Using testnet websocket endpoint
	err := wsCli.SubscribeEvent(ctx, models.SuiXSubscribeEventsRequest{
		SuiEventFilter: queryFilter,
	}, receiveMsgCh)
	if err != nil {
		return fmt.Errorf("failed to subscribe to events: %w", err)
	}
	return nil
}

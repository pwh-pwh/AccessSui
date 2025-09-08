package AccessSui

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

// In a real scenario, you would manage mnemonics securely.
// For testing purposes, we use a placeholder.
// DO NOT use this mnemonic in production.
const TEST_MNEMONIC = "YOUR_TEST_MNEMONIC_HERE"          // Replace with a test mnemonic from Sui Wallet
const TEST_GAS_OBJECT_ID = "YOUR_GAS_OBJECT_ID_HERE"     // Replace with a gas object ID from your test account
const TEST_CREATOR_ADDRESS = "YOUR_CREATOR_ADDRESS_HERE" // Replace with the address derived from your mnemonic
const TEST_PACKAGE_ID = "0x43572497e021468b14d91d4d17e86d960d3f30441efb60290b1dda50868be0a8"

func TestNewSuiContractClient(t *testing.T) {
	if TEST_MNEMONIC == "YOUR_TEST_MNEMONIC_HERE" || TEST_GAS_OBJECT_ID == "YOUR_GAS_OBJECT_ID_HERE" || TEST_CREATOR_ADDRESS == "YOUR_CREATOR_ADDRESS_HERE" {
		t.Skip("Skipping TestNewSuiContractClient: Please replace placeholder constants in sui_contract_test.go")
	}

	client, err := NewSuiContractClient(TEST_MNEMONIC)
	if err != nil {
		t.Fatalf("NewSuiContractClient failed: %v", err)
	}
	if client == nil {
		t.Fatal("NewSuiContractClient returned a nil client")
	}
	if client.signerAccount.Address != TEST_CREATOR_ADDRESS {
		t.Errorf("Expected signer address %s, got %s", TEST_CREATOR_ADDRESS, client.signerAccount.Address)
	}
	fmt.Printf("SuiContractClient created with address: %s\n", client.signerAccount.Address)
}

func TestPublishContent(t *testing.T) {

	client, err := NewSuiContractClient(TEST_MNEMONIC)
	if err != nil {
		t.Fatalf("NewSuiContractClient failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	uri := fmt.Sprintf("http://example.com/content/%d", time.Now().Unix())
	contentTitle := "Test Content Title"
	contentHash := "testcontenthash12345"
	price := "1000" // 1000 MIST

	fmt.Printf("Publishing content from creator: %s\n", TEST_CREATOR_ADDRESS)
	resp, err := client.PublishContent(ctx, TEST_CREATOR_ADDRESS, uri, contentTitle, contentHash, price, TEST_GAS_OBJECT_ID, "100000000")
	if err != nil {
		t.Fatalf("PublishContent failed: %v", err)
	}

	if resp.Digest == "" {
		t.Fatal("PublishContent did not return a transaction digest")
	}

	fmt.Printf("Published Content Transaction Digest: %s\n", resp.Digest)

}

func TestBuyAccessToken(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test TestBuyAccessToken. Set RUN_INTEGRATION_TESTS=true to enable.")
	}
	if TEST_MNEMONIC == "YOUR_TEST_MNEMONIC_HERE" || TEST_GAS_OBJECT_ID == "YOUR_GAS_OBJECT_ID_HERE" || TEST_CREATOR_ADDRESS == "YOUR_CREATOR_ADDRESS_HERE" {
		t.Skip("Skipping TestBuyAccessToken: Please replace placeholder constants in sui_contract_test.go")
	}

	// This test requires an existing Content object and a coin object to pay with.
	const CONTENT_OBJECT_ID_FOR_TOKEN = "YOUR_CONTENT_OBJECT_ID_FOR_TOKEN_HERE" // Replace with an actual Content object ID
	const PAYMENT_COIN_OBJECT_ID = "YOUR_PAYMENT_COIN_OBJECT_ID_HERE"           // Replace with an actual coin object ID

	if CONTENT_OBJECT_ID_FOR_TOKEN == "YOUR_CONTENT_OBJECT_ID_FOR_TOKEN_HERE" || PAYMENT_COIN_OBJECT_ID == "YOUR_PAYMENT_COIN_OBJECT_ID_HERE" {
		t.Skip("Skipping TestBuyAccessToken: Please provide CONTENT_OBJECT_ID_FOR_TOKEN and PAYMENT_COIN_OBJECT_ID in sui_contract_test.go")
	}

	client, err := NewSuiContractClient(TEST_MNEMONIC)
	if err != nil {
		t.Fatalf("NewSuiContractClient failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	duration := "3600" // 1 hour

	fmt.Printf("Buying access token for content %s with payment coin %s\n", CONTENT_OBJECT_ID_FOR_TOKEN, PAYMENT_COIN_OBJECT_ID)
	resp, err := client.BuyAccessToken(ctx, CONTENT_OBJECT_ID_FOR_TOKEN, PAYMENT_COIN_OBJECT_ID, duration, TEST_GAS_OBJECT_ID, "100000000")
	if err != nil {
		t.Fatalf("BuyAccessToken failed: %v", err)
	}

	if resp.Digest == "" {
		t.Fatal("BuyAccessToken did not return a transaction digest")
	}

	fmt.Printf("Buy Access Token Transaction Digest: %s\n", resp.Digest)

	time.Sleep(5 * time.Second) // Give some time for the transaction to be processed

}

func TestAddAccessLevel(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test TestAddAccessLevel. Set RUN_INTEGRATION_TESTS=true to enable.")
	}
	if TEST_MNEMONIC == "YOUR_TEST_MNEMONIC_HERE" || TEST_GAS_OBJECT_ID == "YOUR_GAS_OBJECT_ID_HERE" || TEST_CREATOR_ADDRESS == "YOUR_CREATOR_ADDRESS_HERE" {
		t.Skip("Skipping TestAddAccessLevel: Please replace placeholder constants in sui_contract_test.go")
	}

	// This test requires an existing Content object. For a proper integration test,
	// you might want to publish content first, or use a known pre-existing object ID.
	// For simplicity, we assume an object ID is provided here or published in the previous test.
	const EXISTING_CONTENT_OBJECT_ID = "YOUR_EXISTING_CONTENT_OBJECT_ID_HERE" // Replace with an actual Content object ID you own

	if EXISTING_CONTENT_OBJECT_ID == "YOUR_EXISTING_CONTENT_OBJECT_ID_HERE" {
		t.Skip("Skipping TestAddAccessLevel: Please provide an EXISTING_CONTENT_OBJECT_ID in sui_contract_test.go")
	}

	client, err := NewSuiContractClient(TEST_MNEMONIC)
	if err != nil {
		t.Fatalf("NewSuiContractClient failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	levelName := "premium"
	levelValue := "2"

	fmt.Printf("Adding access level '%s' to content object %s\n", levelName, EXISTING_CONTENT_OBJECT_ID)
	resp, err := client.AddAccessLevel(ctx, EXISTING_CONTENT_OBJECT_ID, levelName, levelValue, TEST_GAS_OBJECT_ID, "100000000")
	if err != nil {
		t.Fatalf("AddAccessLevel failed: %v", err)
	}

	if resp.Digest == "" {
		t.Fatal("AddAccessLevel did not return a transaction digest")
	}

	fmt.Printf("Add Access Level Transaction Digest: %s\n", resp.Digest)

	// Optionally, fetch the object to verify the access level was added
	time.Sleep(5 * time.Second)

	if err != nil {
		t.Fatalf("GetContentObject after AddAccessLevel failed: %v", err)
	}

}

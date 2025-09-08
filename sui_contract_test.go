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
const TEST_MNEMONIC = "YOUR_TEST_MNEMONIC_HERE" // Replace with a test mnemonic from Sui Wallet
const TEST_GAS_OBJECT_ID = "YOUR_GAS_OBJECT_ID_HERE" // Replace with a gas object ID from your test account
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
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test TestPublishContent. Set RUN_INTEGRATION_TESTS=true to enable.")
	}
	if TEST_MNEMONIC == "YOUR_TEST_MNEMONIC_HERE" || TEST_GAS_OBJECT_ID == "YOUR_GAS_OBJECT_ID_HERE" || TEST_CREATOR_ADDRESS == "YOUR_CREATOR_ADDRESS_HERE" {
		t.Skip("Skipping TestPublishContent: Please replace placeholder constants in sui_contract_test.go")
	}

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

	if resp.Error != nil {
		t.Fatalf("PublishContent transaction failed with error: %s", resp.Error.Message)
	}
	if resp.Digest == "" {
		t.Fatal("PublishContent did not return a transaction digest")
	}

	fmt.Printf("Published Content Transaction Digest: %s\n", resp.Digest)

	// Optionally, fetch the object to verify
	time.Sleep(5 * time.Second) // Give some time for the transaction to be processed
	if resp.ObjectChanges == nil || len(*resp.ObjectChanges) == 0 {
		t.Fatalf("No object changes found in transaction response. Cannot verify created object.")
	}

	var newContentObjectId string
	for _, change := range *resp.ObjectChanges {
		if change.Type == "created" && change.ObjectType == "0x" + TEST_PACKAGE_ID[2:] + "::content::Content" {
			newContentObjectId = change.ObjectId
			break
		}
	}

	if newContentObjectId == "" {
		t.Fatal("Failed to find new Content object ID in transaction response")
	}

	fmt.Printf("New Content Object ID: %s\n", newContentObjectId)

	contentObj, err := client.GetContentObject(ctx, newContentObjectId)
	if err != nil {
		t.Fatalf("GetContentObject failed: %v", err)
	}

	if contentObj.Data == nil {
		t.Fatalf("Content object data is nil for ID: %s", newContentObjectId)
	}

	// Further assertions could be made on contentObj.Data.Content.Fields if decoded
	fmt.Printf("Verified Content Object: %+v\n", contentObj.Data)
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

	if resp.Error != nil {
		t.Fatalf("AddAccessLevel transaction failed with error: %s", resp.Error.Message)
	}
	if resp.Digest == "" {
		t.Fatal("AddAccessLevel did not return a transaction digest")
	}

	fmt.Printf("Add Access Level Transaction Digest: %s\n", resp.Digest)

	// Optionally, fetch the object to verify the access level was added
	time.Sleep(5 * time.Second)
	contentObj, err := client.GetContentObject(ctx, EXISTING_CONTENT_OBJECT_ID)
	if err != nil {
		t.Fatalf("GetContentObject after AddAccessLevel failed: %v", err)
	}

	if contentObj.Data == nil {
		t.Fatalf("Content object data is nil after AddAccessLevel for ID: %s", EXISTING_CONTENT_OBJECT_ID)
	}

	// Verification of access level addition would require deserializing the `access_levels` table,
	// which is complex with the current SDK's raw object data.
	// For now, we rely on the transaction being successful.
	fmt.Printf("Verified Content Object after adding access level: %+v\n", contentObj.Data)
}
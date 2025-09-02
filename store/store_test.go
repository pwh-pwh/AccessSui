package store

import (
	"testing"
)

func TestRecData(t *testing.T) {
	data, err := RecData("EP0zTtQmtH3SBbJCcjRC_7FF9z-l96jxT8EZ2evFGZg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestStoreData(t *testing.T) {
	blodId, err := StoreData([]byte("test store data"))
	if err != nil {
		t.Errorf("StoreData() error = %v", err)
	}
	t.Log(blodId)
}

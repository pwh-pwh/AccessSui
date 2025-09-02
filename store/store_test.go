package store

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecData(t *testing.T) {
	data, err := RecData("EP0zTtQmtH3SBbJCcjRC_7FF9z-l96jxT8EZ2evFGZg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	assert.Equal(t, data, []byte("test store data"))
}

func TestStoreData(t *testing.T) {
	blodId, err := StoreData([]byte("test store data"))
	if err != nil {
		t.Errorf("StoreData() error = %v", err)
	}
	t.Log(blodId)
}

func TestRecDataWithKey(t *testing.T) {
	key := bytes.Repeat([]byte{1, 2, 3, 4}, 8)
	data, err := RecDataWithKey("8gzU7K11HRmWatoJSL-8yzZ9DmUTKs621IJN3u4zOig", key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	assert.Equal(t, data, []byte("test store data with key"))
}

func TestStoreDataWithKey(t *testing.T) {
	key := bytes.Repeat([]byte{1, 2, 3, 4}, 8)
	blodId, err := StoreDataWithKey([]byte("test store data with key"), key)
	if err != nil {
		t.Errorf("StoreDataWithKey() error = %v", err)
	}
	t.Log(blodId)
}

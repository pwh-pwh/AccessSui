package store

import (
	"fmt"

	walrus "github.com/namihq/walrus-go"
)

const (
	AGGURL = "https://aggregator.testnet.walrus.atalma.io"
	PUBURL = "https://publisher.walrus-01.tududes.com"
)

var WALRUS_CLIENT = walrus.NewClient(
	walrus.WithAggregatorURLs([]string{AGGURL}),
	walrus.WithPublisherURLs([]string{PUBURL}),
)

func StoreData(data []byte) (blodId string, err error) {
	resp, err := WALRUS_CLIENT.Store(data, &walrus.StoreOptions{Epochs: 1})
	if err != nil {
		return
	}
	if resp.NewlyCreated != nil {
		blobID := resp.NewlyCreated.BlobObject.BlobID
		fmt.Printf("Stored new blob ID: %s with cost: %d\n",
			blobID, resp.NewlyCreated.Cost)
	} else if resp.AlreadyCertified != nil {
		blobID := resp.AlreadyCertified.BlobID
		fmt.Printf("Blob already exists with ID: %s, end epoch: %d\n",
			blobID, resp.AlreadyCertified.EndEpoch)
	}
	return
}

func RecData(blodId string) (data []byte, err error) {
	data, err = WALRUS_CLIENT.Read(blodId, nil)
	return
}

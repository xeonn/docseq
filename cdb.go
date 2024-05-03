package docseq

import (
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
)

var driver = "couch"

var DbHost string
var DbPort int = -1
var DbUser string
var DbPassword string
var DbUsessl bool = false

// Database name in Couchdb. Override to use different database
var TableName string = "sequence"

// Return dbUri for use in http request. (Required couchdb partitioned database)
func DbUri(ctx context.Context) (string, error) {
	if DbHost == "" {
		return "", fmt.Errorf("DbHost not defined")
	}
	if DbPort < 0 {
		return "", fmt.Errorf("DbPort not defined")
	}
	if DbUser == "" {
		return "", fmt.Errorf("DbUser not defined")
	}
	if DbPassword == "" {
		return "", fmt.Errorf("DbPassword not defined")
	}

	protocol := "http"

	if DbUsessl {
		protocol = "https"
	}

	return protocol + "://" + DbUser + ":" + DbPassword + "@" + DbHost + ":" + fmt.Sprintf("%d", DbPort) + "/", nil

}

// Return new client for each call. The callee must call client.Close()
// after use
func NewDb(ctx context.Context) (*kivik.Client, error) {

	uri, err := DbUri(ctx)
	if err != nil {
		return nil, err
	}

	client, err := kivik.New(driver, uri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

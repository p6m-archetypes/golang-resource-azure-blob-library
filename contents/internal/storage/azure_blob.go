package storage

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	appconfig "{{ module_path }}/internal/config"
)

var AzureBlobClient *azblob.Client

func InitAzureBlob(cfg appconfig.AzureBlobConfig) error {
	cred, err := azblob.NewSharedKeyCredential(cfg.AccountName, cfg.AccountKey)
	if err != nil {
		return fmt.Errorf("azure-blob: credential: %w", err)
	}
	client, err := azblob.NewClientWithSharedKeyCredential(cfg.Endpoint, cred, nil)
	if err != nil {
		return fmt.Errorf("azure-blob: client: %w", err)
	}
	AzureBlobClient = client
	return nil
}

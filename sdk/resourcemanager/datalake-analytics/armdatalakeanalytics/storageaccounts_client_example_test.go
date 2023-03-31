//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatalakeanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_ListByAccount.json
func ExampleStorageAccountsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAccountsClient().NewListByAccountPager("contosorg", "contosoadla", &armdatalakeanalytics.StorageAccountsClientListByAccountOptions{Filter: to.Ptr("test_filter"),
		Top:     to.Ptr[int32](1),
		Skip:    to.Ptr[int32](1),
		Select:  to.Ptr("test_select"),
		Orderby: to.Ptr("test_orderby"),
		Count:   to.Ptr(false),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.StorageAccountInformationListResult = armdatalakeanalytics.StorageAccountInformationListResult{
		// 	Value: []*armdatalakeanalytics.StorageAccountInformation{
		// 		{
		// 			Name: to.Ptr("test_storage"),
		// 			Type: to.Ptr("test_type"),
		// 			ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			Properties: &armdatalakeanalytics.StorageAccountInformationProperties{
		// 				Suffix: to.Ptr("test_suffix"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_Add.json
func ExampleStorageAccountsClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStorageAccountsClient().Add(ctx, "contosorg", "contosoadla", "test_storage", armdatalakeanalytics.AddStorageAccountParameters{
		Properties: &armdatalakeanalytics.AddStorageAccountProperties{
			AccessKey: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
			Suffix:    to.Ptr("test_suffix"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_Get.json
func ExampleStorageAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageAccountsClient().Get(ctx, "contosorg", "contosoadla", "test_storage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageAccountInformation = armdatalakeanalytics.StorageAccountInformation{
	// 	Name: to.Ptr("test_storage"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakeanalytics.StorageAccountInformationProperties{
	// 		Suffix: to.Ptr("test_suffix"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_Update.json
func ExampleStorageAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStorageAccountsClient().Update(ctx, "contosorg", "contosoadla", "test_storage", &armdatalakeanalytics.StorageAccountsClientUpdateOptions{Parameters: &armdatalakeanalytics.UpdateStorageAccountParameters{
		Properties: &armdatalakeanalytics.UpdateStorageAccountProperties{
			AccessKey: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
			Suffix:    to.Ptr("test_suffix"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_Delete.json
func ExampleStorageAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStorageAccountsClient().Delete(ctx, "contosorg", "contosoadla", "test_storage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_ListStorageContainers.json
func ExampleStorageAccountsClient_NewListStorageContainersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAccountsClient().NewListStorageContainersPager("contosorg", "contosoadla", "test_storage", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.StorageContainerListResult = armdatalakeanalytics.StorageContainerListResult{
		// 	Value: []*armdatalakeanalytics.StorageContainer{
		// 		{
		// 			Name: to.Ptr("test_storage"),
		// 			Type: to.Ptr("test_type"),
		// 			ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			Properties: &armdatalakeanalytics.StorageContainerProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T13:21:56.6819037-07:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_GetStorageContainer.json
func ExampleStorageAccountsClient_GetStorageContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageAccountsClient().GetStorageContainer(ctx, "contosorg", "contosoadla", "test_storage", "test_container", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageContainer = armdatalakeanalytics.StorageContainer{
	// 	Name: to.Ptr("test_storage"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakeanalytics.StorageContainerProperties{
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-14T13:21:56.6819037-07:00"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/StorageAccounts_ListSasTokens.json
func ExampleStorageAccountsClient_NewListSasTokensPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakeanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAccountsClient().NewListSasTokensPager("contosorg", "contosoadla", "test_storage", "test_container", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SasTokenInformationListResult = armdatalakeanalytics.SasTokenInformationListResult{
		// 	Value: []*armdatalakeanalytics.SasTokenInformation{
		// 		{
		// 			AccessToken: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346"),
		// 	}},
		// }
	}
}
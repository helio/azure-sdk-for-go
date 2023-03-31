//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/GetRegisteredPrefix.json
func ExampleRegisteredPrefixesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegisteredPrefixesClient().Get(ctx, "rgName", "peeringName", "registeredPrefixName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegisteredPrefix = armpeering.RegisteredPrefix{
	// 	Name: to.Ptr("registeredPrefixName"),
	// 	Type: to.Ptr("Microsoft.Peering/registeredPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredPrefixes/registeredPrefixName"),
	// 	Properties: &armpeering.RegisteredPrefixProperties{
	// 		PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Prefix: to.Ptr("10.22.20.0/24"),
	// 		PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateVerified),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreateRegisteredPrefix.json
func ExampleRegisteredPrefixesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegisteredPrefixesClient().CreateOrUpdate(ctx, "rgName", "peeringName", "registeredPrefixName", armpeering.RegisteredPrefix{
		Properties: &armpeering.RegisteredPrefixProperties{
			Prefix: to.Ptr("10.22.20.0/24"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegisteredPrefix = armpeering.RegisteredPrefix{
	// 	Name: to.Ptr("registeredPrefixName"),
	// 	Type: to.Ptr("Microsoft.Peering/registeredPrefixes"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredPrefixes/registeredPrefixName"),
	// 	Properties: &armpeering.RegisteredPrefixProperties{
	// 		PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Prefix: to.Ptr("10.22.20.0/24"),
	// 		PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateVerified),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/DeleteRegisteredPrefix.json
func ExampleRegisteredPrefixesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRegisteredPrefixesClient().Delete(ctx, "rgName", "peeringName", "registeredPrefixName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListRegisteredPrefixesByPeering.json
func ExampleRegisteredPrefixesClient_NewListByPeeringPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegisteredPrefixesClient().NewListByPeeringPager("rgName", "peeringName", nil)
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
		// page.RegisteredPrefixListResult = armpeering.RegisteredPrefixListResult{
		// 	Value: []*armpeering.RegisteredPrefix{
		// 		{
		// 			Name: to.Ptr("registeredPrefixName0"),
		// 			Type: to.Ptr("Microsoft.Peering/registeredPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredPrefixes/registeredPrefixName0"),
		// 			Properties: &armpeering.RegisteredPrefixProperties{
		// 				PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Prefix: to.Ptr("10.22.20.0/24"),
		// 				PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateVerified),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("registeredPrefixName1"),
		// 			Type: to.Ptr("Microsoft.Peering/registeredPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredPrefixes/registeredPrefixName1"),
		// 			Properties: &armpeering.RegisteredPrefixProperties{
		// 				PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Prefix: to.Ptr("10.22.21.0/24"),
		// 				PrefixValidationState: to.Ptr(armpeering.PrefixValidationStateVerified),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("registeredPrefixName2"),
		// 			Type: to.Ptr("Microsoft.Peering/registeredPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredPrefixes/registeredPrefixName2"),
		// 			Properties: &armpeering.RegisteredPrefixProperties{
		// 				Prefix: to.Ptr("10.22.22.0/24"),
		// 				PrefixValidationState: to.Ptr(armpeering.PrefixValidationStatePending),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
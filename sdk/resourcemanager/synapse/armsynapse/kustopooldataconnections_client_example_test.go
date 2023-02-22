//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCheckNameAvailability.json
func ExampleKustoPoolDataConnectionsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "Kustodatabase8", armsynapse.DataConnectionCheckNameRequest{
		Name: to.Ptr("DataConnections8"),
		Type: to.Ptr("Microsoft.Synapse/workspaces/kustoPools/databases/dataConnections"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armsynapse.CheckNameResult{
	// 	Name: to.Ptr("DataConnections8"),
	// 	Message: to.Ptr("Name 'DataConnections8' is already taken. Please specify a different name."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armsynapse.ReasonAlreadyExists),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
func ExampleKustoPoolDataConnectionsClient_BeginDataConnectionValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDataConnectionValidation(ctx, "kustorptest", "kustorptest", "kustoclusterrptest4", "KustoDatabase8", armsynapse.DataConnectionValidation{
		DataConnectionName: to.Ptr("DataConnections8"),
		Properties: &armsynapse.EventHubDataConnection{
			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataConnectionValidationListResult = armsynapse.DataConnectionValidationListResult{
	// 	Value: []*armsynapse.DataConnectionValidationResult{
	// 		{
	// 			ErrorMessage: to.Ptr("Event hub's namespace does not exist"),
	// 		},
	// 		{
	// 			ErrorMessage: to.Ptr("Database does not exist"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsListByDatabase.json
func ExampleKustoPoolDataConnectionsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDatabasePager("kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", nil)
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
		// page.DataConnectionListResult = armsynapse.DataConnectionListResult{
		// 	Value: []armsynapse.DataConnectionClassification{
		// 		&armsynapse.EventHubDataConnection{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8/KustoDataConnection8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8/DataConnections/KustoDataConnection8"),
		// 			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.EventHubConnectionProperties{
		// 				ConsumerGroup: to.Ptr("testConsumerGroup1"),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
		// 			},
		// 		},
		// 		&armsynapse.EventHubDataConnection{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase9/KustoDataConnection9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase9/DataConnections/KustoDataConnection9"),
		// 			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.EventHubConnectionProperties{
		// 				ConsumerGroup: to.Ptr("testConsumerGroup2"),
		// 				EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns2/eventhubs/eventhubTest2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsGet.json
func ExampleKustoPoolDataConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", "DataConnections8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsynapse.KustoPoolDataConnectionsClientGetResponse{
	// 	                            DataConnectionClassification: &armsynapse.EventHubDataConnection{
	// 		Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8/DataConnections8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8/DataConnections/DataConnections8"),
	// 		Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armsynapse.EventHubConnectionProperties{
	// 			ConsumerGroup: to.Ptr("testConsumerGroup1"),
	// 			EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCreateOrUpdate.json
func ExampleKustoPoolDataConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", "DataConnections8", &armsynapse.EventHubDataConnection{
		Kind:     to.Ptr(armsynapse.DataConnectionKindEventHub),
		Location: to.Ptr("westus"),
		Properties: &armsynapse.EventHubConnectionProperties{
			ConsumerGroup:      to.Ptr("testConsumerGroup1"),
			EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsynapse.KustoPoolDataConnectionsClientCreateOrUpdateResponse{
	// 	                            DataConnectionClassification: &armsynapse.EventHubDataConnection{
	// 		Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8/DataConnections8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8/DataConnections/DataConnections8"),
	// 		Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armsynapse.EventHubConnectionProperties{
	// 			ConsumerGroup: to.Ptr("testConsumerGroup1"),
	// 			EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsUpdate.json
func ExampleKustoPoolDataConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", "DataConnections8", &armsynapse.EventHubDataConnection{
		Kind:     to.Ptr(armsynapse.DataConnectionKindEventHub),
		Location: to.Ptr("westus"),
		Properties: &armsynapse.EventHubConnectionProperties{
			ConsumerGroup:      to.Ptr("testConsumerGroup1"),
			EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsynapse.KustoPoolDataConnectionsClientUpdateResponse{
	// 	                            DataConnectionClassification: &armsynapse.EventHubDataConnection{
	// 		Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8/DataConnections8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/DataConnections"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/synapseWorkspaceName/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8/DataConnections/DataConnections8"),
	// 		Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armsynapse.EventHubConnectionProperties{
	// 			ConsumerGroup: to.Ptr("testConsumerGroup1"),
	// 			EventHubResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1"),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsDelete.json
func ExampleKustoPoolDataConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", "kustoeventhubconnection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsphere

// CatalogsClientCountDevicesResponse contains the response from method CatalogsClient.CountDevices.
type CatalogsClientCountDevicesResponse struct {
	// Response to the action call for count devices in a catalog.
	CountDevicesResponse
}

// CatalogsClientCreateOrUpdateResponse contains the response from method CatalogsClient.BeginCreateOrUpdate.
type CatalogsClientCreateOrUpdateResponse struct {
	// An Azure Sphere catalog
	Catalog
}

// CatalogsClientDeleteResponse contains the response from method CatalogsClient.BeginDelete.
type CatalogsClientDeleteResponse struct {
	// placeholder for future response values
}

// CatalogsClientGetResponse contains the response from method CatalogsClient.Get.
type CatalogsClientGetResponse struct {
	// An Azure Sphere catalog
	Catalog
}

// CatalogsClientListByResourceGroupResponse contains the response from method CatalogsClient.NewListByResourceGroupPager.
type CatalogsClientListByResourceGroupResponse struct {
	// The response of a Catalog list operation.
	CatalogListResult
}

// CatalogsClientListBySubscriptionResponse contains the response from method CatalogsClient.NewListBySubscriptionPager.
type CatalogsClientListBySubscriptionResponse struct {
	// The response of a Catalog list operation.
	CatalogListResult
}

// CatalogsClientListDeploymentsResponse contains the response from method CatalogsClient.NewListDeploymentsPager.
type CatalogsClientListDeploymentsResponse struct {
	// The response of a Deployment list operation.
	DeploymentListResult
}

// CatalogsClientListDeviceGroupsResponse contains the response from method CatalogsClient.NewListDeviceGroupsPager.
type CatalogsClientListDeviceGroupsResponse struct {
	// The response of a DeviceGroup list operation.
	DeviceGroupListResult
}

// CatalogsClientListDeviceInsightsResponse contains the response from method CatalogsClient.NewListDeviceInsightsPager.
type CatalogsClientListDeviceInsightsResponse struct {
	// Paged collection of DeviceInsight items
	PagedDeviceInsight
}

// CatalogsClientListDevicesResponse contains the response from method CatalogsClient.NewListDevicesPager.
type CatalogsClientListDevicesResponse struct {
	// The response of a Device list operation.
	DeviceListResult
}

// CatalogsClientUpdateResponse contains the response from method CatalogsClient.Update.
type CatalogsClientUpdateResponse struct {
	// An Azure Sphere catalog
	Catalog
}

// CatalogsClientUploadImageResponse contains the response from method CatalogsClient.BeginUploadImage.
type CatalogsClientUploadImageResponse struct {
	// placeholder for future response values
}

// CertificatesClientGetResponse contains the response from method CertificatesClient.Get.
type CertificatesClientGetResponse struct {
	// An certificate resource belonging to a catalog resource.
	Certificate
}

// CertificatesClientListByCatalogResponse contains the response from method CertificatesClient.NewListByCatalogPager.
type CertificatesClientListByCatalogResponse struct {
	// The response of a Certificate list operation.
	CertificateListResult
}

// CertificatesClientRetrieveCertChainResponse contains the response from method CertificatesClient.RetrieveCertChain.
type CertificatesClientRetrieveCertChainResponse struct {
	// The certificate chain response.
	CertificateChainResponse
}

// CertificatesClientRetrieveProofOfPossessionNonceResponse contains the response from method CertificatesClient.RetrieveProofOfPossessionNonce.
type CertificatesClientRetrieveProofOfPossessionNonceResponse struct {
	// Result of the action to generate a proof of possession nonce
	ProofOfPossessionNonceResponse
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.BeginCreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	// An deployment resource belonging to a device group resource.
	Deployment
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.BeginDelete.
type DeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	// An deployment resource belonging to a device group resource.
	Deployment
}

// DeploymentsClientListByDeviceGroupResponse contains the response from method DeploymentsClient.NewListByDeviceGroupPager.
type DeploymentsClientListByDeviceGroupResponse struct {
	// The response of a Deployment list operation.
	DeploymentListResult
}

// DeviceGroupsClientClaimDevicesResponse contains the response from method DeviceGroupsClient.BeginClaimDevices.
type DeviceGroupsClientClaimDevicesResponse struct {
	// placeholder for future response values
}

// DeviceGroupsClientCountDevicesResponse contains the response from method DeviceGroupsClient.CountDevices.
type DeviceGroupsClientCountDevicesResponse struct {
	// Response to the action call for count devices in a catalog.
	CountDevicesResponse
}

// DeviceGroupsClientCreateOrUpdateResponse contains the response from method DeviceGroupsClient.BeginCreateOrUpdate.
type DeviceGroupsClientCreateOrUpdateResponse struct {
	// An device group resource belonging to a product resource.
	DeviceGroup
}

// DeviceGroupsClientDeleteResponse contains the response from method DeviceGroupsClient.BeginDelete.
type DeviceGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeviceGroupsClientGetResponse contains the response from method DeviceGroupsClient.Get.
type DeviceGroupsClientGetResponse struct {
	// An device group resource belonging to a product resource.
	DeviceGroup
}

// DeviceGroupsClientListByProductResponse contains the response from method DeviceGroupsClient.NewListByProductPager.
type DeviceGroupsClientListByProductResponse struct {
	// The response of a DeviceGroup list operation.
	DeviceGroupListResult
}

// DeviceGroupsClientUpdateResponse contains the response from method DeviceGroupsClient.BeginUpdate.
type DeviceGroupsClientUpdateResponse struct {
	// An device group resource belonging to a product resource.
	DeviceGroup
}

// DevicesClientCreateOrUpdateResponse contains the response from method DevicesClient.BeginCreateOrUpdate.
type DevicesClientCreateOrUpdateResponse struct {
	// An device resource belonging to a device group resource.
	Device
}

// DevicesClientDeleteResponse contains the response from method DevicesClient.BeginDelete.
type DevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DevicesClientGenerateCapabilityImageResponse contains the response from method DevicesClient.BeginGenerateCapabilityImage.
type DevicesClientGenerateCapabilityImageResponse struct {
	// Signed device capability image response
	SignedCapabilityImageResponse
}

// DevicesClientGetResponse contains the response from method DevicesClient.Get.
type DevicesClientGetResponse struct {
	// An device resource belonging to a device group resource.
	Device
}

// DevicesClientListByDeviceGroupResponse contains the response from method DevicesClient.NewListByDeviceGroupPager.
type DevicesClientListByDeviceGroupResponse struct {
	// The response of a Device list operation.
	DeviceListResult
}

// DevicesClientUpdateResponse contains the response from method DevicesClient.BeginUpdate.
type DevicesClientUpdateResponse struct {
	// An device resource belonging to a device group resource.
	Device
}

// ImagesClientCreateOrUpdateResponse contains the response from method ImagesClient.BeginCreateOrUpdate.
type ImagesClientCreateOrUpdateResponse struct {
	// An image resource belonging to a catalog resource.
	Image
}

// ImagesClientDeleteResponse contains the response from method ImagesClient.BeginDelete.
type ImagesClientDeleteResponse struct {
	// placeholder for future response values
}

// ImagesClientGetResponse contains the response from method ImagesClient.Get.
type ImagesClientGetResponse struct {
	// An image resource belonging to a catalog resource.
	Image
}

// ImagesClientListByCatalogResponse contains the response from method ImagesClient.NewListByCatalogPager.
type ImagesClientListByCatalogResponse struct {
	// The response of a Image list operation.
	ImageListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ProductsClientCountDevicesResponse contains the response from method ProductsClient.CountDevices.
type ProductsClientCountDevicesResponse struct {
	// Response to the action call for count devices in a catalog.
	CountDevicesResponse
}

// ProductsClientCreateOrUpdateResponse contains the response from method ProductsClient.BeginCreateOrUpdate.
type ProductsClientCreateOrUpdateResponse struct {
	// An product resource belonging to a catalog resource.
	Product
}

// ProductsClientDeleteResponse contains the response from method ProductsClient.BeginDelete.
type ProductsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProductsClientGenerateDefaultDeviceGroupsResponse contains the response from method ProductsClient.NewGenerateDefaultDeviceGroupsPager.
type ProductsClientGenerateDefaultDeviceGroupsResponse struct {
	// The response of a DeviceGroup list operation.
	DeviceGroupListResult
}

// ProductsClientGetResponse contains the response from method ProductsClient.Get.
type ProductsClientGetResponse struct {
	// An product resource belonging to a catalog resource.
	Product
}

// ProductsClientListByCatalogResponse contains the response from method ProductsClient.NewListByCatalogPager.
type ProductsClientListByCatalogResponse struct {
	// The response of a Product list operation.
	ProductListResult
}

// ProductsClientUpdateResponse contains the response from method ProductsClient.BeginUpdate.
type ProductsClientUpdateResponse struct {
	// An product resource belonging to a catalog resource.
	Product
}
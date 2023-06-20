//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservicefleet

const (
	moduleName    = "armcontainerservicefleet"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// FleetMemberProvisioningState - The provisioning state of the last accepted operation.
type FleetMemberProvisioningState string

const (
	// FleetMemberProvisioningStateCanceled - Resource creation was canceled.
	FleetMemberProvisioningStateCanceled FleetMemberProvisioningState = "Canceled"
	// FleetMemberProvisioningStateFailed - Resource creation failed.
	FleetMemberProvisioningStateFailed  FleetMemberProvisioningState = "Failed"
	FleetMemberProvisioningStateJoining FleetMemberProvisioningState = "Joining"
	FleetMemberProvisioningStateLeaving FleetMemberProvisioningState = "Leaving"
	// FleetMemberProvisioningStateSucceeded - Resource has been created.
	FleetMemberProvisioningStateSucceeded FleetMemberProvisioningState = "Succeeded"
	FleetMemberProvisioningStateUpdating  FleetMemberProvisioningState = "Updating"
)

// PossibleFleetMemberProvisioningStateValues returns the possible values for the FleetMemberProvisioningState const type.
func PossibleFleetMemberProvisioningStateValues() []FleetMemberProvisioningState {
	return []FleetMemberProvisioningState{
		FleetMemberProvisioningStateCanceled,
		FleetMemberProvisioningStateFailed,
		FleetMemberProvisioningStateJoining,
		FleetMemberProvisioningStateLeaving,
		FleetMemberProvisioningStateSucceeded,
		FleetMemberProvisioningStateUpdating,
	}
}

// FleetProvisioningState - The provisioning state of the last accepted operation.
type FleetProvisioningState string

const (
	// FleetProvisioningStateCanceled - Resource creation was canceled.
	FleetProvisioningStateCanceled FleetProvisioningState = "Canceled"
	FleetProvisioningStateCreating FleetProvisioningState = "Creating"
	FleetProvisioningStateDeleting FleetProvisioningState = "Deleting"
	// FleetProvisioningStateFailed - Resource creation failed.
	FleetProvisioningStateFailed FleetProvisioningState = "Failed"
	// FleetProvisioningStateSucceeded - Resource has been created.
	FleetProvisioningStateSucceeded FleetProvisioningState = "Succeeded"
	FleetProvisioningStateUpdating  FleetProvisioningState = "Updating"
)

// PossibleFleetProvisioningStateValues returns the possible values for the FleetProvisioningState const type.
func PossibleFleetProvisioningStateValues() []FleetProvisioningState {
	return []FleetProvisioningState{
		FleetProvisioningStateCanceled,
		FleetProvisioningStateCreating,
		FleetProvisioningStateDeleting,
		FleetProvisioningStateFailed,
		FleetProvisioningStateSucceeded,
		FleetProvisioningStateUpdating,
	}
}

// ManagedClusterUpgradeType - The type of upgrade to perform when targeting ManagedClusters.
type ManagedClusterUpgradeType string

const (
	// ManagedClusterUpgradeTypeFull - Full upgrades the control plane and all agent pools of the target ManagedClusters.
	ManagedClusterUpgradeTypeFull ManagedClusterUpgradeType = "Full"
	// ManagedClusterUpgradeTypeNodeImageOnly - NodeImageOnly upgrades only the node images of the target ManagedClusters.
	ManagedClusterUpgradeTypeNodeImageOnly ManagedClusterUpgradeType = "NodeImageOnly"
)

// PossibleManagedClusterUpgradeTypeValues returns the possible values for the ManagedClusterUpgradeType const type.
func PossibleManagedClusterUpgradeTypeValues() []ManagedClusterUpgradeType {
	return []ManagedClusterUpgradeType{
		ManagedClusterUpgradeTypeFull,
		ManagedClusterUpgradeTypeNodeImageOnly,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// UpdateRunProvisioningState - The provisioning state of the UpdateRun resource.
type UpdateRunProvisioningState string

const (
	// UpdateRunProvisioningStateCanceled - Resource creation was canceled.
	UpdateRunProvisioningStateCanceled UpdateRunProvisioningState = "Canceled"
	// UpdateRunProvisioningStateFailed - Resource creation failed.
	UpdateRunProvisioningStateFailed UpdateRunProvisioningState = "Failed"
	// UpdateRunProvisioningStateSucceeded - Resource has been created.
	UpdateRunProvisioningStateSucceeded UpdateRunProvisioningState = "Succeeded"
)

// PossibleUpdateRunProvisioningStateValues returns the possible values for the UpdateRunProvisioningState const type.
func PossibleUpdateRunProvisioningStateValues() []UpdateRunProvisioningState {
	return []UpdateRunProvisioningState{
		UpdateRunProvisioningStateCanceled,
		UpdateRunProvisioningStateFailed,
		UpdateRunProvisioningStateSucceeded,
	}
}

// UpdateState - The state of the UpdateRun, UpdateStage, UpdateGroup, or MemberUpdate.
type UpdateState string

const (
	UpdateStateCompleted  UpdateState = "Completed"
	UpdateStateFailed     UpdateState = "Failed"
	UpdateStateNotStarted UpdateState = "NotStarted"
	UpdateStateRunning    UpdateState = "Running"
	UpdateStateStopped    UpdateState = "Stopped"
	UpdateStateStopping   UpdateState = "Stopping"
)

// PossibleUpdateStateValues returns the possible values for the UpdateState const type.
func PossibleUpdateStateValues() []UpdateState {
	return []UpdateState{
		UpdateStateCompleted,
		UpdateStateFailed,
		UpdateStateNotStarted,
		UpdateStateRunning,
		UpdateStateStopped,
		UpdateStateStopping,
	}
}

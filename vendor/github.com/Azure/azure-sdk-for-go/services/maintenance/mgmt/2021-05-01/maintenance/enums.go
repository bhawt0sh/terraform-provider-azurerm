package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// ImpactType enumerates the values for impact type.
type ImpactType string

const (
	// ImpactTypeFreeze Pending updates can freeze network or disk io operation on resource.
	ImpactTypeFreeze ImpactType = "Freeze"
	// ImpactTypeNone Pending updates has no impact on resource.
	ImpactTypeNone ImpactType = "None"
	// ImpactTypeRedeploy Pending updates can redeploy resource.
	ImpactTypeRedeploy ImpactType = "Redeploy"
	// ImpactTypeRestart Pending updates can cause resource to restart.
	ImpactTypeRestart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns an array of possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{ImpactTypeFreeze, ImpactTypeNone, ImpactTypeRedeploy, ImpactTypeRestart}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// ScopeExtension This maintenance scope controls extension installation on VM/VMSS
	ScopeExtension Scope = "Extension"
	// ScopeHost This maintenance scope controls installation of azure platform updates i.e. services on
	// physical nodes hosting customer VMs.
	ScopeHost Scope = "Host"
	// ScopeInGuestPatch This maintenance scope controls installation of windows and linux packages on VM/VMSS
	ScopeInGuestPatch Scope = "InGuestPatch"
	// ScopeOSImage This maintenance scope controls os image installation on VM/VMSS
	ScopeOSImage Scope = "OSImage"
	// ScopeSQLDB This maintenance scope controls installation of SQL server platform updates.
	ScopeSQLDB Scope = "SQLDB"
	// ScopeSQLManagedInstance This maintenance scope controls installation of SQL managed instance platform
	// update.
	ScopeSQLManagedInstance Scope = "SQLManagedInstance"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{ScopeExtension, ScopeHost, ScopeInGuestPatch, ScopeOSImage, ScopeSQLDB, ScopeSQLManagedInstance}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// UpdateStatusCompleted All updates are successfully applied.
	UpdateStatusCompleted UpdateStatus = "Completed"
	// UpdateStatusInProgress Updates installation are in progress.
	UpdateStatusInProgress UpdateStatus = "InProgress"
	// UpdateStatusPending There are pending updates to be installed.
	UpdateStatusPending UpdateStatus = "Pending"
	// UpdateStatusRetryLater Updates installation failed and should be retried later.
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	// UpdateStatusRetryNow Updates installation failed but are ready to retry again.
	UpdateStatusRetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{UpdateStatusCompleted, UpdateStatusInProgress, UpdateStatusPending, UpdateStatusRetryLater, UpdateStatusRetryNow}
}

// Visibility enumerates the values for visibility.
type Visibility string

const (
	// VisibilityCustom Only visible to users with permissions.
	VisibilityCustom Visibility = "Custom"
	// VisibilityPublic Visible to all users.
	VisibilityPublic Visibility = "Public"
)

// PossibleVisibilityValues returns an array of possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{VisibilityCustom, VisibilityPublic}
}
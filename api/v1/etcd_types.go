/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PVCRetentionPolicy defines policies for deleting/retaining Etcd PVCs
type PVCRetentionPolicy string

const (
	// PolicyDeleteAll is a constant for a policy type indicating that all the PVCs of etcd instances has to be deleted.
	PolicyDeleteAll PVCRetentionPolicy = "DeleteAll"
	// PolicyRetainMaster is a constant for a policy type indicating that all the PVCs except that of master has to be deleted.
	PolicyRetainMaster PVCRetentionPolicy = "RetainMaster"
	// PolicyRetainAll is a constant for a policy type indicating that all the PVCs of etcd has to be retained.
	PolicyRetainAll PVCRetentionPolicy = "RetRetainAllainMaster"
)

// EtcdConfig will hold information about the name and namespace of the etcd object.
type EtcdConfig struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// TLSConfig will hold information tls config for etcd server.
type TLSConfig struct {
	CACert                string `json:"cacert,omitempty"`
	Cert                  string `json:"cert,omitempty"`
	Key                   string `json:"key,omitempty"`
	InsecureSkipTLSVerify bool   `json:"insecureSkipTLSVerify,omitempty"`
	InsecureTransport     bool   `json:"insecureTransport,omitempty"`
}

// Spec defines the desired state of Etcd
type Spec struct {
	Etcd               EtcdSpec           `json:"etcd,omitempty"`
	Backup             BackupSpec         `json:"backup,omitempty"`
	Store              StoreSpec          `json:"store,omitempty"`
	PVCRetentionPolicy PVCRetentionPolicy `json:"pvcRetentionPolicy,omitempty"`
	Replicas           int                `json:"replicas,omitempty"`
}

// StoreSpec defines parameters related to ObjectStore persisting backups
type StoreSpec struct {
	Region           string                 `json:"region,omitempty"`
	StorageContainer string                 `json:"storageContainer,omitempty"`
	StorePrefix      string                 `json:"storePrefix,omitempty"`
	StorageProvider  string                 `json:"storageProvider,omitempty"`
	SecretRef        corev1.SecretReference `json:"secretRef,omitempty"`
}

// BackupSpec defines parametes associated with the full and delta snapshots of etcd
type BackupSpec struct {
	Labels                         map[string]string           `json:"labels,omitempty"`
	Annotations                    map[string]string           `json:"annotations,omitempty"`
	DeltaSnapshotMemoryLimit       int                         `json:"deltaSnapshotMemoryLimit,omitempty"`
	DeltaSnapshotPeriodSeconds     int                         `json:"deltaSnapshotPeriodSeconds,omitempty"`
	GarbageCollectionPeriodSeconds int                         `json:"garbageCollectionPeriodSeconds,omitempty"`
	Resources                      corev1.ResourceRequirements `json:"resources,omitempty"`
	FullSnapshotSchedule           string                      `json:"fullSnapshotSchedule,omitempty"`
	TempSnapDir                    string                      `json:"tempSnapDir,omitempty"`
}

// EtcdSpec defines parametes associated etcd deployed
type EtcdSpec struct {
	Labels                  map[string]string           `json:"labels,omitempty"`
	Annotations             map[string]string           `json:"annotations,omitempty"`
	Version                 string                      `json:"version,omitempty"`
	DataDir                 string                      `json:"dataDir,omitempty"`
	StorageCapacity         string                      `json:"storageCapacity,omitempty"`
	StorageClass            string                      `json:"storageClass,omitempty"`
	Config                  EtcdConfig                  `json:"etcdConfig,omitempty"`
	EtcdConnectionTimeout   int                         `json:"etcdConnectionTimeout,omitempty"`
	Resources               corev1.ResourceRequirements `json:"resources,omitempty"`
	DefragmentationSchedule string                      `json:"defragmentationSchedule,omitempty"`
	ServiceName             string                      `json:"serviceName,omitempty"`
	TLS                     TLSConfig                   `json:"tls,omitempty"`
}

// CrossVersionObjectReference contains enough information to let you identify the referred resource.
type CrossVersionObjectReference struct {
	// Kind of the referent
	Kind string `json:"kind"`
	// Name of the referent
	Name string `json:"name"`
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
}

// ConditionStatus is the status of a condition.
type ConditionStatus string

// ConditionType is a string alias.
type ConditionType string

const (
	// ConditionAvailable is a condition type for indicating availability.
	ConditionAvailable ConditionType = "Available"

	// ConditionTrue means a resource is in the condition.
	ConditionTrue ConditionStatus = "True"
	// ConditionFalse means a resource is not in the condition.
	ConditionFalse ConditionStatus = "False"
	// ConditionUnknown means Gardener can't decide if a resource is in the condition or not.
	ConditionUnknown ConditionStatus = "Unknown"
	// ConditionProgressing means the condition was seen true, failed but stayed within a predefined failure threshold.
	// In the future, we could add other intermediate conditions, e.g. ConditionDegraded.
	ConditionProgressing ConditionStatus = "Progressing"

	// ConditionCheckError is a constant for a reason in condition.
	ConditionCheckError = "ConditionCheckError"
)

// Condition holds the information about the state of a resource.
type Condition struct {
	// Type of the Shoot condition.
	Type ConditionType
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time
	// Last time the condition was updated.
	LastUpdateTime metav1.Time
	// The reason for the condition's last transition.
	Reason string
	// A human readable message indicating details about the transition.
	Message string
}

// EndpointStatus is the status of a condition.
type EndpointStatus string

// Endpoint holds information about etcd endpoints
type Endpoint struct {
	Name   string         `json:"name"`
	Status EndpointStatus `json:"status"`
	Port   int            `json:"port"`
}

// LastOperationType is a string alias.
type LastOperationType string

const (
	// LastOperationTypeCreate indicates a 'create' operation.
	LastOperationTypeCreate LastOperationType = "Create"
	// LastOperationTypeReconcile indicates a 'reconcile' operation.
	LastOperationTypeReconcile LastOperationType = "Reconcile"
	// LastOperationTypeDelete indicates a 'delete' operation.
	LastOperationTypeDelete LastOperationType = "Delete"
)

// LastOperationState is a string alias.
type LastOperationState string

const (
	// LastOperationStateProcessing indicates that an operation is ongoing.
	LastOperationStateProcessing LastOperationState = "Processing"
	// LastOperationStateSucceeded indicates that an operation has completed successfully.
	LastOperationStateSucceeded LastOperationState = "Succeeded"
	// LastOperationStateError indicates that an operation is completed with errors and will be retried.
	LastOperationStateError LastOperationState = "Error"
	// LastOperationStateFailed indicates that an operation is completed with errors and won't be retried.
	LastOperationStateFailed LastOperationState = "Failed"
	// LastOperationStatePending indicates that an operation cannot be done now, but will be tried in future.
	LastOperationStatePending LastOperationState = "Pending"
	// LastOperationStateAborted indicates that an operation has been aborted.
	LastOperationStateAborted LastOperationState = "Aborted"
)

// LastOperation indicates the type and the state of the last operation, along with a description
// message and a progress indicator.
type LastOperation struct {
	// A human readable message indicating details about the last operation.
	Description string `json:"description"`
	// Last time the operation state transitioned from one to another.
	LastUpdateTime metav1.Time `json:"lastUpdateTime"`
	// The progress in percentage (0-100) of the last operation.
	Progress int `json:"progress"`
	// Status of the last operation, one of Aborted, Processing, Succeeded, Error, Failed.
	State LastOperationState `json:"state"`
	// Type of the last operation, one of Create, Reconcile, Delete.
	Type LastOperationType `json:"type"`
}

// Status defines the observed state of Etcd
type Status struct {
	Etcd               CrossVersionObjectReference `json:"etcd"`
	Backup             CrossVersionObjectReference `json:"backup"`
	Conditions         []Condition                 `json:"conditions"`
	CurrentReplicas    int                         `json:"currentReplicas"`
	CurrentRevision    int                         `json:"currentRevision"`
	Endpoints          []Endpoint                  `json:"endpoints"`
	LastError          string                      `json:"lastError"`
	LastOperation      LastOperation               `json:"lastOperation,omitempty"`
	ObservedGeneration int                         `json:"observedGeneration"`
	Replicas           int                         `json:"replicas"`
	ReadyReplicas      int                         `json:"readyReplicas"`
	Ready              bool                        `json:"ready"`
	UpdatedReplicas    int                         `json:"updatedReplicas"`
	UpdatedRevision    int                         `json:"updatedRevision"`
}

// +kubebuilder:object:root=true

// Etcd is the Schema for the etcds API
type Etcd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Spec   `json:"spec,omitempty"`
	Status Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EtcdList contains a list of Etcd
type EtcdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Etcd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Etcd{}, &EtcdList{})
}

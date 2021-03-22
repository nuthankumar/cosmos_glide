package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBProvisioning is a top-level type
type DBProvisioning struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status DBProvisioningStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec DBProvisioningSpec `json:"spec,omitempty"`
}

// custom spec
type DBProvisioningSpec struct {
	Message      string `json:"message,omitempty"`
	Hello 	     string `json:"hello"`
	Clientid     string `json:"clientid"`
	Clientns     string `json:"clientns"`
	Keyspacename string `json:"keyspacename"`
	Rolename     string `json:"rolename"`
	Servicename  string `json:"servicename"`
	Username     string `json:"username"`
}

// custom status
type DBProvisioningStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type DBProvisioningList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []DBProvisioning `json:"items"`
}


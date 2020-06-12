package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	EMail             string `json:"e-mail,omitempty"`
	PrivKeySecretName string `json:"priv_key_secret,omitempty"`
	Domain            string `json:"Domain,omitempty"`
	TargetSecretName  string `json:"target_secret,omitempty"`
	Server            string `json:"server,omitempty"`
	IngressClass      string `json:"ingress_class,omitempty"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	IsValid   bool   `json:"is_valid,omitempty"`
	ExpiredAt string `json:"expired_at,omitempty"`
	Domain    string `json:"Domain,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Certificate is the Schema for the certificates API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=certificates,scope=Namespaced
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec,omitempty"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateList contains a list of Certificate
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:subresource:status

// KabaneroSpec defines the desired state of Kabanero
// +k8s:openapi-gen=true
type KabaneroSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	Version     string                   `json:"version,omitempty"`
	Github      GithubConfig             `json:"github,omitempty"`
	Collections InstanceCollectionConfig `json:"collections,omitempty"`

	Tekton TektonCustomizationSpec `json:"tekton,omitempty"`

	AppsodyOperator AppsodyCustomizationSpec `json:"appsodyOperator,omitempty"`

	CliServices KabaneroCliServicesCustomizationSpec `json:"cliServices,omitempty"`

	Landing KabaneroLandingCustomizationSpec `json:"landing,omitempty"`
}

type InstanceCollectionConfig struct {
	Repositories []RepositoryConfig `json:"repositories,omitempty"`
}

// This represents the Github information (public or GHE) where
// the organization and teams managing the collections live.  Members
// of the specified team in the specified organization will have admin
// authority in the Kabanero CLI.
type GithubConfig struct {
	Organization string `json:"organization,omitempty"`
	Team         string `json:"team,omitempty"`
	ApiUrl       string `json:"apiUrl,omitempty"`
}

type RepositoryConfig struct {
	Name                       string `json:"name,omitempty"`
	Url                        string `json:"url,omitempty"`
	ActivateDefaultCollections bool   `json:"activateDefaultCollections,omitempty"`
	SkipCertVerification       bool   `json:"skipCertVerification,omitempty"`
}

type TektonCustomizationSpec struct {
	Enable  string `json:"enable,omitempty"`
	Version string `json:"version,omitempty"`
}

type AppsodyCustomizationSpec struct {
	Enable     bool   `json:"enable,omitempty"`
	Version    string `json:"version,omitempty"`
	Image      string `json:"image,omitempty"`
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

type KabaneroCliServicesCustomizationSpec struct {
	Enable     bool   `json:"enable,omitempty"`
	Version    string `json:"version,omitempty"`
	Image      string `json:"image,omitempty"`
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

type KabaneroLandingCustomizationSpec struct {
	Enable bool `json:"enable,omitempty"`
}

// KabaneroStatus defines the observed state of the Kabanero instance
// +k8s:openapi-gen=true
type KabaneroStatus struct {
	// Kabanero operator instance readiness status. The status is directly correlated to the availability of resources dependencies.
	KabaneroInstance KabaneroInstanceStatus `json:"kabaneroInstance,omitempty"`

	// Knative eventing instance readiness status.
	KnativeEventing KnativeEventingStatus `json:"knativeEventing,omitempty"`

	// Knative serving instance readiness status.
	KnativeServing KnativeServingStatus `json:"knativeServing,omitempty"`

	// Tekton instance readiness status.
	Tekton TektonStatus `json:"tekton,omitempty"`

	// CLI readiness status.
	Cli CliStatus `json:"cli,omitempty"`

	// Kabanero Landing page readiness status.
	Landing KabaneroLandingPageStatus `json:"landing,omitempty"`

	// Appsody instance readiness status.
	Appsody AppsodyStatus `json:"appsody,omitempty"`
}

type KabaneroInstanceStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Version      string `json:"version,omitempty"`
}

type TektonStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Version      string `json:"version,omitempty"`
}

type KnativeEventingStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Version      string `json:"version,omitempty"`
}

type KnativeServingStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Version      string `json:"version,omitempty"`
}

type CliStatus struct {
	Ready        string   `json:"ready, omitempty"`
	ErrorMessage string   `json:"errorMessage,omitempty"`
	Hostnames    []string `json:"hostnames,omitempty"`
}

type KabaneroLandingPageStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Version      string `json:"version,omitempty"`
}

type AppsodyStatus struct {
	Ready        string `json:"ready,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kabanero is the Schema for the kabaneros API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Kabanero struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KabaneroSpec   `json:"spec,omitempty"`
	Status KabaneroStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KabaneroList contains a list of Kabanero
type KabaneroList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kabanero `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kabanero{}, &KabaneroList{})
}

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthorizationAttemptInfoInitParameters struct {
}

type AuthorizationAttemptInfoObservation struct {

	// (Output)
	// Human readable explanation about the issue. Provided to help address
	// the configuration issues.
	// Not guaranteed to be stable. For programmatic access use reason field.
	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	// (Output)
	// Domain name of the authorization attempt.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (Output)
	// Reason for failure of the authorization attempt for the domain.
	FailureReason *string `json:"failureReason,omitempty" tf:"failure_reason,omitempty"`

	// (Output)
	// A state of this Managed Certificate.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type AuthorizationAttemptInfoParameters struct {
}

type CertificateInitParameters struct {

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of label tags associated with the Certificate resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed *ManagedInitParameters `json:"managed,omitempty" tf:"managed,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The scope of the certificate.
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
	// See https://cloud.google.com/vpc/docs/edge-locations.
	// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
	// See https://cloud.google.com/compute/docs/regions-zones
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged *SelfManagedInitParameters `json:"selfManaged,omitempty" tf:"self_managed,omitempty"`
}

type CertificateObservation struct {

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/certificates/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of label tags associated with the Certificate resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	Managed *ManagedObservation `json:"managed,omitempty" tf:"managed,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The scope of the certificate.
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
	// See https://cloud.google.com/vpc/docs/edge-locations.
	// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
	// See https://cloud.google.com/compute/docs/regions-zones
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	SelfManaged *SelfManagedObservation `json:"selfManaged,omitempty" tf:"self_managed,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type CertificateParameters struct {

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of label tags associated with the Certificate resource.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Configuration and state of a Managed Certificate.
	// Certificate Manager provisions and renews Managed Certificates
	// automatically, for as long as it's authorized to do so.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Managed *ManagedParameters `json:"managed,omitempty" tf:"managed,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The scope of the certificate.
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
	// See https://cloud.google.com/vpc/docs/edge-locations.
	// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
	// See https://cloud.google.com/compute/docs/regions-zones
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Certificate data for a SelfManaged Certificate.
	// SelfManaged Certificates are uploaded by the user. Updating such
	// certificates before they expire remains the user's responsibility.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SelfManaged *SelfManagedParameters `json:"selfManaged,omitempty" tf:"self_managed,omitempty"`
}

type ManagedInitParameters struct {

	// Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.DNSAuthorization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DNSAuthorizations []*string `json:"dnsAuthorizations,omitempty" tf:"dns_authorizations,omitempty"`

	// References to DNSAuthorization in certificatemanager to populate dnsAuthorizations.
	// +kubebuilder:validation:Optional
	DNSAuthorizationsRefs []v1.Reference `json:"dnsAuthorizationsRefs,omitempty" tf:"-"`

	// Selector for a list of DNSAuthorization in certificatemanager to populate dnsAuthorizations.
	// +kubebuilder:validation:Optional
	DNSAuthorizationsSelector *v1.Selector `json:"dnsAuthorizationsSelector,omitempty" tf:"-"`

	// The domains for which a managed SSL certificate will be generated.
	// Wildcard domains are only supported with DNS challenge resolution
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.DNSAuthorization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain",false)
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// References to DNSAuthorization in certificatemanager to populate domains.
	// +kubebuilder:validation:Optional
	DomainsRefs []v1.Reference `json:"domainsRefs,omitempty" tf:"-"`

	// Selector for a list of DNSAuthorization in certificatemanager to populate domains.
	// +kubebuilder:validation:Optional
	DomainsSelector *v1.Selector `json:"domainsSelector,omitempty" tf:"-"`

	// The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects//locations//certificateIssuanceConfigs/*.
	// If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	// Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	IssuanceConfig *string `json:"issuanceConfig,omitempty" tf:"issuance_config,omitempty"`
}

type ManagedObservation struct {

	// (Output)
	// Detailed state of the latest authorization attempt for each domain
	// specified for this Managed Certificate.
	// Structure is documented below.
	AuthorizationAttemptInfo []AuthorizationAttemptInfoObservation `json:"authorizationAttemptInfo,omitempty" tf:"authorization_attempt_info,omitempty"`

	// Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	DNSAuthorizations []*string `json:"dnsAuthorizations,omitempty" tf:"dns_authorizations,omitempty"`

	// The domains for which a managed SSL certificate will be generated.
	// Wildcard domains are only supported with DNS challenge resolution
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects//locations//certificateIssuanceConfigs/*.
	// If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	// Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	IssuanceConfig *string `json:"issuanceConfig,omitempty" tf:"issuance_config,omitempty"`

	// (Output)
	// Information about issues with provisioning this Managed Certificate.
	// Structure is documented below.
	ProvisioningIssue []ProvisioningIssueObservation `json:"provisioningIssue,omitempty" tf:"provisioning_issue,omitempty"`

	// (Output)
	// A state of this Managed Certificate.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ManagedParameters struct {

	// Authorizations that will be used for performing domain authorization. Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.DNSAuthorization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DNSAuthorizations []*string `json:"dnsAuthorizations,omitempty" tf:"dns_authorizations,omitempty"`

	// References to DNSAuthorization in certificatemanager to populate dnsAuthorizations.
	// +kubebuilder:validation:Optional
	DNSAuthorizationsRefs []v1.Reference `json:"dnsAuthorizationsRefs,omitempty" tf:"-"`

	// Selector for a list of DNSAuthorization in certificatemanager to populate dnsAuthorizations.
	// +kubebuilder:validation:Optional
	DNSAuthorizationsSelector *v1.Selector `json:"dnsAuthorizationsSelector,omitempty" tf:"-"`

	// The domains for which a managed SSL certificate will be generated.
	// Wildcard domains are only supported with DNS challenge resolution
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.DNSAuthorization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain",false)
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// References to DNSAuthorization in certificatemanager to populate domains.
	// +kubebuilder:validation:Optional
	DomainsRefs []v1.Reference `json:"domainsRefs,omitempty" tf:"-"`

	// Selector for a list of DNSAuthorization in certificatemanager to populate domains.
	// +kubebuilder:validation:Optional
	DomainsSelector *v1.Selector `json:"domainsSelector,omitempty" tf:"-"`

	// The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects//locations//certificateIssuanceConfigs/*.
	// If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	// Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
	// +kubebuilder:validation:Optional
	IssuanceConfig *string `json:"issuanceConfig,omitempty" tf:"issuance_config,omitempty"`
}

type ProvisioningIssueInitParameters struct {
}

type ProvisioningIssueObservation struct {

	// (Output)
	// Human readable explanation about the issue. Provided to help address
	// the configuration issues.
	// Not guaranteed to be stable. For programmatic access use reason field.
	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	// (Output)
	// Reason for provisioning failures.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type ProvisioningIssueParameters struct {
}

type SelfManagedInitParameters struct {

	// The certificate chain in PEM-encoded form.
	// Leaf certificate comes first, followed by intermediate ones if any.
	// Note: This property is sensitive and will not be displayed in the plan.
	CertificatePemSecretRef *v1.SecretKeySelector `json:"certificatePemSecretRef,omitempty" tf:"-"`

	// The certificate chain in PEM-encoded form.
	// Leaf certificate comes first, followed by intermediate ones if any.
	PemCertificate *string `json:"pemCertificate,omitempty" tf:"pem_certificate,omitempty"`

	// The private key of the leaf certificate in PEM-encoded form.
	// Note: This property is sensitive and will not be displayed in the plan.
	PemPrivateKeySecretRef *v1.SecretKeySelector `json:"pemPrivateKeySecretRef,omitempty" tf:"-"`

	// The private key of the leaf certificate in PEM-encoded form.
	// Note: This property is sensitive and will not be displayed in the plan.
	PrivateKeyPemSecretRef *v1.SecretKeySelector `json:"privateKeyPemSecretRef,omitempty" tf:"-"`
}

type SelfManagedObservation struct {

	// The certificate chain in PEM-encoded form.
	// Leaf certificate comes first, followed by intermediate ones if any.
	PemCertificate *string `json:"pemCertificate,omitempty" tf:"pem_certificate,omitempty"`
}

type SelfManagedParameters struct {

	// The certificate chain in PEM-encoded form.
	// Leaf certificate comes first, followed by intermediate ones if any.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	CertificatePemSecretRef *v1.SecretKeySelector `json:"certificatePemSecretRef,omitempty" tf:"-"`

	// The certificate chain in PEM-encoded form.
	// Leaf certificate comes first, followed by intermediate ones if any.
	// +kubebuilder:validation:Optional
	PemCertificate *string `json:"pemCertificate,omitempty" tf:"pem_certificate,omitempty"`

	// The private key of the leaf certificate in PEM-encoded form.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PemPrivateKeySecretRef *v1.SecretKeySelector `json:"pemPrivateKeySecretRef,omitempty" tf:"-"`

	// The private key of the leaf certificate in PEM-encoded form.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PrivateKeyPemSecretRef *v1.SecretKeySelector `json:"privateKeyPemSecretRef,omitempty" tf:"-"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Certificate is the Schema for the Certificates API. Certificate represents a HTTP-reachable backend for a Certificate.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}

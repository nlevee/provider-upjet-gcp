/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoHealingPoliciesObservation struct {
}

type AutoHealingPoliciesParameters struct {

	// The health check resource that signals autohealing.
	// +crossplane:generate:reference:type=HealthCheck
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	HealthCheck *string `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Reference to a HealthCheck to populate healthCheck.
	// +kubebuilder:validation:Optional
	HealthCheckRef *v1.Reference `json:"healthCheckRef,omitempty" tf:"-"`

	// Selector for a HealthCheck to populate healthCheck.
	// +kubebuilder:validation:Optional
	HealthCheckSelector *v1.Selector `json:"healthCheckSelector,omitempty" tf:"-"`

	// The number of seconds that the managed instance group waits before
	// it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.
	// +kubebuilder:validation:Required
	InitialDelaySec *float64 `json:"initialDelaySec" tf:"initial_delay_sec,omitempty"`
}

type InstanceGroupManagerNamedPortObservation struct {
}

type InstanceGroupManagerNamedPortParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The port number.
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`
}

type InstanceGroupManagerObservation struct {

	// The fingerprint of the instance group manager.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The full URL of the instance group created by the manager.
	InstanceGroup *string `json:"instanceGroup,omitempty" tf:"instance_group,omitempty"`

	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The URL of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The status of this managed instance group.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceGroupManagerParameters struct {

	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the official documentation.
	// +kubebuilder:validation:Optional
	AutoHealingPolicies []AutoHealingPoliciesParameters `json:"autoHealingPolicies,omitempty" tf:"auto_healing_policies,omitempty"`

	// The base instance name to use for
	// instances in this group. The value must be a valid
	// RFC1035 name. Supported characters
	// are lowercase letters, numbers, and hyphens . Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	// +kubebuilder:validation:Required
	BaseInstanceName *string `json:"baseInstanceName" tf:"base_instance_name,omitempty"`

	// An optional textual description of the instance
	// group manager.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The named port configuration. See the section below
	// for details on configuration.
	// +kubebuilder:validation:Optional
	NamedPort []InstanceGroupManagerNamedPortParameters `json:"namedPort,omitempty" tf:"named_port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// ) Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the official documentation.
	// +kubebuilder:validation:Optional
	StatefulDisk []StatefulDiskParameters `json:"statefulDisk,omitempty" tf:"stateful_disk,omitempty"`

	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	// +crossplane:generate:reference:type=TargetPool
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	TargetPools []*string `json:"targetPools,omitempty" tf:"target_pools,omitempty"`

	// References to TargetPool to populate targetPools.
	// +kubebuilder:validation:Optional
	TargetPoolsRefs []v1.Reference `json:"targetPoolsRefs,omitempty" tf:"-"`

	// Selector for a list of TargetPool to populate targetPools.
	// +kubebuilder:validation:Optional
	TargetPoolsSelector *v1.Selector `json:"targetPoolsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetSize *float64 `json:"targetSize,omitempty" tf:"target_size,omitempty"`

	// The update policy for this managed instance group. Structure is documented below. For more information, see the official documentation and API
	// +kubebuilder:validation:Optional
	UpdatePolicy []UpdatePolicyParameters `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Version []VersionParameters `json:"version" tf:"version,omitempty"`

	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, Terraform will
	// continue trying until it times out.
	// +kubebuilder:validation:Optional
	WaitForInstances *bool `json:"waitForInstances,omitempty" tf:"wait_for_instances,omitempty"`

	// When used with wait_for_instances it specifies the status to wait for.
	// When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are STABLE and UPDATED
	// +kubebuilder:validation:Optional
	WaitForInstancesStatus *string `json:"waitForInstancesStatus,omitempty" tf:"wait_for_instances_status,omitempty"`

	// The zone that instances in this group should be created
	// in.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type PerInstanceConfigsObservation struct {

	// A bit indicating if all of the group's per-instance configs  have status EFFECTIVE or there are no per-instance-configs.
	AllEffective *bool `json:"allEffective,omitempty" tf:"all_effective,omitempty"`
}

type PerInstanceConfigsParameters struct {
}

type StatefulDiskObservation struct {
}

type StatefulDiskParameters struct {

	// , A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION. NEVER - detach the disk when the VM is deleted, but do not delete the disk. ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently deleted from the instance group. The default is NEVER.
	// +kubebuilder:validation:Optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// , The device name of the disk to be attached.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`
}

type StatefulObservation struct {

	// A bit indicating whether the managed instance group has stateful configuration, that is, if you have configured any items in a stateful policy or in per-instance configs. The group might report that it has no stateful config even when there is still some preserved state on a managed instance, for example, if you have deleted all PICs but not yet applied those deletions.
	HasStatefulConfig *bool `json:"hasStatefulConfig,omitempty" tf:"has_stateful_config,omitempty"`

	// Status of per-instance configs on the instance.
	PerInstanceConfigs []PerInstanceConfigsObservation `json:"perInstanceConfigs,omitempty" tf:"per_instance_configs,omitempty"`
}

type StatefulParameters struct {
}

type StatusObservation struct {

	// A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change ; no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
	IsStable *bool `json:"isStable,omitempty" tf:"is_stable,omitempty"`

	// Stateful status of the given Instance Group Manager.
	Stateful []StatefulObservation `json:"stateful,omitempty" tf:"stateful,omitempty"`

	VersionTarget []VersionTargetObservation `json:"versionTarget,omitempty" tf:"version_target,omitempty"`
}

type StatusParameters struct {
}

type TargetSizeObservation struct {
}

type TargetSizeParameters struct {

	// , The number of instances which are managed for this version. Conflicts with percent.
	// +kubebuilder:validation:Optional
	Fixed *float64 `json:"fixed,omitempty" tf:"fixed,omitempty"`

	// , The number of instances  which are managed for this version. Conflicts with fixed.
	// Note that when using percent, rounding will be in favor of explicitly set target_size values; a managed instance group with 2 instances and 2 versions,
	// one of which has a target_size.percent of 60 will create 2 instances of that version.
	// +kubebuilder:validation:Optional
	Percent *float64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

type UpdatePolicyObservation struct {
}

type UpdatePolicyParameters struct {

	// , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with max_surge_percent. If neither is set, defaults to 1
	// +kubebuilder:validation:Optional
	MaxSurgeFixed *float64 `json:"maxSurgeFixed,omitempty" tf:"max_surge_fixed,omitempty"`

	// , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with max_surge_fixed.
	// +kubebuilder:validation:Optional
	MaxSurgePercent *float64 `json:"maxSurgePercent,omitempty" tf:"max_surge_percent,omitempty"`

	// , The maximum number of instances that can be unavailable during the update process. Conflicts with max_unavailable_percent. If neither is set, defaults to 1
	// +kubebuilder:validation:Optional
	MaxUnavailableFixed *float64 `json:"maxUnavailableFixed,omitempty" tf:"max_unavailable_fixed,omitempty"`

	// , The maximum number of instances that can be unavailable during the update process. Conflicts with max_unavailable_fixed.
	// +kubebuilder:validation:Optional
	MaxUnavailablePercent *float64 `json:"maxUnavailablePercent,omitempty" tf:"max_unavailable_percent,omitempty"`

	// - Minimal action to be taken on an instance. You can specify either REFRESH to update without stopping instances, RESTART to restart existing instances or REPLACE to delete and create new instances from the target template. If you specify a REFRESH, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
	// +kubebuilder:validation:Required
	MinimalAction *string `json:"minimalAction" tf:"minimal_action,omitempty"`

	// - Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
	// +kubebuilder:validation:Optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// , The instance replacement method for managed instance groups. Valid values are: "RECREATE", "SUBSTITUTE". If SUBSTITUTE , the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
	// +kubebuilder:validation:Optional
	ReplacementMethod *string `json:"replacementMethod,omitempty" tf:"replacement_method,omitempty"`

	// - The type of update process. You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions .
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VersionObservation struct {
}

type VersionParameters struct {

	// - The full URL to an instance template from which all new instances of this version will be created.
	// +crossplane:generate:reference:type=InstanceTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceTemplate *string `json:"instanceTemplate,omitempty" tf:"instance_template,omitempty"`

	// Reference to a InstanceTemplate to populate instanceTemplate.
	// +kubebuilder:validation:Optional
	InstanceTemplateRef *v1.Reference `json:"instanceTemplateRef,omitempty" tf:"-"`

	// Selector for a InstanceTemplate to populate instanceTemplate.
	// +kubebuilder:validation:Optional
	InstanceTemplateSelector *v1.Selector `json:"instanceTemplateSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TargetSize []TargetSizeParameters `json:"targetSize,omitempty" tf:"target_size,omitempty"`
}

type VersionTargetObservation struct {
	IsReached *bool `json:"isReached,omitempty" tf:"is_reached,omitempty"`
}

type VersionTargetParameters struct {
}

// InstanceGroupManagerSpec defines the desired state of InstanceGroupManager
type InstanceGroupManagerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceGroupManagerParameters `json:"forProvider"`
}

// InstanceGroupManagerStatus defines the observed state of InstanceGroupManager.
type InstanceGroupManagerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceGroupManagerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupManager is the Schema for the InstanceGroupManagers API. Manages an Instance Group within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupManagerSpec   `json:"spec"`
	Status            InstanceGroupManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupManagerList contains a list of InstanceGroupManagers
type InstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceGroupManager `json:"items"`
}

// Repository type metadata.
var (
	InstanceGroupManager_Kind             = "InstanceGroupManager"
	InstanceGroupManager_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceGroupManager_Kind}.String()
	InstanceGroupManager_KindAPIVersion   = InstanceGroupManager_Kind + "." + CRDGroupVersion.String()
	InstanceGroupManager_GroupVersionKind = CRDGroupVersion.WithKind(InstanceGroupManager_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceGroupManager{}, &InstanceGroupManagerList{})
}

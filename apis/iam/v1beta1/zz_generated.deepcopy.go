//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsObservation) DeepCopyInto(out *AwsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsObservation.
func (in *AwsObservation) DeepCopy() *AwsObservation {
	if in == nil {
		return nil
	}
	out := new(AwsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsParameters) DeepCopyInto(out *AwsParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsParameters.
func (in *AwsParameters) DeepCopy() *AwsParameters {
	if in == nil {
		return nil
	}
	out := new(AwsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcObservation) DeepCopyInto(out *OidcObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcObservation.
func (in *OidcObservation) DeepCopy() *OidcObservation {
	if in == nil {
		return nil
	}
	out := new(OidcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcParameters) DeepCopyInto(out *OidcParameters) {
	*out = *in
	if in.AllowedAudiences != nil {
		in, out := &in.AllowedAudiences, &out.AllowedAudiences
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IssuerURI != nil {
		in, out := &in.IssuerURI, &out.IssuerURI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcParameters.
func (in *OidcParameters) DeepCopy() *OidcParameters {
	if in == nil {
		return nil
	}
	out := new(OidcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPool) DeepCopyInto(out *WorkloadIdentityPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPool.
func (in *WorkloadIdentityPool) DeepCopy() *WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolList) DeepCopyInto(out *WorkloadIdentityPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadIdentityPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolList.
func (in *WorkloadIdentityPoolList) DeepCopy() *WorkloadIdentityPoolList {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolObservation) DeepCopyInto(out *WorkloadIdentityPoolObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolObservation.
func (in *WorkloadIdentityPoolObservation) DeepCopy() *WorkloadIdentityPoolObservation {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolParameters) DeepCopyInto(out *WorkloadIdentityPoolParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolParameters.
func (in *WorkloadIdentityPoolParameters) DeepCopy() *WorkloadIdentityPoolParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProvider) DeepCopyInto(out *WorkloadIdentityPoolProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProvider.
func (in *WorkloadIdentityPoolProvider) DeepCopy() *WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityPoolProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProviderList) DeepCopyInto(out *WorkloadIdentityPoolProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadIdentityPoolProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProviderList.
func (in *WorkloadIdentityPoolProviderList) DeepCopy() *WorkloadIdentityPoolProviderList {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadIdentityPoolProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProviderObservation) DeepCopyInto(out *WorkloadIdentityPoolProviderObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProviderObservation.
func (in *WorkloadIdentityPoolProviderObservation) DeepCopy() *WorkloadIdentityPoolProviderObservation {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProviderParameters) DeepCopyInto(out *WorkloadIdentityPoolProviderParameters) {
	*out = *in
	if in.AttributeCondition != nil {
		in, out := &in.AttributeCondition, &out.AttributeCondition
		*out = new(string)
		**out = **in
	}
	if in.AttributeMapping != nil {
		in, out := &in.AttributeMapping, &out.AttributeMapping
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Aws != nil {
		in, out := &in.Aws, &out.Aws
		*out = make([]AwsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = make([]OidcParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.WorkloadIdentityPoolID != nil {
		in, out := &in.WorkloadIdentityPoolID, &out.WorkloadIdentityPoolID
		*out = new(string)
		**out = **in
	}
	if in.WorkloadIdentityPoolIDRef != nil {
		in, out := &in.WorkloadIdentityPoolIDRef, &out.WorkloadIdentityPoolIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadIdentityPoolIDSelector != nil {
		in, out := &in.WorkloadIdentityPoolIDSelector, &out.WorkloadIdentityPoolIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProviderParameters.
func (in *WorkloadIdentityPoolProviderParameters) DeepCopy() *WorkloadIdentityPoolProviderParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProviderSpec) DeepCopyInto(out *WorkloadIdentityPoolProviderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProviderSpec.
func (in *WorkloadIdentityPoolProviderSpec) DeepCopy() *WorkloadIdentityPoolProviderSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolProviderStatus) DeepCopyInto(out *WorkloadIdentityPoolProviderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolProviderStatus.
func (in *WorkloadIdentityPoolProviderStatus) DeepCopy() *WorkloadIdentityPoolProviderStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolSpec) DeepCopyInto(out *WorkloadIdentityPoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolSpec.
func (in *WorkloadIdentityPoolSpec) DeepCopy() *WorkloadIdentityPoolSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityPoolStatus) DeepCopyInto(out *WorkloadIdentityPoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityPoolStatus.
func (in *WorkloadIdentityPoolStatus) DeepCopy() *WorkloadIdentityPoolStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityPoolStatus)
	in.DeepCopyInto(out)
	return out
}

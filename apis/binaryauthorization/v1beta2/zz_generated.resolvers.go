// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Attestor.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Attestor) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.AttestationAuthorityNote != nil {
		{
			m, l, err = apisresolver.GetManagedResource("containeranalysis.gcp.upbound.io", "v1beta2", "Note", "NoteList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AttestationAuthorityNote.NoteReference),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.AttestationAuthorityNote.NoteReferenceRef,
				Selector:     mg.Spec.ForProvider.AttestationAuthorityNote.NoteReferenceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AttestationAuthorityNote.NoteReference")
		}
		mg.Spec.ForProvider.AttestationAuthorityNote.NoteReference = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AttestationAuthorityNote.NoteReferenceRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.AttestationAuthorityNote != nil {
		{
			m, l, err = apisresolver.GetManagedResource("containeranalysis.gcp.upbound.io", "v1beta2", "Note", "NoteList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AttestationAuthorityNote.NoteReference),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.AttestationAuthorityNote.NoteReferenceRef,
				Selector:     mg.Spec.InitProvider.AttestationAuthorityNote.NoteReferenceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AttestationAuthorityNote.NoteReference")
		}
		mg.Spec.InitProvider.AttestationAuthorityNote.NoteReference = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AttestationAuthorityNote.NoteReferenceRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ClusterAdmissionRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("binaryauthorization.gcp.upbound.io", "v1beta2", "Attestor", "AttestorList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsBy),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsByRefs,
				Selector:      mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsBySelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsBy")
		}
		mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsBy = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.ClusterAdmissionRules[i3].RequireAttestationsByRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.ForProvider.DefaultAdmissionRule != nil {
		{
			m, l, err = apisresolver.GetManagedResource("binaryauthorization.gcp.upbound.io", "v1beta2", "Attestor", "AttestorList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsBy),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsByRefs,
				Selector:      mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsBySelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsBy")
		}
		mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsBy = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.DefaultAdmissionRule.RequireAttestationsByRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ClusterAdmissionRules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("binaryauthorization.gcp.upbound.io", "v1beta2", "Attestor", "AttestorList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsBy),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsByRefs,
				Selector:      mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsBySelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsBy")
		}
		mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsBy = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.ClusterAdmissionRules[i3].RequireAttestationsByRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.DefaultAdmissionRule != nil {
		{
			m, l, err = apisresolver.GetManagedResource("binaryauthorization.gcp.upbound.io", "v1beta2", "Attestor", "AttestorList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsBy),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsByRefs,
				Selector:      mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsBySelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsBy")
		}
		mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsBy = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.DefaultAdmissionRule.RequireAttestationsByRefs = mrsp.ResolvedReferences

	}

	return nil
}

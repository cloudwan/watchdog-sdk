// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/activation_custom.proto
// DO NOT EDIT!!!

package activation_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account_key"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &iam_service_account_key.ServiceAccountKey{}
	_ = &probe.Probe{}
)

func (o *ActivationRequest) GotenObjectExt() {}

func (o *ActivationRequest) MakeFullFieldMask() *ActivationRequest_FieldMask {
	return FullActivationRequest_FieldMask()
}

func (o *ActivationRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationRequest_FieldMask()
}

func (o *ActivationRequest) MakeDiffFieldMask(other *ActivationRequest) *ActivationRequest_FieldMask {
	if o == nil && other == nil {
		return &ActivationRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationRequest_FieldMask()
	}

	res := &ActivationRequest_FieldMask{}
	{
		subMask := o.GetActivate().MakeDiffFieldMask(other.GetActivate())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationRequest_FieldTerminalPath{selector: ActivationRequest_FieldPathSelectorActivate})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationRequest_FieldSubPath{selector: ActivationRequest_FieldPathSelectorActivate, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetConfirmActivation().MakeDiffFieldMask(other.GetConfirmActivation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationRequest_FieldTerminalPath{selector: ActivationRequest_FieldPathSelectorConfirmActivation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationRequest_FieldSubPath{selector: ActivationRequest_FieldPathSelectorConfirmActivation, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ActivationRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationRequest))
}

func (o *ActivationRequest) Clone() *ActivationRequest {
	if o == nil {
		return nil
	}
	result := &ActivationRequest{}
	if o, ok := o.Msg.(*ActivationRequest_Activate_); ok {
		result.Msg = (*ActivationRequest_Activate_)(nil)
		if o != nil {
			result.Msg = &ActivationRequest_Activate_{}
			result := result.Msg.(*ActivationRequest_Activate_)
			result.Activate = o.Activate.Clone()
		}
	}
	if o, ok := o.Msg.(*ActivationRequest_ConfirmActivation_); ok {
		result.Msg = (*ActivationRequest_ConfirmActivation_)(nil)
		if o != nil {
			result.Msg = &ActivationRequest_ConfirmActivation_{}
			result := result.Msg.(*ActivationRequest_ConfirmActivation_)
			result.ConfirmActivation = o.ConfirmActivation.Clone()
		}
	}
	return result
}

func (o *ActivationRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationRequest) Merge(source *ActivationRequest) {
	if source, ok := source.GetMsg().(*ActivationRequest_Activate_); ok {
		if dstOneOf, ok := o.Msg.(*ActivationRequest_Activate_); !ok || dstOneOf == nil {
			o.Msg = &ActivationRequest_Activate_{}
		}
		if source != nil {
			o := o.Msg.(*ActivationRequest_Activate_)
			if source.Activate != nil {
				if o.Activate == nil {
					o.Activate = new(ActivationRequest_Activate)
				}
				o.Activate.Merge(source.Activate)
			}
		}
	}
	if source, ok := source.GetMsg().(*ActivationRequest_ConfirmActivation_); ok {
		if dstOneOf, ok := o.Msg.(*ActivationRequest_ConfirmActivation_); !ok || dstOneOf == nil {
			o.Msg = &ActivationRequest_ConfirmActivation_{}
		}
		if source != nil {
			o := o.Msg.(*ActivationRequest_ConfirmActivation_)
			if source.ConfirmActivation != nil {
				if o.ConfirmActivation == nil {
					o.ConfirmActivation = new(ActivationRequest_ConfirmActivation)
				}
				o.ConfirmActivation.Merge(source.ConfirmActivation)
			}
		}
	}
}

func (o *ActivationRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationRequest))
}

func (o *ActivationRequest_Activate) GotenObjectExt() {}

func (o *ActivationRequest_Activate) MakeFullFieldMask() *ActivationRequest_Activate_FieldMask {
	return FullActivationRequest_Activate_FieldMask()
}

func (o *ActivationRequest_Activate) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationRequest_Activate_FieldMask()
}

func (o *ActivationRequest_Activate) MakeDiffFieldMask(other *ActivationRequest_Activate) *ActivationRequest_Activate_FieldMask {
	if o == nil && other == nil {
		return &ActivationRequest_Activate_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationRequest_Activate_FieldMask()
	}

	res := &ActivationRequest_Activate_FieldMask{}
	if o.GetToken() != other.GetToken() {
		res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorToken})
	}
	if o.GetClientCookie() != other.GetClientCookie() {
		res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorClientCookie})
	}
	{
		subMask := o.GetProbe().MakeDiffFieldMask(other.GetProbe())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorProbe})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationRequestActivate_FieldSubPath{selector: ActivationRequestActivate_FieldPathSelectorProbe, subPath: subpath})
			}
		}
	}

	if len(o.GetMetadata()) == len(other.GetMetadata()) {
		for i, lValue := range o.GetMetadata() {
			rValue := other.GetMetadata()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorMetadata})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorMetadata})
	}
	if o.GetPublicKeyData() != other.GetPublicKeyData() {
		res.Paths = append(res.Paths, &ActivationRequestActivate_FieldTerminalPath{selector: ActivationRequestActivate_FieldPathSelectorPublicKeyData})
	}
	return res
}

func (o *ActivationRequest_Activate) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationRequest_Activate))
}

func (o *ActivationRequest_Activate) Clone() *ActivationRequest_Activate {
	if o == nil {
		return nil
	}
	result := &ActivationRequest_Activate{}
	result.Token = o.Token
	result.ClientCookie = o.ClientCookie
	result.Probe = o.Probe.Clone()
	result.Metadata = map[string]string{}
	for key, sourceValue := range o.Metadata {
		result.Metadata[key] = sourceValue
	}
	result.PublicKeyData = o.PublicKeyData
	return result
}

func (o *ActivationRequest_Activate) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationRequest_Activate) Merge(source *ActivationRequest_Activate) {
	o.Token = source.GetToken()
	o.ClientCookie = source.GetClientCookie()
	if source.GetProbe() != nil {
		if o.Probe == nil {
			o.Probe = new(probe.Probe)
		}
		o.Probe.Merge(source.GetProbe())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = make(map[string]string, len(source.GetMetadata()))
		}
		for key, sourceValue := range source.GetMetadata() {
			o.Metadata[key] = sourceValue
		}
	}
	o.PublicKeyData = source.GetPublicKeyData()
}

func (o *ActivationRequest_Activate) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationRequest_Activate))
}

func (o *ActivationRequest_ConfirmActivation) GotenObjectExt() {}

func (o *ActivationRequest_ConfirmActivation) MakeFullFieldMask() *ActivationRequest_ConfirmActivation_FieldMask {
	return FullActivationRequest_ConfirmActivation_FieldMask()
}

func (o *ActivationRequest_ConfirmActivation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationRequest_ConfirmActivation_FieldMask()
}

func (o *ActivationRequest_ConfirmActivation) MakeDiffFieldMask(other *ActivationRequest_ConfirmActivation) *ActivationRequest_ConfirmActivation_FieldMask {
	if o == nil && other == nil {
		return &ActivationRequest_ConfirmActivation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationRequest_ConfirmActivation_FieldMask()
	}

	res := &ActivationRequest_ConfirmActivation_FieldMask{}
	if o.GetClientCookie() != other.GetClientCookie() {
		res.Paths = append(res.Paths, &ActivationRequestConfirmActivation_FieldTerminalPath{selector: ActivationRequestConfirmActivation_FieldPathSelectorClientCookie})
	}
	return res
}

func (o *ActivationRequest_ConfirmActivation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationRequest_ConfirmActivation))
}

func (o *ActivationRequest_ConfirmActivation) Clone() *ActivationRequest_ConfirmActivation {
	if o == nil {
		return nil
	}
	result := &ActivationRequest_ConfirmActivation{}
	result.ClientCookie = o.ClientCookie
	return result
}

func (o *ActivationRequest_ConfirmActivation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationRequest_ConfirmActivation) Merge(source *ActivationRequest_ConfirmActivation) {
	o.ClientCookie = source.GetClientCookie()
}

func (o *ActivationRequest_ConfirmActivation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationRequest_ConfirmActivation))
}

func (o *ActivationResponse) GotenObjectExt() {}

func (o *ActivationResponse) MakeFullFieldMask() *ActivationResponse_FieldMask {
	return FullActivationResponse_FieldMask()
}

func (o *ActivationResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationResponse_FieldMask()
}

func (o *ActivationResponse) MakeDiffFieldMask(other *ActivationResponse) *ActivationResponse_FieldMask {
	if o == nil && other == nil {
		return &ActivationResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationResponse_FieldMask()
	}

	res := &ActivationResponse_FieldMask{}
	{
		subMask := o.GetActivated().MakeDiffFieldMask(other.GetActivated())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationResponse_FieldTerminalPath{selector: ActivationResponse_FieldPathSelectorActivated})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationResponse_FieldSubPath{selector: ActivationResponse_FieldPathSelectorActivated, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetConfirmActivationAck().MakeDiffFieldMask(other.GetConfirmActivationAck())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationResponse_FieldTerminalPath{selector: ActivationResponse_FieldPathSelectorConfirmActivationAck})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationResponse_FieldSubPath{selector: ActivationResponse_FieldPathSelectorConfirmActivationAck, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ActivationResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationResponse))
}

func (o *ActivationResponse) Clone() *ActivationResponse {
	if o == nil {
		return nil
	}
	result := &ActivationResponse{}
	if o, ok := o.Msg.(*ActivationResponse_Activated_); ok {
		result.Msg = (*ActivationResponse_Activated_)(nil)
		if o != nil {
			result.Msg = &ActivationResponse_Activated_{}
			result := result.Msg.(*ActivationResponse_Activated_)
			result.Activated = o.Activated.Clone()
		}
	}
	if o, ok := o.Msg.(*ActivationResponse_ConfirmActivationAck_); ok {
		result.Msg = (*ActivationResponse_ConfirmActivationAck_)(nil)
		if o != nil {
			result.Msg = &ActivationResponse_ConfirmActivationAck_{}
			result := result.Msg.(*ActivationResponse_ConfirmActivationAck_)
			result.ConfirmActivationAck = o.ConfirmActivationAck.Clone()
		}
	}
	return result
}

func (o *ActivationResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationResponse) Merge(source *ActivationResponse) {
	if source, ok := source.GetMsg().(*ActivationResponse_Activated_); ok {
		if dstOneOf, ok := o.Msg.(*ActivationResponse_Activated_); !ok || dstOneOf == nil {
			o.Msg = &ActivationResponse_Activated_{}
		}
		if source != nil {
			o := o.Msg.(*ActivationResponse_Activated_)
			if source.Activated != nil {
				if o.Activated == nil {
					o.Activated = new(ActivationResponse_Activated)
				}
				o.Activated.Merge(source.Activated)
			}
		}
	}
	if source, ok := source.GetMsg().(*ActivationResponse_ConfirmActivationAck_); ok {
		if dstOneOf, ok := o.Msg.(*ActivationResponse_ConfirmActivationAck_); !ok || dstOneOf == nil {
			o.Msg = &ActivationResponse_ConfirmActivationAck_{}
		}
		if source != nil {
			o := o.Msg.(*ActivationResponse_ConfirmActivationAck_)
			if source.ConfirmActivationAck != nil {
				if o.ConfirmActivationAck == nil {
					o.ConfirmActivationAck = new(ActivationResponse_ConfirmActivationAck)
				}
				o.ConfirmActivationAck.Merge(source.ConfirmActivationAck)
			}
		}
	}
}

func (o *ActivationResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationResponse))
}

func (o *ActivationResponse_Activated) GotenObjectExt() {}

func (o *ActivationResponse_Activated) MakeFullFieldMask() *ActivationResponse_Activated_FieldMask {
	return FullActivationResponse_Activated_FieldMask()
}

func (o *ActivationResponse_Activated) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationResponse_Activated_FieldMask()
}

func (o *ActivationResponse_Activated) MakeDiffFieldMask(other *ActivationResponse_Activated) *ActivationResponse_Activated_FieldMask {
	if o == nil && other == nil {
		return &ActivationResponse_Activated_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationResponse_Activated_FieldMask()
	}

	res := &ActivationResponse_Activated_FieldMask{}
	{
		subMask := o.GetProbe().MakeDiffFieldMask(other.GetProbe())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationResponseActivated_FieldTerminalPath{selector: ActivationResponseActivated_FieldPathSelectorProbe})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationResponseActivated_FieldSubPath{selector: ActivationResponseActivated_FieldPathSelectorProbe, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetServiceAccountKey().MakeDiffFieldMask(other.GetServiceAccountKey())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ActivationResponseActivated_FieldTerminalPath{selector: ActivationResponseActivated_FieldPathSelectorServiceAccountKey})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ActivationResponseActivated_FieldSubPath{selector: ActivationResponseActivated_FieldPathSelectorServiceAccountKey, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ActivationResponse_Activated) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationResponse_Activated))
}

func (o *ActivationResponse_Activated) Clone() *ActivationResponse_Activated {
	if o == nil {
		return nil
	}
	result := &ActivationResponse_Activated{}
	result.Probe = o.Probe.Clone()
	result.ServiceAccountKey = o.ServiceAccountKey.Clone()
	return result
}

func (o *ActivationResponse_Activated) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationResponse_Activated) Merge(source *ActivationResponse_Activated) {
	if source.GetProbe() != nil {
		if o.Probe == nil {
			o.Probe = new(probe.Probe)
		}
		o.Probe.Merge(source.GetProbe())
	}
	if source.GetServiceAccountKey() != nil {
		if o.ServiceAccountKey == nil {
			o.ServiceAccountKey = new(iam_service_account_key.ServiceAccountKey)
		}
		o.ServiceAccountKey.Merge(source.GetServiceAccountKey())
	}
}

func (o *ActivationResponse_Activated) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationResponse_Activated))
}

func (o *ActivationResponse_ConfirmActivationAck) GotenObjectExt() {}

func (o *ActivationResponse_ConfirmActivationAck) MakeFullFieldMask() *ActivationResponse_ConfirmActivationAck_FieldMask {
	return FullActivationResponse_ConfirmActivationAck_FieldMask()
}

func (o *ActivationResponse_ConfirmActivationAck) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActivationResponse_ConfirmActivationAck_FieldMask()
}

func (o *ActivationResponse_ConfirmActivationAck) MakeDiffFieldMask(other *ActivationResponse_ConfirmActivationAck) *ActivationResponse_ConfirmActivationAck_FieldMask {
	if o == nil && other == nil {
		return &ActivationResponse_ConfirmActivationAck_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActivationResponse_ConfirmActivationAck_FieldMask()
	}

	res := &ActivationResponse_ConfirmActivationAck_FieldMask{}
	return res
}

func (o *ActivationResponse_ConfirmActivationAck) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ActivationResponse_ConfirmActivationAck))
}

func (o *ActivationResponse_ConfirmActivationAck) Clone() *ActivationResponse_ConfirmActivationAck {
	if o == nil {
		return nil
	}
	result := &ActivationResponse_ConfirmActivationAck{}
	return result
}

func (o *ActivationResponse_ConfirmActivationAck) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ActivationResponse_ConfirmActivationAck) Merge(source *ActivationResponse_ConfirmActivationAck) {
}

func (o *ActivationResponse_ConfirmActivationAck) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ActivationResponse_ConfirmActivationAck))
}

func (o *SendActivationInvitationRequest) GotenObjectExt() {}

func (o *SendActivationInvitationRequest) MakeFullFieldMask() *SendActivationInvitationRequest_FieldMask {
	return FullSendActivationInvitationRequest_FieldMask()
}

func (o *SendActivationInvitationRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullSendActivationInvitationRequest_FieldMask()
}

func (o *SendActivationInvitationRequest) MakeDiffFieldMask(other *SendActivationInvitationRequest) *SendActivationInvitationRequest_FieldMask {
	if o == nil && other == nil {
		return &SendActivationInvitationRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullSendActivationInvitationRequest_FieldMask()
	}

	res := &SendActivationInvitationRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &SendActivationInvitationRequest_FieldTerminalPath{selector: SendActivationInvitationRequest_FieldPathSelectorName})
	}
	return res
}

func (o *SendActivationInvitationRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*SendActivationInvitationRequest))
}

func (o *SendActivationInvitationRequest) Clone() *SendActivationInvitationRequest {
	if o == nil {
		return nil
	}
	result := &SendActivationInvitationRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &probe.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *SendActivationInvitationRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *SendActivationInvitationRequest) Merge(source *SendActivationInvitationRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &probe.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *SendActivationInvitationRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*SendActivationInvitationRequest))
}

func (o *ResetActivationRequest) GotenObjectExt() {}

func (o *ResetActivationRequest) MakeFullFieldMask() *ResetActivationRequest_FieldMask {
	return FullResetActivationRequest_FieldMask()
}

func (o *ResetActivationRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullResetActivationRequest_FieldMask()
}

func (o *ResetActivationRequest) MakeDiffFieldMask(other *ResetActivationRequest) *ResetActivationRequest_FieldMask {
	if o == nil && other == nil {
		return &ResetActivationRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullResetActivationRequest_FieldMask()
	}

	res := &ResetActivationRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ResetActivationRequest_FieldTerminalPath{selector: ResetActivationRequest_FieldPathSelectorName})
	}
	return res
}

func (o *ResetActivationRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ResetActivationRequest))
}

func (o *ResetActivationRequest) Clone() *ResetActivationRequest {
	if o == nil {
		return nil
	}
	result := &ResetActivationRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &probe.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *ResetActivationRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ResetActivationRequest) Merge(source *ResetActivationRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &probe.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *ResetActivationRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ResetActivationRequest))
}

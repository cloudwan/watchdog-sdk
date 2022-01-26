// Code generated by protoc-gen-goten-resource
// Resource: QualityProfile
// DO NOT EDIT!!!

package quality_profile

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *QualityProfile_FieldMask) *QualityProfile_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "description")
		break
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &QualityProfile_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}

// Code generated by protoc-gen-goten-resource
// Resource change: QualityProfileChange
// DO NOT EDIT!!!

package quality_profile

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
	_ = &project.Project{}
)

func (c *QualityProfileChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*QualityProfileChange_Added_)
	return ok
}

func (c *QualityProfileChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*QualityProfileChange_Modified_)
	return ok
}

func (c *QualityProfileChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*QualityProfileChange_Current_)
	return ok
}

func (c *QualityProfileChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*QualityProfileChange_Removed_)
	return ok
}

func (c *QualityProfileChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *QualityProfileChange_Added_:
		return cType.Added.ViewIndex
	case *QualityProfileChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *QualityProfileChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *QualityProfileChange_Removed_:
		return cType.Removed.ViewIndex
	case *QualityProfileChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *QualityProfileChange) GetQualityProfile() *QualityProfile {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *QualityProfileChange_Added_:
		return cType.Added.QualityProfile
	case *QualityProfileChange_Modified_:
		return cType.Modified.QualityProfile
	case *QualityProfileChange_Current_:
		return cType.Current.QualityProfile
	case *QualityProfileChange_Removed_:
		return nil
	}
	return nil
}

func (c *QualityProfileChange) GetRawResource() gotenresource.Resource {
	return c.GetQualityProfile()
}

func (c *QualityProfileChange) GetQualityProfileName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *QualityProfileChange_Added_:
		return cType.Added.QualityProfile.GetName()
	case *QualityProfileChange_Modified_:
		return cType.Modified.Name
	case *QualityProfileChange_Current_:
		return cType.Current.QualityProfile.GetName()
	case *QualityProfileChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *QualityProfileChange) GetRawName() gotenresource.Name {
	return c.GetQualityProfileName()
}

func (c *QualityProfileChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &QualityProfileChange_Added_{
		Added: &QualityProfileChange_Added{
			QualityProfile: snapshot.(*QualityProfile),
			ViewIndex:      int32(idx),
		},
	}
}

func (c *QualityProfileChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &QualityProfileChange_Modified_{
		Modified: &QualityProfileChange_Modified{
			Name:              name.(*Name),
			QualityProfile:    snapshot.(*QualityProfile),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *QualityProfileChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &QualityProfileChange_Current_{
		Current: &QualityProfileChange_Current{
			QualityProfile: snapshot.(*QualityProfile),
		},
	}
}

func (c *QualityProfileChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &QualityProfileChange_Removed_{
		Removed: &QualityProfileChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

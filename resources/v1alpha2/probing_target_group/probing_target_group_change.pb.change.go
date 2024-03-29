// Code generated by protoc-gen-goten-resource
// Resource change: ProbingTargetGroupChange
// DO NOT EDIT!!!

package probing_target_group

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
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
	_ = &project.Project{}
)

func (c *ProbingTargetGroupChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetGroupChange_Added_)
	return ok
}

func (c *ProbingTargetGroupChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetGroupChange_Modified_)
	return ok
}

func (c *ProbingTargetGroupChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetGroupChange_Current_)
	return ok
}

func (c *ProbingTargetGroupChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetGroupChange_Removed_)
	return ok
}

func (c *ProbingTargetGroupChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetGroupChange_Added_:
		return cType.Added.ViewIndex
	case *ProbingTargetGroupChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProbingTargetGroupChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetGroupChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProbingTargetGroupChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProbingTargetGroupChange) GetProbingTargetGroup() *ProbingTargetGroup {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetGroupChange_Added_:
		return cType.Added.ProbingTargetGroup
	case *ProbingTargetGroupChange_Modified_:
		return cType.Modified.ProbingTargetGroup
	case *ProbingTargetGroupChange_Current_:
		return cType.Current.ProbingTargetGroup
	case *ProbingTargetGroupChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProbingTargetGroupChange) GetRawResource() gotenresource.Resource {
	return c.GetProbingTargetGroup()
}

func (c *ProbingTargetGroupChange) GetProbingTargetGroupName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetGroupChange_Added_:
		return cType.Added.ProbingTargetGroup.GetName()
	case *ProbingTargetGroupChange_Modified_:
		return cType.Modified.Name
	case *ProbingTargetGroupChange_Current_:
		return cType.Current.ProbingTargetGroup.GetName()
	case *ProbingTargetGroupChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProbingTargetGroupChange) GetRawName() gotenresource.Name {
	return c.GetProbingTargetGroupName()
}

func (c *ProbingTargetGroupChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProbingTargetGroupChange_Added_{
		Added: &ProbingTargetGroupChange_Added{
			ProbingTargetGroup: snapshot.(*ProbingTargetGroup),
			ViewIndex:          int32(idx),
		},
	}
}

func (c *ProbingTargetGroupChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProbingTargetGroupChange_Modified_{
		Modified: &ProbingTargetGroupChange_Modified{
			Name:               name.(*Name),
			ProbingTargetGroup: snapshot.(*ProbingTargetGroup),
			PreviousViewIndex:  int32(prevIdx),
			ViewIndex:          int32(newIdx),
		},
	}
}

func (c *ProbingTargetGroupChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProbingTargetGroupChange_Current_{
		Current: &ProbingTargetGroupChange_Current{
			ProbingTargetGroup: snapshot.(*ProbingTargetGroup),
		},
	}
}

func (c *ProbingTargetGroupChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProbingTargetGroupChange_Removed_{
		Removed: &ProbingTargetGroupChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

// Code generated by protoc-gen-goten-resource
// Resource change: ProbingTargetChange
// DO NOT EDIT!!!

package probing_target

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

func (c *ProbingTargetChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetChange_Added_)
	return ok
}

func (c *ProbingTargetChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetChange_Modified_)
	return ok
}

func (c *ProbingTargetChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetChange_Current_)
	return ok
}

func (c *ProbingTargetChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingTargetChange_Removed_)
	return ok
}

func (c *ProbingTargetChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetChange_Added_:
		return cType.Added.ViewIndex
	case *ProbingTargetChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProbingTargetChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProbingTargetChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProbingTargetChange) GetProbingTarget() *ProbingTarget {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetChange_Added_:
		return cType.Added.ProbingTarget
	case *ProbingTargetChange_Modified_:
		return cType.Modified.ProbingTarget
	case *ProbingTargetChange_Current_:
		return cType.Current.ProbingTarget
	case *ProbingTargetChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProbingTargetChange) GetRawResource() gotenresource.Resource {
	return c.GetProbingTarget()
}

func (c *ProbingTargetChange) GetProbingTargetName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingTargetChange_Added_:
		return cType.Added.ProbingTarget.GetName()
	case *ProbingTargetChange_Modified_:
		return cType.Modified.Name
	case *ProbingTargetChange_Current_:
		return cType.Current.ProbingTarget.GetName()
	case *ProbingTargetChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProbingTargetChange) GetRawName() gotenresource.Name {
	return c.GetProbingTargetName()
}

func (c *ProbingTargetChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProbingTargetChange_Added_{
		Added: &ProbingTargetChange_Added{
			ProbingTarget: snapshot.(*ProbingTarget),
			ViewIndex:     int32(idx),
		},
	}
}

func (c *ProbingTargetChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProbingTargetChange_Modified_{
		Modified: &ProbingTargetChange_Modified{
			Name:              name.(*Name),
			ProbingTarget:     snapshot.(*ProbingTarget),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ProbingTargetChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProbingTargetChange_Current_{
		Current: &ProbingTargetChange_Current{
			ProbingTarget: snapshot.(*ProbingTarget),
		},
	}
}

func (c *ProbingTargetChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProbingTargetChange_Removed_{
		Removed: &ProbingTargetChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

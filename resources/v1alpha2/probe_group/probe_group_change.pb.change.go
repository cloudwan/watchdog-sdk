// Code generated by protoc-gen-goten-resource
// Resource change: ProbeGroupChange
// DO NOT EDIT!!!

package probe_group

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

func (c *ProbeGroupChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbeGroupChange_Added_)
	return ok
}

func (c *ProbeGroupChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbeGroupChange_Modified_)
	return ok
}

func (c *ProbeGroupChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbeGroupChange_Current_)
	return ok
}

func (c *ProbeGroupChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbeGroupChange_Removed_)
	return ok
}

func (c *ProbeGroupChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbeGroupChange_Added_:
		return cType.Added.ViewIndex
	case *ProbeGroupChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProbeGroupChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbeGroupChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProbeGroupChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProbeGroupChange) GetProbeGroup() *ProbeGroup {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbeGroupChange_Added_:
		return cType.Added.ProbeGroup
	case *ProbeGroupChange_Modified_:
		return cType.Modified.ProbeGroup
	case *ProbeGroupChange_Current_:
		return cType.Current.ProbeGroup
	case *ProbeGroupChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProbeGroupChange) GetRawResource() gotenresource.Resource {
	return c.GetProbeGroup()
}

func (c *ProbeGroupChange) GetProbeGroupName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbeGroupChange_Added_:
		return cType.Added.ProbeGroup.GetName()
	case *ProbeGroupChange_Modified_:
		return cType.Modified.Name
	case *ProbeGroupChange_Current_:
		return cType.Current.ProbeGroup.GetName()
	case *ProbeGroupChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProbeGroupChange) GetRawName() gotenresource.Name {
	return c.GetProbeGroupName()
}

func (c *ProbeGroupChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProbeGroupChange_Added_{
		Added: &ProbeGroupChange_Added{
			ProbeGroup: snapshot.(*ProbeGroup),
			ViewIndex:  int32(idx),
		},
	}
}

func (c *ProbeGroupChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProbeGroupChange_Modified_{
		Modified: &ProbeGroupChange_Modified{
			Name:              name.(*Name),
			ProbeGroup:        snapshot.(*ProbeGroup),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ProbeGroupChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProbeGroupChange_Current_{
		Current: &ProbeGroupChange_Current{
			ProbeGroup: snapshot.(*ProbeGroup),
		},
	}
}

func (c *ProbeGroupChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProbeGroupChange_Removed_{
		Removed: &ProbeGroupChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

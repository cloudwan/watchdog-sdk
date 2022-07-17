// Code generated by protoc-gen-goten-resource
// Resource change: ProbingSessionChange
// DO NOT EDIT!!!

package probing_session

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

func (c *ProbingSessionChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingSessionChange_Added_)
	return ok
}

func (c *ProbingSessionChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingSessionChange_Modified_)
	return ok
}

func (c *ProbingSessionChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingSessionChange_Current_)
	return ok
}

func (c *ProbingSessionChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingSessionChange_Removed_)
	return ok
}

func (c *ProbingSessionChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingSessionChange_Added_:
		return cType.Added.ViewIndex
	case *ProbingSessionChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProbingSessionChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingSessionChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProbingSessionChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProbingSessionChange) GetProbingSession() *ProbingSession {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingSessionChange_Added_:
		return cType.Added.ProbingSession
	case *ProbingSessionChange_Modified_:
		return cType.Modified.ProbingSession
	case *ProbingSessionChange_Current_:
		return cType.Current.ProbingSession
	case *ProbingSessionChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProbingSessionChange) GetRawResource() gotenresource.Resource {
	return c.GetProbingSession()
}

func (c *ProbingSessionChange) GetProbingSessionName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingSessionChange_Added_:
		return cType.Added.ProbingSession.GetName()
	case *ProbingSessionChange_Modified_:
		return cType.Modified.Name
	case *ProbingSessionChange_Current_:
		return cType.Current.ProbingSession.GetName()
	case *ProbingSessionChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProbingSessionChange) GetRawName() gotenresource.Name {
	return c.GetProbingSessionName()
}

func (c *ProbingSessionChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProbingSessionChange_Added_{
		Added: &ProbingSessionChange_Added{
			ProbingSession: snapshot.(*ProbingSession),
			ViewIndex:      int32(idx),
		},
	}
}

func (c *ProbingSessionChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProbingSessionChange_Modified_{
		Modified: &ProbingSessionChange_Modified{
			Name:              name.(*Name),
			ProbingSession:    snapshot.(*ProbingSession),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ProbingSessionChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProbingSessionChange_Current_{
		Current: &ProbingSessionChange_Current{
			ProbingSession: snapshot.(*ProbingSession),
		},
	}
}

func (c *ProbingSessionChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProbingSessionChange_Removed_{
		Removed: &ProbingSessionChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

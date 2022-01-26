// Code generated by protoc-gen-goten-resource
// Resource change: ProbingDistributionChange
// DO NOT EDIT!!!

package probing_distribution

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

func (c *ProbingDistributionChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingDistributionChange_Added_)
	return ok
}

func (c *ProbingDistributionChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingDistributionChange_Modified_)
	return ok
}

func (c *ProbingDistributionChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingDistributionChange_Current_)
	return ok
}

func (c *ProbingDistributionChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProbingDistributionChange_Removed_)
	return ok
}

func (c *ProbingDistributionChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingDistributionChange_Added_:
		return cType.Added.ViewIndex
	case *ProbingDistributionChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProbingDistributionChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProbingDistributionChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProbingDistributionChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProbingDistributionChange) GetProbingDistribution() *ProbingDistribution {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingDistributionChange_Added_:
		return cType.Added.ProbingDistribution
	case *ProbingDistributionChange_Modified_:
		return cType.Modified.ProbingDistribution
	case *ProbingDistributionChange_Current_:
		return cType.Current.ProbingDistribution
	case *ProbingDistributionChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProbingDistributionChange) GetResource() gotenresource.Resource {
	return c.GetProbingDistribution()
}

func (c *ProbingDistributionChange) GetProbingDistributionName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProbingDistributionChange_Added_:
		return cType.Added.ProbingDistribution.GetName()
	case *ProbingDistributionChange_Modified_:
		return cType.Modified.Name
	case *ProbingDistributionChange_Current_:
		return cType.Current.ProbingDistribution.GetName()
	case *ProbingDistributionChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProbingDistributionChange) GetRawName() gotenresource.Name {
	return c.GetProbingDistributionName()
}

func (c *ProbingDistributionChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProbingDistributionChange_Added_{
		Added: &ProbingDistributionChange_Added{
			ProbingDistribution: snapshot.(*ProbingDistribution),
			ViewIndex:           int32(idx),
		},
	}
}

func (c *ProbingDistributionChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProbingDistributionChange_Modified_{
		Modified: &ProbingDistributionChange_Modified{
			Name:                name.(*Name),
			ProbingDistribution: snapshot.(*ProbingDistribution),
			PreviousViewIndex:   int32(prevIdx),
			ViewIndex:           int32(newIdx),
		},
	}
}

func (c *ProbingDistributionChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProbingDistributionChange_Current_{
		Current: &ProbingDistributionChange_Current{
			ProbingDistribution: snapshot.(*ProbingDistribution),
		},
	}
}

func (c *ProbingDistributionChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProbingDistributionChange_Removed_{
		Removed: &ProbingDistributionChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}

// Code generated by protoc-gen-goten-resource
// Resource: QualityProfile
// DO NOT EDIT!!!

package quality_profile

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
)

type QualityProfileList []*QualityProfile

func (l QualityProfileList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*QualityProfile))
}

func (l QualityProfileList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(QualityProfileList)...)
}

func (l QualityProfileList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l QualityProfileList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*QualityProfile)
}

func (l QualityProfileList) Length() int {
	return len(l)
}

type QualityProfileChangeList []*QualityProfileChange

func (l QualityProfileChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*QualityProfileChange))
}

func (l QualityProfileChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(QualityProfileChangeList)...)
}

func (l QualityProfileChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l QualityProfileChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*QualityProfileChange)
}

func (l QualityProfileChangeList) Length() int {
	return len(l)
}

type QualityProfileNameList []*Name

func (l QualityProfileNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l QualityProfileNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(QualityProfileNameList)...)
}

func (l QualityProfileNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l QualityProfileNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l QualityProfileNameList) Length() int {
	return len(l)
}

type QualityProfileReferenceList []*Reference

func (l QualityProfileReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l QualityProfileReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(QualityProfileReferenceList)...)
}

func (l QualityProfileReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l QualityProfileReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l QualityProfileReferenceList) Length() int {
	return len(l)
}

type QualityProfileParentNameList []*ParentName

func (l QualityProfileParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l QualityProfileParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(QualityProfileParentNameList)...)
}

func (l QualityProfileParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l QualityProfileParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l QualityProfileParentNameList) Length() int {
	return len(l)
}

type QualityProfileParentReferenceList []*ParentReference

func (l QualityProfileParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l QualityProfileParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(QualityProfileParentReferenceList)...)
}

func (l QualityProfileParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l QualityProfileParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l QualityProfileParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l QualityProfileParentReferenceList) Length() int {
	return len(l)
}

type QualityProfileMap map[Name]*QualityProfile

func (m QualityProfileMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m QualityProfileMap) Set(res gotenresource.Resource) {
	tRes := res.(*QualityProfile)
	m[*tRes.Name] = tRes
}

func (m QualityProfileMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m QualityProfileMap) Length() int {
	return len(m)
}

func (m QualityProfileMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type QualityProfileChangeMap map[Name]*QualityProfileChange

func (m QualityProfileChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m QualityProfileChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*QualityProfileChange)
	m[*tChange.GetQualityProfileName()] = tChange
}

func (m QualityProfileChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m QualityProfileChangeMap) Length() int {
	return len(m)
}

func (m QualityProfileChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

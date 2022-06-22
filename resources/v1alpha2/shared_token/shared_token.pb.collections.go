// Code generated by protoc-gen-goten-resource
// Resource: SharedToken
// DO NOT EDIT!!!

package shared_token

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

type SharedTokenList []*SharedToken

func (l SharedTokenList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*SharedToken))
}

func (l SharedTokenList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(SharedTokenList)...)
}

func (l SharedTokenList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l SharedTokenList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*SharedToken)
}

func (l SharedTokenList) Length() int {
	return len(l)
}

type SharedTokenChangeList []*SharedTokenChange

func (l SharedTokenChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*SharedTokenChange))
}

func (l SharedTokenChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(SharedTokenChangeList)...)
}

func (l SharedTokenChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l SharedTokenChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*SharedTokenChange)
}

func (l SharedTokenChangeList) Length() int {
	return len(l)
}

type SharedTokenNameList []*Name

func (l SharedTokenNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l SharedTokenNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(SharedTokenNameList)...)
}

func (l SharedTokenNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l SharedTokenNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l SharedTokenNameList) Length() int {
	return len(l)
}

type SharedTokenReferenceList []*Reference

func (l SharedTokenReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l SharedTokenReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(SharedTokenReferenceList)...)
}

func (l SharedTokenReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l SharedTokenReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l SharedTokenReferenceList) Length() int {
	return len(l)
}

type SharedTokenParentNameList []*ParentName

func (l SharedTokenParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l SharedTokenParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(SharedTokenParentNameList)...)
}

func (l SharedTokenParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l SharedTokenParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l SharedTokenParentNameList) Length() int {
	return len(l)
}

type SharedTokenParentReferenceList []*ParentReference

func (l SharedTokenParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l SharedTokenParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(SharedTokenParentReferenceList)...)
}

func (l SharedTokenParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l SharedTokenParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l SharedTokenParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l SharedTokenParentReferenceList) Length() int {
	return len(l)
}

type SharedTokenMap map[Name]*SharedToken

func (m SharedTokenMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m SharedTokenMap) Set(res gotenresource.Resource) {
	tRes := res.(*SharedToken)
	m[*tRes.Name] = tRes
}

func (m SharedTokenMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m SharedTokenMap) Length() int {
	return len(m)
}

func (m SharedTokenMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type SharedTokenChangeMap map[Name]*SharedTokenChange

func (m SharedTokenChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m SharedTokenChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*SharedTokenChange)
	m[*tChange.GetSharedTokenName()] = tChange
}

func (m SharedTokenChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m SharedTokenChangeMap) Length() int {
	return len(m)
}

func (m SharedTokenChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

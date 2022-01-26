// Code generated by protoc-gen-goten-resource
// Resource: InternetQualityRating
// DO NOT EDIT!!!

package internet_quality_rating

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probe"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &timestamp.Timestamp{}
	_ = &probe.Probe{}
)

type InternetQualityRatingList []*InternetQualityRating

func (l InternetQualityRatingList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*InternetQualityRating))
}

func (l InternetQualityRatingList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(InternetQualityRatingList)...)
}

func (l InternetQualityRatingList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l InternetQualityRatingList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*InternetQualityRating)
}

func (l InternetQualityRatingList) Length() int {
	return len(l)
}

type InternetQualityRatingChangeList []*InternetQualityRatingChange

func (l InternetQualityRatingChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*InternetQualityRatingChange))
}

func (l InternetQualityRatingChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(InternetQualityRatingChangeList)...)
}

func (l InternetQualityRatingChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l InternetQualityRatingChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*InternetQualityRatingChange)
}

func (l InternetQualityRatingChangeList) Length() int {
	return len(l)
}

type InternetQualityRatingNameList []*Name

func (l InternetQualityRatingNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l InternetQualityRatingNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(InternetQualityRatingNameList)...)
}

func (l InternetQualityRatingNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l InternetQualityRatingNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l InternetQualityRatingNameList) Length() int {
	return len(l)
}

type InternetQualityRatingReferenceList []*Reference

func (l InternetQualityRatingReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l InternetQualityRatingReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(InternetQualityRatingReferenceList)...)
}

func (l InternetQualityRatingReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l InternetQualityRatingReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l InternetQualityRatingReferenceList) Length() int {
	return len(l)
}

type InternetQualityRatingParentNameList []*ParentName

func (l InternetQualityRatingParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l InternetQualityRatingParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(InternetQualityRatingParentNameList)...)
}

func (l InternetQualityRatingParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l InternetQualityRatingParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l InternetQualityRatingParentNameList) Length() int {
	return len(l)
}

type InternetQualityRatingParentReferenceList []*ParentReference

func (l InternetQualityRatingParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l InternetQualityRatingParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(InternetQualityRatingParentReferenceList)...)
}

func (l InternetQualityRatingParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l InternetQualityRatingParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l InternetQualityRatingParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l InternetQualityRatingParentReferenceList) Length() int {
	return len(l)
}

type InternetQualityRatingMap map[Name]*InternetQualityRating

func (m InternetQualityRatingMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m InternetQualityRatingMap) Set(res gotenresource.Resource) {
	tRes := res.(*InternetQualityRating)
	m[*tRes.Name] = tRes
}

func (m InternetQualityRatingMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m InternetQualityRatingMap) Length() int {
	return len(m)
}

func (m InternetQualityRatingMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type InternetQualityRatingChangeMap map[Name]*InternetQualityRatingChange

func (m InternetQualityRatingChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m InternetQualityRatingChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*InternetQualityRatingChange)
	m[*tChange.GetInternetQualityRatingName()] = tChange
}

func (m InternetQualityRatingChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m InternetQualityRatingChangeMap) Length() int {
	return len(m)
}

func (m InternetQualityRatingChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

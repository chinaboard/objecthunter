package objecthunter

import (
	"github.com/chinaboard/objecthunter/oss"
)

type ObjectHandle func(properties oss.ObjectProperties) bool
type ListHandle func(properties oss.ListObjectsResultV2) bool
type HunterFilter struct {
	objectHandle ObjectHandle
	listHandle   ListHandle
}

func NewHunterFilter(oh ObjectHandle, lh ListHandle) *HunterFilter {
	if lh == nil {
		lh = defaultListHandle
	}
	if oh == nil {
		oh = defaultObjectHandle
	}
	return &HunterFilter{objectHandle: oh, listHandle: lh}
}

func defaultListHandle(oss.ListObjectsResultV2) bool {
	return true
}

func defaultObjectHandle(oss.ObjectProperties) bool {
	return true
}

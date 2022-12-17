package objecthunter

import (
	"fmt"
	"github.com/chinaboard/objecthunter/oss"
)

type ObjectHandle func(properties oss.ObjectProperties) bool

type ListHandle func(properties oss.ListObjectsResultV2) bool

type Options map[string]string

type HunterFilter struct {
	objectHandle ObjectHandle
	listHandle   ListHandle
	parameters   string
}

func NewHunterFilter(oh ObjectHandle, lh ListHandle, opt Options) *HunterFilter {
	if lh == nil {
		lh = defaultListHandle
	}
	if oh == nil {
		oh = defaultObjectHandle
	}
	if opt == nil {
		opt = Options{"max-keys": "100"}
	}
	parameters := ""
	for k, v := range opt {
		parameters += fmt.Sprintf("&%s=%s", k, v)
	}
	return &HunterFilter{objectHandle: oh, listHandle: lh, parameters: parameters}
}

func defaultListHandle(oss.ListObjectsResultV2) bool {
	return true
}

func defaultObjectHandle(oss.ObjectProperties) bool {
	return true
}

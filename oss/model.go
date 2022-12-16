package oss

import (
	"encoding/xml"
	"time"
)

// Owner defines Bucket/Object's owner
type Owner struct {
	XMLName     xml.Name `xml:"Owner"`
	ID          string   `xml:"ID"`          // Owner ID
	DisplayName string   `xml:"DisplayName"` // Owner's display name
}

// ObjectProperties defines Objecct properties
type ObjectProperties struct {
	XMLName      xml.Name  `xml:"Contents"`
	Key          string    `xml:"Key"`                   // Object key
	Type         string    `xml:"Type"`                  // Object type
	Size         int64     `xml:"Size"`                  // Object size
	ETag         string    `xml:"ETag"`                  // Object ETag
	Owner        Owner     `xml:"Owner"`                 // Object owner information
	LastModified time.Time `xml:"LastModified"`          // Object last modified time
	StorageClass string    `xml:"StorageClass"`          // Object storage class (Standard, IA, Archive)
	RestoreInfo  string    `xml:"RestoreInfo,omitempty"` // Object restoreInfo
}

// ListObjectsResultV2 defines the result from ListObjectsV2 request
type ListObjectsResultV2 struct {
	XMLName               xml.Name           `xml:"ListBucketResult"`
	Prefix                string             `xml:"Prefix"`                // The object prefix
	StartAfter            string             `xml:"StartAfter"`            // the input StartAfter
	ContinuationToken     string             `xml:"ContinuationToken"`     // the input ContinuationToken
	MaxKeys               int                `xml:"MaxKeys"`               // Max keys to return
	Delimiter             string             `xml:"Delimiter"`             // The delimiter for grouping objects' name
	IsTruncated           bool               `xml:"IsTruncated"`           // Flag indicates if all results are returned (when it's false)
	NextContinuationToken string             `xml:"NextContinuationToken"` // The start point of the next NextContinuationToken
	Objects               []ObjectProperties `xml:"Contents"`              // Object list
	CommonPrefixes        []string           `xml:"CommonPrefixes>Prefix"` // You can think of commonprefixes as "folders" whose names end with the delimiter
	KeyCount              int                `xml:"KeyCount"`
}

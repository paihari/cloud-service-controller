// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// BucketCommandDescriptor Command descriptor for querylanguage BUCKET command.
type BucketCommandDescriptor struct {

	// Command fragment display string from user specified query string formatted by query builder.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Command fragment internal string from user specified query string formatted by query builder.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// querylanguage command designation for example; reporting vs filtering
	Category *string `mandatory:"false" json:"category"`

	// Fields referenced in command fragment from user specified query string.
	ReferencedFields []AbstractField `mandatory:"false" json:"referencedFields"`

	// Fields declared in command fragment from user specified query string.
	DeclaredFields []AbstractField `mandatory:"false" json:"declaredFields"`

	// number of auto calculated ranges to compute if specified.
	MaxBuckets *int `mandatory:"false" json:"maxBuckets"`

	// Size of each numeric range if specified. Data type should match numeric field data type specified in the query string.
	Span *float32 `mandatory:"false" json:"span"`

	// List of the specified numeric ranges.
	Ranges []BucketRange `mandatory:"false" json:"ranges"`

	// Default value to use in place of null if a result does not fit into any of the specified / calculated ranges.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m BucketCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m BucketCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m BucketCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m BucketCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m BucketCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m BucketCommandDescriptor) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BucketCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBucketCommandDescriptor BucketCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeBucketCommandDescriptor
	}{
		"BUCKET",
		(MarshalTypeBucketCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *BucketCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		MaxBuckets          *int            `json:"maxBuckets"`
		Span                *float32        `json:"span"`
		Ranges              []BucketRange   `json:"ranges"`
		DefaultValue        *string         `json:"defaultValue"`
		DisplayQueryString  *string         `json:"displayQueryString"`
		InternalQueryString *string         `json:"internalQueryString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Category = model.Category

	m.ReferencedFields = make([]AbstractField, len(model.ReferencedFields))
	for i, n := range model.ReferencedFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ReferencedFields[i] = nn.(AbstractField)
		} else {
			m.ReferencedFields[i] = nil
		}
	}

	m.DeclaredFields = make([]AbstractField, len(model.DeclaredFields))
	for i, n := range model.DeclaredFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeclaredFields[i] = nn.(AbstractField)
		} else {
			m.DeclaredFields[i] = nil
		}
	}

	m.MaxBuckets = model.MaxBuckets

	m.Span = model.Span

	m.Ranges = make([]BucketRange, len(model.Ranges))
	for i, n := range model.Ranges {
		m.Ranges[i] = n
	}

	m.DefaultValue = model.DefaultValue

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}

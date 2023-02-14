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

// EventStatsCommandDescriptor Command descriptor for querylanguage EVENTSTATS command.
type EventStatsCommandDescriptor struct {

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

	// Group by fields if specified in the query string.
	GroupByFields []AbstractField `mandatory:"false" json:"groupByFields"`

	// Statistical functions specified in the query string. Atleast 1 is required for a EVENTSTATS command.
	Functions []FunctionField `mandatory:"false" json:"functions"`
}

//GetDisplayQueryString returns DisplayQueryString
func (m EventStatsCommandDescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m EventStatsCommandDescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m EventStatsCommandDescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m EventStatsCommandDescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m EventStatsCommandDescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m EventStatsCommandDescriptor) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EventStatsCommandDescriptor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEventStatsCommandDescriptor EventStatsCommandDescriptor
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeEventStatsCommandDescriptor
	}{
		"EVENT_STATS",
		(MarshalTypeEventStatsCommandDescriptor)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *EventStatsCommandDescriptor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Category            *string         `json:"category"`
		ReferencedFields    []abstractfield `json:"referencedFields"`
		DeclaredFields      []abstractfield `json:"declaredFields"`
		GroupByFields       []abstractfield `json:"groupByFields"`
		Functions           []FunctionField `json:"functions"`
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

	m.GroupByFields = make([]AbstractField, len(model.GroupByFields))
	for i, n := range model.GroupByFields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.GroupByFields[i] = nn.(AbstractField)
		} else {
			m.GroupByFields[i] = nil
		}
	}

	m.Functions = make([]FunctionField, len(model.Functions))
	for i, n := range model.Functions {
		m.Functions[i] = n
	}

	m.DisplayQueryString = model.DisplayQueryString

	m.InternalQueryString = model.InternalQueryString

	return
}

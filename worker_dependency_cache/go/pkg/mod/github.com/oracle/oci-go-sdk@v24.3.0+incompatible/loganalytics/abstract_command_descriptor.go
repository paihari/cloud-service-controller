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

// AbstractCommandDescriptor Generic command descriptor defining all attributes common to all querylanguage commands for parse output.
type AbstractCommandDescriptor interface {

	// Command fragment display string from user specified query string formatted by query builder.
	GetDisplayQueryString() *string

	// Command fragment internal string from user specified query string formatted by query builder.
	GetInternalQueryString() *string

	// querylanguage command designation for example; reporting vs filtering
	GetCategory() *string

	// Fields referenced in command fragment from user specified query string.
	GetReferencedFields() []AbstractField

	// Fields declared in command fragment from user specified query string.
	GetDeclaredFields() []AbstractField
}

type abstractcommanddescriptor struct {
	JsonData            []byte
	DisplayQueryString  *string         `mandatory:"true" json:"displayQueryString"`
	InternalQueryString *string         `mandatory:"true" json:"internalQueryString"`
	Category            *string         `mandatory:"false" json:"category"`
	ReferencedFields    []AbstractField `mandatory:"false" json:"referencedFields"`
	DeclaredFields      []AbstractField `mandatory:"false" json:"declaredFields"`
	Name                string          `json:"name"`
}

// UnmarshalJSON unmarshals json
func (m *abstractcommanddescriptor) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractcommanddescriptor abstractcommanddescriptor
	s := struct {
		Model Unmarshalerabstractcommanddescriptor
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayQueryString = s.Model.DisplayQueryString
	m.InternalQueryString = s.Model.InternalQueryString
	m.Category = s.Model.Category
	m.ReferencedFields = s.Model.ReferencedFields
	m.DeclaredFields = s.Model.DeclaredFields
	m.Name = s.Model.Name

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractcommanddescriptor) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Name {
	case "TOP":
		mm := TopCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HIGHLIGHT":
		mm := HighlightCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MULTI_SEARCH":
		mm := MultiSearchCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATS":
		mm := StatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_COMPARE":
		mm := TimeCompareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAIL":
		mm := TailCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REGEX":
		mm := RegexCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELTA":
		mm := DeltaCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOOKUP":
		mm := LookupCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEMO_MODE":
		mm := DemoModeCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD_SUMMARY":
		mm := FieldSummaryCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EVENT_STATS":
		mm := EventStatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WHERE":
		mm := WhereCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_SPLIT":
		mm := ClusterSplitCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TIME_STATS":
		mm := TimeStatsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER":
		mm := ClusterCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_DETAILS":
		mm := ClusterDetailsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELETE":
		mm := DeleteCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER_COMPARE":
		mm := ClusterCompareCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH":
		mm := SearchCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUCKET":
		mm := BucketCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMMAND":
		mm := CommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISTINCT":
		mm := DistinctCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LINK":
		mm := LinkCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SORT":
		mm := SortCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTRACT":
		mm := ExtractCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BOTTOM":
		mm := BottomCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELDS":
		mm := FieldsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HIGHLIGHT_ROWS":
		mm := HighlightRowsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACRO":
		mm := MacroCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLASSIFY":
		mm := ClassifyCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LINK_DETAILS":
		mm := LinkDetailsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SEARCH_LOOKUP":
		mm := SearchLookupCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HEAD":
		mm := HeadCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ADD_FIELDS":
		mm := AddFieldsCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EVAL":
		mm := EvalCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RENAME":
		mm := RenameCommandDescriptor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayQueryString returns DisplayQueryString
func (m abstractcommanddescriptor) GetDisplayQueryString() *string {
	return m.DisplayQueryString
}

//GetInternalQueryString returns InternalQueryString
func (m abstractcommanddescriptor) GetInternalQueryString() *string {
	return m.InternalQueryString
}

//GetCategory returns Category
func (m abstractcommanddescriptor) GetCategory() *string {
	return m.Category
}

//GetReferencedFields returns ReferencedFields
func (m abstractcommanddescriptor) GetReferencedFields() []AbstractField {
	return m.ReferencedFields
}

//GetDeclaredFields returns DeclaredFields
func (m abstractcommanddescriptor) GetDeclaredFields() []AbstractField {
	return m.DeclaredFields
}

func (m abstractcommanddescriptor) String() string {
	return common.PointerString(m)
}

// AbstractCommandDescriptorNameEnum Enum with underlying type: string
type AbstractCommandDescriptorNameEnum string

// Set of constants representing the allowable values for AbstractCommandDescriptorNameEnum
const (
	AbstractCommandDescriptorNameCommand        AbstractCommandDescriptorNameEnum = "COMMAND"
	AbstractCommandDescriptorNameSearch         AbstractCommandDescriptorNameEnum = "SEARCH"
	AbstractCommandDescriptorNameStats          AbstractCommandDescriptorNameEnum = "STATS"
	AbstractCommandDescriptorNameTimeStats      AbstractCommandDescriptorNameEnum = "TIME_STATS"
	AbstractCommandDescriptorNameSort           AbstractCommandDescriptorNameEnum = "SORT"
	AbstractCommandDescriptorNameFields         AbstractCommandDescriptorNameEnum = "FIELDS"
	AbstractCommandDescriptorNameAddFields      AbstractCommandDescriptorNameEnum = "ADD_FIELDS"
	AbstractCommandDescriptorNameLink           AbstractCommandDescriptorNameEnum = "LINK"
	AbstractCommandDescriptorNameLinkDetails    AbstractCommandDescriptorNameEnum = "LINK_DETAILS"
	AbstractCommandDescriptorNameCluster        AbstractCommandDescriptorNameEnum = "CLUSTER"
	AbstractCommandDescriptorNameClusterDetails AbstractCommandDescriptorNameEnum = "CLUSTER_DETAILS"
	AbstractCommandDescriptorNameCuslterSplit   AbstractCommandDescriptorNameEnum = "CUSLTER_SPLIT"
	AbstractCommandDescriptorNameEval           AbstractCommandDescriptorNameEnum = "EVAL"
	AbstractCommandDescriptorNameExtract        AbstractCommandDescriptorNameEnum = "EXTRACT"
	AbstractCommandDescriptorNameEventStats     AbstractCommandDescriptorNameEnum = "EVENT_STATS"
	AbstractCommandDescriptorNameBucket         AbstractCommandDescriptorNameEnum = "BUCKET"
	AbstractCommandDescriptorNameClassify       AbstractCommandDescriptorNameEnum = "CLASSIFY"
	AbstractCommandDescriptorNameTop            AbstractCommandDescriptorNameEnum = "TOP"
	AbstractCommandDescriptorNameBottom         AbstractCommandDescriptorNameEnum = "BOTTOM"
	AbstractCommandDescriptorNameHead           AbstractCommandDescriptorNameEnum = "HEAD"
	AbstractCommandDescriptorNameTail           AbstractCommandDescriptorNameEnum = "TAIL"
	AbstractCommandDescriptorNameFieldSummary   AbstractCommandDescriptorNameEnum = "FIELD_SUMMARY"
	AbstractCommandDescriptorNameRegex          AbstractCommandDescriptorNameEnum = "REGEX"
	AbstractCommandDescriptorNameRename         AbstractCommandDescriptorNameEnum = "RENAME"
	AbstractCommandDescriptorNameTimeCompare    AbstractCommandDescriptorNameEnum = "TIME_COMPARE"
	AbstractCommandDescriptorNameWhere          AbstractCommandDescriptorNameEnum = "WHERE"
	AbstractCommandDescriptorNameClusterCompare AbstractCommandDescriptorNameEnum = "CLUSTER_COMPARE"
	AbstractCommandDescriptorNameDelete         AbstractCommandDescriptorNameEnum = "DELETE"
	AbstractCommandDescriptorNameDelta          AbstractCommandDescriptorNameEnum = "DELTA"
	AbstractCommandDescriptorNameDistinct       AbstractCommandDescriptorNameEnum = "DISTINCT"
	AbstractCommandDescriptorNameSearchLookup   AbstractCommandDescriptorNameEnum = "SEARCH_LOOKUP"
	AbstractCommandDescriptorNameLookup         AbstractCommandDescriptorNameEnum = "LOOKUP"
	AbstractCommandDescriptorNameDemoMode       AbstractCommandDescriptorNameEnum = "DEMO_MODE"
	AbstractCommandDescriptorNameMacro          AbstractCommandDescriptorNameEnum = "MACRO"
	AbstractCommandDescriptorNameMultiSearch    AbstractCommandDescriptorNameEnum = "MULTI_SEARCH"
	AbstractCommandDescriptorNameHighlight      AbstractCommandDescriptorNameEnum = "HIGHLIGHT"
	AbstractCommandDescriptorNameHighlightRows  AbstractCommandDescriptorNameEnum = "HIGHLIGHT_ROWS"
)

var mappingAbstractCommandDescriptorName = map[string]AbstractCommandDescriptorNameEnum{
	"COMMAND":         AbstractCommandDescriptorNameCommand,
	"SEARCH":          AbstractCommandDescriptorNameSearch,
	"STATS":           AbstractCommandDescriptorNameStats,
	"TIME_STATS":      AbstractCommandDescriptorNameTimeStats,
	"SORT":            AbstractCommandDescriptorNameSort,
	"FIELDS":          AbstractCommandDescriptorNameFields,
	"ADD_FIELDS":      AbstractCommandDescriptorNameAddFields,
	"LINK":            AbstractCommandDescriptorNameLink,
	"LINK_DETAILS":    AbstractCommandDescriptorNameLinkDetails,
	"CLUSTER":         AbstractCommandDescriptorNameCluster,
	"CLUSTER_DETAILS": AbstractCommandDescriptorNameClusterDetails,
	"CUSLTER_SPLIT":   AbstractCommandDescriptorNameCuslterSplit,
	"EVAL":            AbstractCommandDescriptorNameEval,
	"EXTRACT":         AbstractCommandDescriptorNameExtract,
	"EVENT_STATS":     AbstractCommandDescriptorNameEventStats,
	"BUCKET":          AbstractCommandDescriptorNameBucket,
	"CLASSIFY":        AbstractCommandDescriptorNameClassify,
	"TOP":             AbstractCommandDescriptorNameTop,
	"BOTTOM":          AbstractCommandDescriptorNameBottom,
	"HEAD":            AbstractCommandDescriptorNameHead,
	"TAIL":            AbstractCommandDescriptorNameTail,
	"FIELD_SUMMARY":   AbstractCommandDescriptorNameFieldSummary,
	"REGEX":           AbstractCommandDescriptorNameRegex,
	"RENAME":          AbstractCommandDescriptorNameRename,
	"TIME_COMPARE":    AbstractCommandDescriptorNameTimeCompare,
	"WHERE":           AbstractCommandDescriptorNameWhere,
	"CLUSTER_COMPARE": AbstractCommandDescriptorNameClusterCompare,
	"DELETE":          AbstractCommandDescriptorNameDelete,
	"DELTA":           AbstractCommandDescriptorNameDelta,
	"DISTINCT":        AbstractCommandDescriptorNameDistinct,
	"SEARCH_LOOKUP":   AbstractCommandDescriptorNameSearchLookup,
	"LOOKUP":          AbstractCommandDescriptorNameLookup,
	"DEMO_MODE":       AbstractCommandDescriptorNameDemoMode,
	"MACRO":           AbstractCommandDescriptorNameMacro,
	"MULTI_SEARCH":    AbstractCommandDescriptorNameMultiSearch,
	"HIGHLIGHT":       AbstractCommandDescriptorNameHighlight,
	"HIGHLIGHT_ROWS":  AbstractCommandDescriptorNameHighlightRows,
}

// GetAbstractCommandDescriptorNameEnumValues Enumerates the set of values for AbstractCommandDescriptorNameEnum
func GetAbstractCommandDescriptorNameEnumValues() []AbstractCommandDescriptorNameEnum {
	values := make([]AbstractCommandDescriptorNameEnum, 0)
	for _, v := range mappingAbstractCommandDescriptorName {
		values = append(values, v)
	}
	return values
}

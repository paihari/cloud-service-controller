// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogAnalyticsParserSummary LoganParserDetails
type LogAnalyticsParserSummary struct {

	// content
	Content *string `mandatory:"false" json:"content"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// encoding
	Encoding *string `mandatory:"false" json:"encoding"`

	// example content
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// fields Maps
	FieldMaps []LogAnalyticsParserField `mandatory:"false" json:"fieldMaps"`

	// footer regular expression
	FooterContent *string `mandatory:"false" json:"footerContent"`

	// header content
	HeaderContent *string `mandatory:"false" json:"headerContent"`

	// Name
	Name *string `mandatory:"false" json:"name"`

	// is default flag
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// is single line content
	IsSingleLineContent *bool `mandatory:"false" json:"isSingleLineContent"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// language
	Language *string `mandatory:"false" json:"language"`

	// last updated date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// log type test request version
	LogTypeTestRequestVersion *int `mandatory:"false" json:"logTypeTestRequestVersion"`

	// mapped parser list
	MappedParsers []LogAnalyticsParser `mandatory:"false" json:"mappedParsers"`

	// parser ignore line characters
	ParserIgnorelineCharacters *string `mandatory:"false" json:"parserIgnorelineCharacters"`

	// is hidden flag
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// sequence
	ParserSequence *int `mandatory:"false" json:"parserSequence"`

	// time zone
	ParserTimezone *string `mandatory:"false" json:"parserTimezone"`

	ParserFilter *LogAnalyticsParserFilter `mandatory:"false" json:"parserFilter"`

	// write once
	IsParserWrittenOnce *bool `mandatory:"false" json:"isParserWrittenOnce"`

	// plugin instance list
	ParserFunctions []LogAnalyticsParserFunction `mandatory:"false" json:"parserFunctions"`

	// sources using this parser
	SourcesCount *int64 `mandatory:"false" json:"sourcesCount"`

	// sources using list
	Sources []LogAnalyticsSource `mandatory:"false" json:"sources"`

	// tokenize original text
	ShouldTokenizeOriginalText *bool `mandatory:"false" json:"shouldTokenizeOriginalText"`

	// type
	Type LogAnalyticsParserSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// user deleted flag
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`
}

func (m LogAnalyticsParserSummary) String() string {
	return common.PointerString(m)
}

// LogAnalyticsParserSummaryTypeEnum Enum with underlying type: string
type LogAnalyticsParserSummaryTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsParserSummaryTypeEnum
const (
	LogAnalyticsParserSummaryTypeXml   LogAnalyticsParserSummaryTypeEnum = "XML"
	LogAnalyticsParserSummaryTypeJson  LogAnalyticsParserSummaryTypeEnum = "JSON"
	LogAnalyticsParserSummaryTypeRegex LogAnalyticsParserSummaryTypeEnum = "REGEX"
	LogAnalyticsParserSummaryTypeOdl   LogAnalyticsParserSummaryTypeEnum = "ODL"
)

var mappingLogAnalyticsParserSummaryType = map[string]LogAnalyticsParserSummaryTypeEnum{
	"XML":   LogAnalyticsParserSummaryTypeXml,
	"JSON":  LogAnalyticsParserSummaryTypeJson,
	"REGEX": LogAnalyticsParserSummaryTypeRegex,
	"ODL":   LogAnalyticsParserSummaryTypeOdl,
}

// GetLogAnalyticsParserSummaryTypeEnumValues Enumerates the set of values for LogAnalyticsParserSummaryTypeEnum
func GetLogAnalyticsParserSummaryTypeEnumValues() []LogAnalyticsParserSummaryTypeEnum {
	values := make([]LogAnalyticsParserSummaryTypeEnum, 0)
	for _, v := range mappingLogAnalyticsParserSummaryType {
		values = append(values, v)
	}
	return values
}

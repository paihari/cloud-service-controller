// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ForeignKey The foreign key object.
type ForeignKey struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// attributeRefs
	AttributeRefs []KeyAttribute `mandatory:"false" json:"attributeRefs"`

	// updateRule
	UpdateRule *int `mandatory:"false" json:"updateRule"`

	// deleteRule
	DeleteRule *int `mandatory:"false" json:"deleteRule"`

	ReferenceUniqueKey *UniqueKey `mandatory:"false" json:"referenceUniqueKey"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m ForeignKey) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ForeignKey) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeForeignKey ForeignKey
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeForeignKey
	}{
		"FOREIGN_KEY",
		(MarshalTypeForeignKey)(m),
	}

	return json.Marshal(&s)
}

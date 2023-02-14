// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Tag The tag use for filtering.
type Tag struct {

	// The tag namespace.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The key of the tag.
	Key *string `mandatory:"false" json:"key"`

	// The value of the tag.
	Value *string `mandatory:"false" json:"value"`
}

func (m Tag) String() string {
	return common.PointerString(m)
}

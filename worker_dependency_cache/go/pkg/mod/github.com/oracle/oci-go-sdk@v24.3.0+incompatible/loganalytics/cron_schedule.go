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

// CronSchedule Cron schedule for a scheduled task.
type CronSchedule struct {

	// Value in cron format.
	Expression *string `mandatory:"true" json:"expression"`

	// Time zone, by default UTC.
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// Schedule misfire retry policy.
	MisfirePolicy ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
}

//GetMisfirePolicy returns MisfirePolicy
func (m CronSchedule) GetMisfirePolicy() ScheduleMisfirePolicyEnum {
	return m.MisfirePolicy
}

func (m CronSchedule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CronSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCronSchedule CronSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCronSchedule
	}{
		"CRON",
		(MarshalTypeCronSchedule)(m),
	}

	return json.Marshal(&s)
}

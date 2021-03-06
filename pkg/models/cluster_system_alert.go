// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterSystemAlert Information about a system alert on an Elasticsearch cluster.
// swagger:model ClusterSystemAlert
type ClusterSystemAlert struct {

	// Type of system alert
	// Required: true
	// Enum: [automatic_restart heap_dump unknown_event]
	AlertType *string `json:"alert_type"`

	// The exit_code related to the event. Only applicable for alert_type: slain
	ExitCode int32 `json:"exit_code,omitempty"`

	// Instance that caused the system alert
	// Required: true
	InstanceName *string `json:"instance_name"`

	// Timestamp marking the system alert
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// The URL related to the event. Only applicable for alert_type: heap_dump
	URL string `json:"url,omitempty"`
}

// Validate validates this cluster system alert
func (m *ClusterSystemAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterSystemAlertTypeAlertTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["automatic_restart","heap_dump","unknown_event"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSystemAlertTypeAlertTypePropEnum = append(clusterSystemAlertTypeAlertTypePropEnum, v)
	}
}

const (

	// ClusterSystemAlertAlertTypeAutomaticRestart captures enum value "automatic_restart"
	ClusterSystemAlertAlertTypeAutomaticRestart string = "automatic_restart"

	// ClusterSystemAlertAlertTypeHeapDump captures enum value "heap_dump"
	ClusterSystemAlertAlertTypeHeapDump string = "heap_dump"

	// ClusterSystemAlertAlertTypeUnknownEvent captures enum value "unknown_event"
	ClusterSystemAlertAlertTypeUnknownEvent string = "unknown_event"
)

// prop value enum
func (m *ClusterSystemAlert) validateAlertTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterSystemAlertTypeAlertTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterSystemAlert) validateAlertType(formats strfmt.Registry) error {

	if err := validate.Required("alert_type", "body", m.AlertType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlertTypeEnum("alert_type", "body", *m.AlertType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSystemAlert) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instance_name", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSystemAlert) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSystemAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSystemAlert) UnmarshalBinary(b []byte) error {
	var res ClusterSystemAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

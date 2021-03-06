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

package deployment_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeploymentTemplatesV2Params creates a new GetDeploymentTemplatesV2Params object
// with the default values initialized.
func NewGetDeploymentTemplatesV2Params() *GetDeploymentTemplatesV2Params {
	var (
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(true)
	)
	return &GetDeploymentTemplatesV2Params{
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTemplatesV2ParamsWithTimeout creates a new GetDeploymentTemplatesV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTemplatesV2ParamsWithTimeout(timeout time.Duration) *GetDeploymentTemplatesV2Params {
	var (
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(true)
	)
	return &GetDeploymentTemplatesV2Params{
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		timeout: timeout,
	}
}

// NewGetDeploymentTemplatesV2ParamsWithContext creates a new GetDeploymentTemplatesV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTemplatesV2ParamsWithContext(ctx context.Context) *GetDeploymentTemplatesV2Params {
	var (
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(true)
	)
	return &GetDeploymentTemplatesV2Params{
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		Context: ctx,
	}
}

// NewGetDeploymentTemplatesV2ParamsWithHTTPClient creates a new GetDeploymentTemplatesV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTemplatesV2ParamsWithHTTPClient(client *http.Client) *GetDeploymentTemplatesV2Params {
	var (
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(true)
	)
	return &GetDeploymentTemplatesV2Params{
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,
		HTTPClient:                 client,
	}
}

/*GetDeploymentTemplatesV2Params contains all the parameters to send to the API endpoint
for the get deployment templates v2 operation typically these are written to a http.Request
*/
type GetDeploymentTemplatesV2Params struct {

	/*Metadata
	  An optional key/value pair in the form of (key:value) that will act as a filter and exclude any templates that do not have a matching metadata item associated.

	*/
	Metadata *string
	/*Region
	  Region of the deployment templates

	*/
	Region string
	/*ShowHidden
	  If true, templates flagged as hidden will be returned.

	*/
	ShowHidden *bool
	/*ShowInstanceConfigurations
	  If true, will return details for each instance configuration referenced by the template.

	*/
	ShowInstanceConfigurations *bool
	/*StackVersion
	  If present, it will cause the returned deployment templates to be adapted to return only the elements allowed in that version.

	*/
	StackVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithTimeout(timeout time.Duration) *GetDeploymentTemplatesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithContext(ctx context.Context) *GetDeploymentTemplatesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithHTTPClient(client *http.Client) *GetDeploymentTemplatesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithMetadata(metadata *string) *GetDeploymentTemplatesV2Params {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetMetadata(metadata *string) {
	o.Metadata = metadata
}

// WithRegion adds the region to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithRegion(region string) *GetDeploymentTemplatesV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetRegion(region string) {
	o.Region = region
}

// WithShowHidden adds the showHidden to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithShowHidden(showHidden *bool) *GetDeploymentTemplatesV2Params {
	o.SetShowHidden(showHidden)
	return o
}

// SetShowHidden adds the showHidden to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetShowHidden(showHidden *bool) {
	o.ShowHidden = showHidden
}

// WithShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithShowInstanceConfigurations(showInstanceConfigurations *bool) *GetDeploymentTemplatesV2Params {
	o.SetShowInstanceConfigurations(showInstanceConfigurations)
	return o
}

// SetShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetShowInstanceConfigurations(showInstanceConfigurations *bool) {
	o.ShowInstanceConfigurations = showInstanceConfigurations
}

// WithStackVersion adds the stackVersion to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) WithStackVersion(stackVersion *string) *GetDeploymentTemplatesV2Params {
	o.SetStackVersion(stackVersion)
	return o
}

// SetStackVersion adds the stackVersion to the get deployment templates v2 params
func (o *GetDeploymentTemplatesV2Params) SetStackVersion(stackVersion *string) {
	o.StackVersion = stackVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTemplatesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string
		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {
			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
				return err
			}
		}

	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {
		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if o.ShowHidden != nil {

		// query param show_hidden
		var qrShowHidden bool
		if o.ShowHidden != nil {
			qrShowHidden = *o.ShowHidden
		}
		qShowHidden := swag.FormatBool(qrShowHidden)
		if qShowHidden != "" {
			if err := r.SetQueryParam("show_hidden", qShowHidden); err != nil {
				return err
			}
		}

	}

	if o.ShowInstanceConfigurations != nil {

		// query param show_instance_configurations
		var qrShowInstanceConfigurations bool
		if o.ShowInstanceConfigurations != nil {
			qrShowInstanceConfigurations = *o.ShowInstanceConfigurations
		}
		qShowInstanceConfigurations := swag.FormatBool(qrShowInstanceConfigurations)
		if qShowInstanceConfigurations != "" {
			if err := r.SetQueryParam("show_instance_configurations", qShowInstanceConfigurations); err != nil {
				return err
			}
		}

	}

	if o.StackVersion != nil {

		// query param stack_version
		var qrStackVersion string
		if o.StackVersion != nil {
			qrStackVersion = *o.StackVersion
		}
		qStackVersion := qrStackVersion
		if qStackVersion != "" {
			if err := r.SetQueryParam("stack_version", qStackVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

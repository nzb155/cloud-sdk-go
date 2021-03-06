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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteKibProxyRequestsParams creates a new DeleteKibProxyRequestsParams object
// with the default values initialized.
func NewDeleteKibProxyRequestsParams() *DeleteKibProxyRequestsParams {
	var ()
	return &DeleteKibProxyRequestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKibProxyRequestsParamsWithTimeout creates a new DeleteKibProxyRequestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKibProxyRequestsParamsWithTimeout(timeout time.Duration) *DeleteKibProxyRequestsParams {
	var ()
	return &DeleteKibProxyRequestsParams{

		timeout: timeout,
	}
}

// NewDeleteKibProxyRequestsParamsWithContext creates a new DeleteKibProxyRequestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteKibProxyRequestsParamsWithContext(ctx context.Context) *DeleteKibProxyRequestsParams {
	var ()
	return &DeleteKibProxyRequestsParams{

		Context: ctx,
	}
}

// NewDeleteKibProxyRequestsParamsWithHTTPClient creates a new DeleteKibProxyRequestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteKibProxyRequestsParamsWithHTTPClient(client *http.Client) *DeleteKibProxyRequestsParams {
	var ()
	return &DeleteKibProxyRequestsParams{
		HTTPClient: client,
	}
}

/*DeleteKibProxyRequestsParams contains all the parameters to send to the API endpoint
for the delete kib proxy requests operation typically these are written to a http.Request
*/
type DeleteKibProxyRequestsParams struct {

	/*XManagementRequest
	  When set to `true`, includes the X-Management-Request header value.

	*/
	XManagementRequest string
	/*ClusterID
	  The Kibana deployment identifier

	*/
	ClusterID string
	/*KibanaPath
	  The URL part to proxy to the Kibana cluster. Example: /api/spaces/space or /api/upgrade_assistant/status

	*/
	KibanaPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithTimeout(timeout time.Duration) *DeleteKibProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithContext(ctx context.Context) *DeleteKibProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithHTTPClient(client *http.Client) *DeleteKibProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *DeleteKibProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithClusterID adds the clusterID to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithClusterID(clusterID string) *DeleteKibProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithKibanaPath adds the kibanaPath to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) WithKibanaPath(kibanaPath string) *DeleteKibProxyRequestsParams {
	o.SetKibanaPath(kibanaPath)
	return o
}

// SetKibanaPath adds the kibanaPath to the delete kib proxy requests params
func (o *DeleteKibProxyRequestsParams) SetKibanaPath(kibanaPath string) {
	o.KibanaPath = kibanaPath
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKibProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param kibana_path
	if err := r.SetPathParam("kibana_path", o.KibanaPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

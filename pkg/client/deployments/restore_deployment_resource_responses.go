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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// RestoreDeploymentResourceReader is a Reader for the RestoreDeploymentResource structure.
type RestoreDeploymentResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDeploymentResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreDeploymentResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestoreDeploymentResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreDeploymentResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewRestoreDeploymentResourceRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestoreDeploymentResourceOK creates a RestoreDeploymentResourceOK with default headers values
func NewRestoreDeploymentResourceOK() *RestoreDeploymentResourceOK {
	return &RestoreDeploymentResourceOK{}
}

/*RestoreDeploymentResourceOK handles this case with default header values.

Standard Deployment Resource Crud Response
*/
type RestoreDeploymentResourceOK struct {
	Payload *models.DeploymentResourceCrudResponse
}

func (o *RestoreDeploymentResourceOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/_restore][%d] restoreDeploymentResourceOK  %+v", 200, o.Payload)
}

func (o *RestoreDeploymentResourceOK) GetPayload() *models.DeploymentResourceCrudResponse {
	return o.Payload
}

func (o *RestoreDeploymentResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentResourceBadRequest creates a RestoreDeploymentResourceBadRequest with default headers values
func NewRestoreDeploymentResourceBadRequest() *RestoreDeploymentResourceBadRequest {
	return &RestoreDeploymentResourceBadRequest{}
}

/*RestoreDeploymentResourceBadRequest handles this case with default header values.

* The Resource does not have a pending plan. (code: `deployments.resource_does_not_have_a_pending_plan`)
* The resource is not shut down. (code: `deployments.resource_not_shutdown`)
 */
type RestoreDeploymentResourceBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestoreDeploymentResourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/_restore][%d] restoreDeploymentResourceBadRequest  %+v", 400, o.Payload)
}

func (o *RestoreDeploymentResourceBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentResourceNotFound creates a RestoreDeploymentResourceNotFound with default headers values
func NewRestoreDeploymentResourceNotFound() *RestoreDeploymentResourceNotFound {
	return &RestoreDeploymentResourceNotFound{}
}

/*RestoreDeploymentResourceNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type RestoreDeploymentResourceNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestoreDeploymentResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/_restore][%d] restoreDeploymentResourceNotFound  %+v", 404, o.Payload)
}

func (o *RestoreDeploymentResourceNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentResourceRetryWith creates a RestoreDeploymentResourceRetryWith with default headers values
func NewRestoreDeploymentResourceRetryWith() *RestoreDeploymentResourceRetryWith {
	return &RestoreDeploymentResourceRetryWith{}
}

/*RestoreDeploymentResourceRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type RestoreDeploymentResourceRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestoreDeploymentResourceRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/_restore][%d] restoreDeploymentResourceRetryWith  %+v", 449, o.Payload)
}

func (o *RestoreDeploymentResourceRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentResourceRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

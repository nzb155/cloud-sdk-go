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

package deployments_notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateDeploymentNoteReader is a Reader for the CreateDeploymentNote structure.
type CreateDeploymentNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeploymentNoteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateDeploymentNoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateDeploymentNoteRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDeploymentNoteCreated creates a CreateDeploymentNoteCreated with default headers values
func NewCreateDeploymentNoteCreated() *CreateDeploymentNoteCreated {
	return &CreateDeploymentNoteCreated{}
}

/*CreateDeploymentNoteCreated handles this case with default header values.

List of deployment notes after the new deployment note has been added
*/
type CreateDeploymentNoteCreated struct {
	Payload *models.Notes
}

func (o *CreateDeploymentNoteCreated) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/notes][%d] createDeploymentNoteCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentNoteCreated) GetPayload() *models.Notes {
	return o.Payload
}

func (o *CreateDeploymentNoteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Notes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentNoteNotFound creates a CreateDeploymentNoteNotFound with default headers values
func NewCreateDeploymentNoteNotFound() *CreateDeploymentNoteNotFound {
	return &CreateDeploymentNoteNotFound{}
}

/*CreateDeploymentNoteNotFound handles this case with default header values.

The deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type CreateDeploymentNoteNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentNoteNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/notes][%d] createDeploymentNoteNotFound  %+v", 404, o.Payload)
}

func (o *CreateDeploymentNoteNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentNoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentNoteRetryWith creates a CreateDeploymentNoteRetryWith with default headers values
func NewCreateDeploymentNoteRetryWith() *CreateDeploymentNoteRetryWith {
	return &CreateDeploymentNoteRetryWith{}
}

/*CreateDeploymentNoteRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateDeploymentNoteRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentNoteRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/notes][%d] createDeploymentNoteRetryWith  %+v", 449, o.Payload)
}

func (o *CreateDeploymentNoteRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentNoteRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

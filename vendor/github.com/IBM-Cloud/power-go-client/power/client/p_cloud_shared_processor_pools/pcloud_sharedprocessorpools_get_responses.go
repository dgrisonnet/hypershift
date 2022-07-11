// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_shared_processor_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudSharedprocessorpoolsGetReader is a Reader for the PcloudSharedprocessorpoolsGet structure.
type PcloudSharedprocessorpoolsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSharedprocessorpoolsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSharedprocessorpoolsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSharedprocessorpoolsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSharedprocessorpoolsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudSharedprocessorpoolsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSharedprocessorpoolsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudSharedprocessorpoolsGetOK creates a PcloudSharedprocessorpoolsGetOK with default headers values
func NewPcloudSharedprocessorpoolsGetOK() *PcloudSharedprocessorpoolsGetOK {
	return &PcloudSharedprocessorpoolsGetOK{}
}

/* PcloudSharedprocessorpoolsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSharedprocessorpoolsGetOK struct {
	Payload *models.SharedProcessorPool
}

func (o *PcloudSharedprocessorpoolsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools/{shared_processor_pool_id}][%d] pcloudSharedprocessorpoolsGetOK  %+v", 200, o.Payload)
}
func (o *PcloudSharedprocessorpoolsGetOK) GetPayload() *models.SharedProcessorPool {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SharedProcessorPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetBadRequest creates a PcloudSharedprocessorpoolsGetBadRequest with default headers values
func NewPcloudSharedprocessorpoolsGetBadRequest() *PcloudSharedprocessorpoolsGetBadRequest {
	return &PcloudSharedprocessorpoolsGetBadRequest{}
}

/* PcloudSharedprocessorpoolsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSharedprocessorpoolsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudSharedprocessorpoolsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools/{shared_processor_pool_id}][%d] pcloudSharedprocessorpoolsGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudSharedprocessorpoolsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetUnauthorized creates a PcloudSharedprocessorpoolsGetUnauthorized with default headers values
func NewPcloudSharedprocessorpoolsGetUnauthorized() *PcloudSharedprocessorpoolsGetUnauthorized {
	return &PcloudSharedprocessorpoolsGetUnauthorized{}
}

/* PcloudSharedprocessorpoolsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSharedprocessorpoolsGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudSharedprocessorpoolsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools/{shared_processor_pool_id}][%d] pcloudSharedprocessorpoolsGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudSharedprocessorpoolsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetNotFound creates a PcloudSharedprocessorpoolsGetNotFound with default headers values
func NewPcloudSharedprocessorpoolsGetNotFound() *PcloudSharedprocessorpoolsGetNotFound {
	return &PcloudSharedprocessorpoolsGetNotFound{}
}

/* PcloudSharedprocessorpoolsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudSharedprocessorpoolsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudSharedprocessorpoolsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools/{shared_processor_pool_id}][%d] pcloudSharedprocessorpoolsGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudSharedprocessorpoolsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSharedprocessorpoolsGetInternalServerError creates a PcloudSharedprocessorpoolsGetInternalServerError with default headers values
func NewPcloudSharedprocessorpoolsGetInternalServerError() *PcloudSharedprocessorpoolsGetInternalServerError {
	return &PcloudSharedprocessorpoolsGetInternalServerError{}
}

/* PcloudSharedprocessorpoolsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSharedprocessorpoolsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudSharedprocessorpoolsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/shared-processor-pools/{shared_processor_pool_id}][%d] pcloudSharedprocessorpoolsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudSharedprocessorpoolsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSharedprocessorpoolsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
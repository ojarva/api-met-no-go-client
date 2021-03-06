// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ojarva/api-met-no-go-client/models"
)

// GetCompactReader is a Reader for the GetCompact structure.
type GetCompactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCompactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCompactOK creates a GetCompactOK with default headers values
func NewGetCompactOK() *GetCompactOK {
	return &GetCompactOK{}
}

/* GetCompactOK describes a response with status code 200, with default header values.

Success
*/
type GetCompactOK struct {
	Payload *models.METJSONForecast
}

func (o *GetCompactOK) Error() string {
	return fmt.Sprintf("[GET /compact][%d] getCompactOK  %+v", 200, o.Payload)
}
func (o *GetCompactOK) GetPayload() *models.METJSONForecast {
	return o.Payload
}

func (o *GetCompactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.METJSONForecast)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetClassicJSONReader is a Reader for the GetClassicJSON structure.
type GetClassicJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClassicJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClassicJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClassicJSONOK creates a GetClassicJSONOK with default headers values
func NewGetClassicJSONOK() *GetClassicJSONOK {
	return &GetClassicJSONOK{}
}

/* GetClassicJSONOK describes a response with status code 200, with default header values.

Success
*/
type GetClassicJSONOK struct {
	Payload string
}

func (o *GetClassicJSONOK) Error() string {
	return fmt.Sprintf("[GET /classic.json][%d] getClassicJsonOK  %+v", 200, o.Payload)
}
func (o *GetClassicJSONOK) GetPayload() string {
	return o.Payload
}

func (o *GetClassicJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

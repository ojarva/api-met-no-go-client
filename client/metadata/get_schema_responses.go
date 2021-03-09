// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSchemaReader is a Reader for the GetSchema structure.
type GetSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchemaOK creates a GetSchemaOK with default headers values
func NewGetSchemaOK() *GetSchemaOK {
	return &GetSchemaOK{}
}

/* GetSchemaOK describes a response with status code 200, with default header values.

OK
*/
type GetSchemaOK struct {
}

func (o *GetSchemaOK) Error() string {
	return fmt.Sprintf("[GET /schema][%d] getSchemaOK ", 200)
}

func (o *GetSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
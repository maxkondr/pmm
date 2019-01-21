// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

(empty)
*/
type GetServiceOK struct {
	Payload *GetServiceOKBody
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Services/Get][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceBody get service body
swagger:model GetServiceBody
*/
type GetServiceBody struct {

	// Unique Service identifier.
	ID string `json:"id,omitempty"`
}

// Validate validates this get service body
func (o *GetServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceBody) UnmarshalBinary(b []byte) error {
	var res GetServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBody get service o k body
swagger:model GetServiceOKBody
*/
type GetServiceOKBody struct {

	// mysql
	Mysql *GetServiceOKBodyMysql `json:"mysql,omitempty"`
}

// Validate validates this get service o k body
func (o *GetServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	if o.Mysql != nil {
		if err := o.Mysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOK" + "." + "mysql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMysql MySQLService represents MySQL-compatible Service configuration.
swagger:model GetServiceOKBodyMysql
*/
type GetServiceOKBodyMysql struct {

	// MySQL access address (DNS name or IP address).
	Address string `json:"address,omitempty"`

	// host node info
	HostNodeInfo *GetServiceOKBodyMysqlHostNodeInfo `json:"host_node_info,omitempty"`

	// Unique Service identifier.
	ID string `json:"id,omitempty"`

	// Unique user-defined Service name.
	Name string `json:"name,omitempty"`

	// MySQL access port.
	Port int64 `json:"port,omitempty"`

	// MySQL access UNIX socket path.
	UnixSocket string `json:"unix_socket,omitempty"`
}

// Validate validates this get service o k body mysql
func (o *GetServiceOKBodyMysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHostNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceOKBodyMysql) validateHostNodeInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.HostNodeInfo) { // not required
		return nil
	}

	if o.HostNodeInfo != nil {
		if err := o.HostNodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOK" + "." + "mysql" + "." + "host_node_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMysqlHostNodeInfo HostNodeInfo describes the way Service or Agent runs on Node.
swagger:model GetServiceOKBodyMysqlHostNodeInfo
*/
type GetServiceOKBodyMysqlHostNodeInfo struct {

	// Linux distribution used if any.
	Distro string `json:"distro,omitempty"`

	// Linux distribution version used if any.
	DistroVersion string `json:"distro_version,omitempty"`

	// Node identifier where Service or Agent runs.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get service o k body mysql host node info
func (o *GetServiceOKBodyMysqlHostNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMysqlHostNodeInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMysqlHostNodeInfo) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMysqlHostNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

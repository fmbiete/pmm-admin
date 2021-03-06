// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMongoDBExporterReader is a Reader for the AddMongoDBExporter structure.
type AddMongoDBExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMongoDBExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMongoDBExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBExporterOK creates a AddMongoDBExporterOK with default headers values
func NewAddMongoDBExporterOK() *AddMongoDBExporterOK {
	return &AddMongoDBExporterOK{}
}

/*AddMongoDBExporterOK handles this case with default header values.

A successful response.
*/
type AddMongoDBExporterOK struct {
	Payload *AddMongoDBExporterOKBody
}

func (o *AddMongoDBExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] addMongoDbExporterOk  %+v", 200, o.Payload)
}

func (o *AddMongoDBExporterOK) GetPayload() *AddMongoDBExporterOKBody {
	return o.Payload
}

func (o *AddMongoDBExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBExporterDefault creates a AddMongoDBExporterDefault with default headers values
func NewAddMongoDBExporterDefault(code int) *AddMongoDBExporterDefault {
	return &AddMongoDBExporterDefault{
		_statusCode: code,
	}
}

/*AddMongoDBExporterDefault handles this case with default header values.

An unexpected error response
*/
type AddMongoDBExporterDefault struct {
	_statusCode int

	Payload *AddMongoDBExporterDefaultBody
}

// Code gets the status code for the add mongo DB exporter default response
func (o *AddMongoDBExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] AddMongoDBExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddMongoDBExporterDefault) GetPayload() *AddMongoDBExporterDefaultBody {
	return o.Payload
}

func (o *AddMongoDBExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMongoDBExporterBody add mongo DB exporter body
swagger:model AddMongoDBExporterBody
*/
type AddMongoDBExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// MongoDB password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`
}

// Validate validates this add mongo DB exporter body
func (o *AddMongoDBExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterDefaultBody add mongo DB exporter default body
swagger:model AddMongoDBExporterDefaultBody
*/
type AddMongoDBExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add mongo DB exporter default body
func (o *AddMongoDBExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterOKBody add mongo DB exporter OK body
swagger:model AddMongoDBExporterOKBody
*/
type AddMongoDBExporterOKBody struct {

	// mongodb exporter
	MongodbExporter *AddMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`
}

// Validate validates this add mongo DB exporter OK body
func (o *AddMongoDBExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model AddMongoDBExporterOKBodyMongodbExporter
*/
type AddMongoDBExporterOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this add mongo DB exporter OK body mongodb exporter
func (o *AddMongoDBExporterOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum = append(addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	AddMongoDBExporterOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	AddMongoDBExporterOKBodyMongodbExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DetailsItems0) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(o.Value) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

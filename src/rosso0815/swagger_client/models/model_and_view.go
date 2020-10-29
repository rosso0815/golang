// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelAndView ModelAndView
// swagger:model ModelAndView
type ModelAndView struct {

	// empty
	Empty bool `json:"empty,omitempty"`

	// model
	Model interface{} `json:"model,omitempty"`

	// model map
	ModelMap map[string]interface{} `json:"modelMap,omitempty"`

	// reference
	Reference bool `json:"reference,omitempty"`

	// status
	// Enum: [100 CONTINUE 101 SWITCHING_PROTOCOLS 102 PROCESSING 103 CHECKPOINT 200 OK 201 CREATED 202 ACCEPTED 203 NON_AUTHORITATIVE_INFORMATION 204 NO_CONTENT 205 RESET_CONTENT 206 PARTIAL_CONTENT 207 MULTI_STATUS 208 ALREADY_REPORTED 226 IM_USED 300 MULTIPLE_CHOICES 301 MOVED_PERMANENTLY 302 FOUND 302 MOVED_TEMPORARILY 303 SEE_OTHER 304 NOT_MODIFIED 305 USE_PROXY 307 TEMPORARY_REDIRECT 308 PERMANENT_REDIRECT 400 BAD_REQUEST 401 UNAUTHORIZED 402 PAYMENT_REQUIRED 403 FORBIDDEN 404 NOT_FOUND 405 METHOD_NOT_ALLOWED 406 NOT_ACCEPTABLE 407 PROXY_AUTHENTICATION_REQUIRED 408 REQUEST_TIMEOUT 409 CONFLICT 410 GONE 411 LENGTH_REQUIRED 412 PRECONDITION_FAILED 413 PAYLOAD_TOO_LARGE 413 REQUEST_ENTITY_TOO_LARGE 414 URI_TOO_LONG 414 REQUEST_URI_TOO_LONG 415 UNSUPPORTED_MEDIA_TYPE 416 REQUESTED_RANGE_NOT_SATISFIABLE 417 EXPECTATION_FAILED 418 I_AM_A_TEAPOT 419 INSUFFICIENT_SPACE_ON_RESOURCE 420 METHOD_FAILURE 421 DESTINATION_LOCKED 422 UNPROCESSABLE_ENTITY 423 LOCKED 424 FAILED_DEPENDENCY 426 UPGRADE_REQUIRED 428 PRECONDITION_REQUIRED 429 TOO_MANY_REQUESTS 431 REQUEST_HEADER_FIELDS_TOO_LARGE 451 UNAVAILABLE_FOR_LEGAL_REASONS 500 INTERNAL_SERVER_ERROR 501 NOT_IMPLEMENTED 502 BAD_GATEWAY 503 SERVICE_UNAVAILABLE 504 GATEWAY_TIMEOUT 505 HTTP_VERSION_NOT_SUPPORTED 506 VARIANT_ALSO_NEGOTIATES 507 INSUFFICIENT_STORAGE 508 LOOP_DETECTED 509 BANDWIDTH_LIMIT_EXCEEDED 510 NOT_EXTENDED 511 NETWORK_AUTHENTICATION_REQUIRED]
	Status string `json:"status,omitempty"`

	// view
	View *View `json:"view,omitempty"`

	// view name
	ViewName string `json:"viewName,omitempty"`
}

// Validate validates this model and view
func (m *ModelAndView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateView(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modelAndViewTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["100 CONTINUE","101 SWITCHING_PROTOCOLS","102 PROCESSING","103 CHECKPOINT","200 OK","201 CREATED","202 ACCEPTED","203 NON_AUTHORITATIVE_INFORMATION","204 NO_CONTENT","205 RESET_CONTENT","206 PARTIAL_CONTENT","207 MULTI_STATUS","208 ALREADY_REPORTED","226 IM_USED","300 MULTIPLE_CHOICES","301 MOVED_PERMANENTLY","302 FOUND","302 MOVED_TEMPORARILY","303 SEE_OTHER","304 NOT_MODIFIED","305 USE_PROXY","307 TEMPORARY_REDIRECT","308 PERMANENT_REDIRECT","400 BAD_REQUEST","401 UNAUTHORIZED","402 PAYMENT_REQUIRED","403 FORBIDDEN","404 NOT_FOUND","405 METHOD_NOT_ALLOWED","406 NOT_ACCEPTABLE","407 PROXY_AUTHENTICATION_REQUIRED","408 REQUEST_TIMEOUT","409 CONFLICT","410 GONE","411 LENGTH_REQUIRED","412 PRECONDITION_FAILED","413 PAYLOAD_TOO_LARGE","413 REQUEST_ENTITY_TOO_LARGE","414 URI_TOO_LONG","414 REQUEST_URI_TOO_LONG","415 UNSUPPORTED_MEDIA_TYPE","416 REQUESTED_RANGE_NOT_SATISFIABLE","417 EXPECTATION_FAILED","418 I_AM_A_TEAPOT","419 INSUFFICIENT_SPACE_ON_RESOURCE","420 METHOD_FAILURE","421 DESTINATION_LOCKED","422 UNPROCESSABLE_ENTITY","423 LOCKED","424 FAILED_DEPENDENCY","426 UPGRADE_REQUIRED","428 PRECONDITION_REQUIRED","429 TOO_MANY_REQUESTS","431 REQUEST_HEADER_FIELDS_TOO_LARGE","451 UNAVAILABLE_FOR_LEGAL_REASONS","500 INTERNAL_SERVER_ERROR","501 NOT_IMPLEMENTED","502 BAD_GATEWAY","503 SERVICE_UNAVAILABLE","504 GATEWAY_TIMEOUT","505 HTTP_VERSION_NOT_SUPPORTED","506 VARIANT_ALSO_NEGOTIATES","507 INSUFFICIENT_STORAGE","508 LOOP_DETECTED","509 BANDWIDTH_LIMIT_EXCEEDED","510 NOT_EXTENDED","511 NETWORK_AUTHENTICATION_REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelAndViewTypeStatusPropEnum = append(modelAndViewTypeStatusPropEnum, v)
	}
}

const (

	// ModelAndViewStatusNr100CONTINUE captures enum value "100 CONTINUE"
	ModelAndViewStatusNr100CONTINUE string = "100 CONTINUE"

	// ModelAndViewStatusNr101SWITCHINGPROTOCOLS captures enum value "101 SWITCHING_PROTOCOLS"
	ModelAndViewStatusNr101SWITCHINGPROTOCOLS string = "101 SWITCHING_PROTOCOLS"

	// ModelAndViewStatusNr102PROCESSING captures enum value "102 PROCESSING"
	ModelAndViewStatusNr102PROCESSING string = "102 PROCESSING"

	// ModelAndViewStatusNr103CHECKPOINT captures enum value "103 CHECKPOINT"
	ModelAndViewStatusNr103CHECKPOINT string = "103 CHECKPOINT"

	// ModelAndViewStatusNr200OK captures enum value "200 OK"
	ModelAndViewStatusNr200OK string = "200 OK"

	// ModelAndViewStatusNr201CREATED captures enum value "201 CREATED"
	ModelAndViewStatusNr201CREATED string = "201 CREATED"

	// ModelAndViewStatusNr202ACCEPTED captures enum value "202 ACCEPTED"
	ModelAndViewStatusNr202ACCEPTED string = "202 ACCEPTED"

	// ModelAndViewStatusNr203NONAUTHORITATIVEINFORMATION captures enum value "203 NON_AUTHORITATIVE_INFORMATION"
	ModelAndViewStatusNr203NONAUTHORITATIVEINFORMATION string = "203 NON_AUTHORITATIVE_INFORMATION"

	// ModelAndViewStatusNr204NOCONTENT captures enum value "204 NO_CONTENT"
	ModelAndViewStatusNr204NOCONTENT string = "204 NO_CONTENT"

	// ModelAndViewStatusNr205RESETCONTENT captures enum value "205 RESET_CONTENT"
	ModelAndViewStatusNr205RESETCONTENT string = "205 RESET_CONTENT"

	// ModelAndViewStatusNr206PARTIALCONTENT captures enum value "206 PARTIAL_CONTENT"
	ModelAndViewStatusNr206PARTIALCONTENT string = "206 PARTIAL_CONTENT"

	// ModelAndViewStatusNr207MULTISTATUS captures enum value "207 MULTI_STATUS"
	ModelAndViewStatusNr207MULTISTATUS string = "207 MULTI_STATUS"

	// ModelAndViewStatusNr208ALREADYREPORTED captures enum value "208 ALREADY_REPORTED"
	ModelAndViewStatusNr208ALREADYREPORTED string = "208 ALREADY_REPORTED"

	// ModelAndViewStatusNr226IMUSED captures enum value "226 IM_USED"
	ModelAndViewStatusNr226IMUSED string = "226 IM_USED"

	// ModelAndViewStatusNr300MULTIPLECHOICES captures enum value "300 MULTIPLE_CHOICES"
	ModelAndViewStatusNr300MULTIPLECHOICES string = "300 MULTIPLE_CHOICES"

	// ModelAndViewStatusNr301MOVEDPERMANENTLY captures enum value "301 MOVED_PERMANENTLY"
	ModelAndViewStatusNr301MOVEDPERMANENTLY string = "301 MOVED_PERMANENTLY"

	// ModelAndViewStatusNr302FOUND captures enum value "302 FOUND"
	ModelAndViewStatusNr302FOUND string = "302 FOUND"

	// ModelAndViewStatusNr302MOVEDTEMPORARILY captures enum value "302 MOVED_TEMPORARILY"
	ModelAndViewStatusNr302MOVEDTEMPORARILY string = "302 MOVED_TEMPORARILY"

	// ModelAndViewStatusNr303SEEOTHER captures enum value "303 SEE_OTHER"
	ModelAndViewStatusNr303SEEOTHER string = "303 SEE_OTHER"

	// ModelAndViewStatusNr304NOTMODIFIED captures enum value "304 NOT_MODIFIED"
	ModelAndViewStatusNr304NOTMODIFIED string = "304 NOT_MODIFIED"

	// ModelAndViewStatusNr305USEPROXY captures enum value "305 USE_PROXY"
	ModelAndViewStatusNr305USEPROXY string = "305 USE_PROXY"

	// ModelAndViewStatusNr307TEMPORARYREDIRECT captures enum value "307 TEMPORARY_REDIRECT"
	ModelAndViewStatusNr307TEMPORARYREDIRECT string = "307 TEMPORARY_REDIRECT"

	// ModelAndViewStatusNr308PERMANENTREDIRECT captures enum value "308 PERMANENT_REDIRECT"
	ModelAndViewStatusNr308PERMANENTREDIRECT string = "308 PERMANENT_REDIRECT"

	// ModelAndViewStatusNr400BADREQUEST captures enum value "400 BAD_REQUEST"
	ModelAndViewStatusNr400BADREQUEST string = "400 BAD_REQUEST"

	// ModelAndViewStatusNr401UNAUTHORIZED captures enum value "401 UNAUTHORIZED"
	ModelAndViewStatusNr401UNAUTHORIZED string = "401 UNAUTHORIZED"

	// ModelAndViewStatusNr402PAYMENTREQUIRED captures enum value "402 PAYMENT_REQUIRED"
	ModelAndViewStatusNr402PAYMENTREQUIRED string = "402 PAYMENT_REQUIRED"

	// ModelAndViewStatusNr403FORBIDDEN captures enum value "403 FORBIDDEN"
	ModelAndViewStatusNr403FORBIDDEN string = "403 FORBIDDEN"

	// ModelAndViewStatusNr404NOTFOUND captures enum value "404 NOT_FOUND"
	ModelAndViewStatusNr404NOTFOUND string = "404 NOT_FOUND"

	// ModelAndViewStatusNr405METHODNOTALLOWED captures enum value "405 METHOD_NOT_ALLOWED"
	ModelAndViewStatusNr405METHODNOTALLOWED string = "405 METHOD_NOT_ALLOWED"

	// ModelAndViewStatusNr406NOTACCEPTABLE captures enum value "406 NOT_ACCEPTABLE"
	ModelAndViewStatusNr406NOTACCEPTABLE string = "406 NOT_ACCEPTABLE"

	// ModelAndViewStatusNr407PROXYAUTHENTICATIONREQUIRED captures enum value "407 PROXY_AUTHENTICATION_REQUIRED"
	ModelAndViewStatusNr407PROXYAUTHENTICATIONREQUIRED string = "407 PROXY_AUTHENTICATION_REQUIRED"

	// ModelAndViewStatusNr408REQUESTTIMEOUT captures enum value "408 REQUEST_TIMEOUT"
	ModelAndViewStatusNr408REQUESTTIMEOUT string = "408 REQUEST_TIMEOUT"

	// ModelAndViewStatusNr409CONFLICT captures enum value "409 CONFLICT"
	ModelAndViewStatusNr409CONFLICT string = "409 CONFLICT"

	// ModelAndViewStatusNr410GONE captures enum value "410 GONE"
	ModelAndViewStatusNr410GONE string = "410 GONE"

	// ModelAndViewStatusNr411LENGTHREQUIRED captures enum value "411 LENGTH_REQUIRED"
	ModelAndViewStatusNr411LENGTHREQUIRED string = "411 LENGTH_REQUIRED"

	// ModelAndViewStatusNr412PRECONDITIONFAILED captures enum value "412 PRECONDITION_FAILED"
	ModelAndViewStatusNr412PRECONDITIONFAILED string = "412 PRECONDITION_FAILED"

	// ModelAndViewStatusNr413PAYLOADTOOLARGE captures enum value "413 PAYLOAD_TOO_LARGE"
	ModelAndViewStatusNr413PAYLOADTOOLARGE string = "413 PAYLOAD_TOO_LARGE"

	// ModelAndViewStatusNr413REQUESTENTITYTOOLARGE captures enum value "413 REQUEST_ENTITY_TOO_LARGE"
	ModelAndViewStatusNr413REQUESTENTITYTOOLARGE string = "413 REQUEST_ENTITY_TOO_LARGE"

	// ModelAndViewStatusNr414URITOOLONG captures enum value "414 URI_TOO_LONG"
	ModelAndViewStatusNr414URITOOLONG string = "414 URI_TOO_LONG"

	// ModelAndViewStatusNr414REQUESTURITOOLONG captures enum value "414 REQUEST_URI_TOO_LONG"
	ModelAndViewStatusNr414REQUESTURITOOLONG string = "414 REQUEST_URI_TOO_LONG"

	// ModelAndViewStatusNr415UNSUPPORTEDMEDIATYPE captures enum value "415 UNSUPPORTED_MEDIA_TYPE"
	ModelAndViewStatusNr415UNSUPPORTEDMEDIATYPE string = "415 UNSUPPORTED_MEDIA_TYPE"

	// ModelAndViewStatusNr416REQUESTEDRANGENOTSATISFIABLE captures enum value "416 REQUESTED_RANGE_NOT_SATISFIABLE"
	ModelAndViewStatusNr416REQUESTEDRANGENOTSATISFIABLE string = "416 REQUESTED_RANGE_NOT_SATISFIABLE"

	// ModelAndViewStatusNr417EXPECTATIONFAILED captures enum value "417 EXPECTATION_FAILED"
	ModelAndViewStatusNr417EXPECTATIONFAILED string = "417 EXPECTATION_FAILED"

	// ModelAndViewStatusNr418IAMATEAPOT captures enum value "418 I_AM_A_TEAPOT"
	ModelAndViewStatusNr418IAMATEAPOT string = "418 I_AM_A_TEAPOT"

	// ModelAndViewStatusNr419INSUFFICIENTSPACEONRESOURCE captures enum value "419 INSUFFICIENT_SPACE_ON_RESOURCE"
	ModelAndViewStatusNr419INSUFFICIENTSPACEONRESOURCE string = "419 INSUFFICIENT_SPACE_ON_RESOURCE"

	// ModelAndViewStatusNr420METHODFAILURE captures enum value "420 METHOD_FAILURE"
	ModelAndViewStatusNr420METHODFAILURE string = "420 METHOD_FAILURE"

	// ModelAndViewStatusNr421DESTINATIONLOCKED captures enum value "421 DESTINATION_LOCKED"
	ModelAndViewStatusNr421DESTINATIONLOCKED string = "421 DESTINATION_LOCKED"

	// ModelAndViewStatusNr422UNPROCESSABLEENTITY captures enum value "422 UNPROCESSABLE_ENTITY"
	ModelAndViewStatusNr422UNPROCESSABLEENTITY string = "422 UNPROCESSABLE_ENTITY"

	// ModelAndViewStatusNr423LOCKED captures enum value "423 LOCKED"
	ModelAndViewStatusNr423LOCKED string = "423 LOCKED"

	// ModelAndViewStatusNr424FAILEDDEPENDENCY captures enum value "424 FAILED_DEPENDENCY"
	ModelAndViewStatusNr424FAILEDDEPENDENCY string = "424 FAILED_DEPENDENCY"

	// ModelAndViewStatusNr426UPGRADEREQUIRED captures enum value "426 UPGRADE_REQUIRED"
	ModelAndViewStatusNr426UPGRADEREQUIRED string = "426 UPGRADE_REQUIRED"

	// ModelAndViewStatusNr428PRECONDITIONREQUIRED captures enum value "428 PRECONDITION_REQUIRED"
	ModelAndViewStatusNr428PRECONDITIONREQUIRED string = "428 PRECONDITION_REQUIRED"

	// ModelAndViewStatusNr429TOOMANYREQUESTS captures enum value "429 TOO_MANY_REQUESTS"
	ModelAndViewStatusNr429TOOMANYREQUESTS string = "429 TOO_MANY_REQUESTS"

	// ModelAndViewStatusNr431REQUESTHEADERFIELDSTOOLARGE captures enum value "431 REQUEST_HEADER_FIELDS_TOO_LARGE"
	ModelAndViewStatusNr431REQUESTHEADERFIELDSTOOLARGE string = "431 REQUEST_HEADER_FIELDS_TOO_LARGE"

	// ModelAndViewStatusNr451UNAVAILABLEFORLEGALREASONS captures enum value "451 UNAVAILABLE_FOR_LEGAL_REASONS"
	ModelAndViewStatusNr451UNAVAILABLEFORLEGALREASONS string = "451 UNAVAILABLE_FOR_LEGAL_REASONS"

	// ModelAndViewStatusNr500INTERNALSERVERERROR captures enum value "500 INTERNAL_SERVER_ERROR"
	ModelAndViewStatusNr500INTERNALSERVERERROR string = "500 INTERNAL_SERVER_ERROR"

	// ModelAndViewStatusNr501NOTIMPLEMENTED captures enum value "501 NOT_IMPLEMENTED"
	ModelAndViewStatusNr501NOTIMPLEMENTED string = "501 NOT_IMPLEMENTED"

	// ModelAndViewStatusNr502BADGATEWAY captures enum value "502 BAD_GATEWAY"
	ModelAndViewStatusNr502BADGATEWAY string = "502 BAD_GATEWAY"

	// ModelAndViewStatusNr503SERVICEUNAVAILABLE captures enum value "503 SERVICE_UNAVAILABLE"
	ModelAndViewStatusNr503SERVICEUNAVAILABLE string = "503 SERVICE_UNAVAILABLE"

	// ModelAndViewStatusNr504GATEWAYTIMEOUT captures enum value "504 GATEWAY_TIMEOUT"
	ModelAndViewStatusNr504GATEWAYTIMEOUT string = "504 GATEWAY_TIMEOUT"

	// ModelAndViewStatusNr505HTTPVERSIONNOTSUPPORTED captures enum value "505 HTTP_VERSION_NOT_SUPPORTED"
	ModelAndViewStatusNr505HTTPVERSIONNOTSUPPORTED string = "505 HTTP_VERSION_NOT_SUPPORTED"

	// ModelAndViewStatusNr506VARIANTALSONEGOTIATES captures enum value "506 VARIANT_ALSO_NEGOTIATES"
	ModelAndViewStatusNr506VARIANTALSONEGOTIATES string = "506 VARIANT_ALSO_NEGOTIATES"

	// ModelAndViewStatusNr507INSUFFICIENTSTORAGE captures enum value "507 INSUFFICIENT_STORAGE"
	ModelAndViewStatusNr507INSUFFICIENTSTORAGE string = "507 INSUFFICIENT_STORAGE"

	// ModelAndViewStatusNr508LOOPDETECTED captures enum value "508 LOOP_DETECTED"
	ModelAndViewStatusNr508LOOPDETECTED string = "508 LOOP_DETECTED"

	// ModelAndViewStatusNr509BANDWIDTHLIMITEXCEEDED captures enum value "509 BANDWIDTH_LIMIT_EXCEEDED"
	ModelAndViewStatusNr509BANDWIDTHLIMITEXCEEDED string = "509 BANDWIDTH_LIMIT_EXCEEDED"

	// ModelAndViewStatusNr510NOTEXTENDED captures enum value "510 NOT_EXTENDED"
	ModelAndViewStatusNr510NOTEXTENDED string = "510 NOT_EXTENDED"

	// ModelAndViewStatusNr511NETWORKAUTHENTICATIONREQUIRED captures enum value "511 NETWORK_AUTHENTICATION_REQUIRED"
	ModelAndViewStatusNr511NETWORKAUTHENTICATIONREQUIRED string = "511 NETWORK_AUTHENTICATION_REQUIRED"
)

// prop value enum
func (m *ModelAndView) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, modelAndViewTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ModelAndView) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ModelAndView) validateView(formats strfmt.Registry) error {

	if swag.IsZero(m.View) { // not required
		return nil
	}

	if m.View != nil {
		if err := m.View.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("view")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelAndView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAndView) UnmarshalBinary(b []byte) error {
	var res ModelAndView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Timesheet Timesheet
// swagger:model Timesheet
type Timesheet struct {

	// finish
	// Format: date-time
	Finish strfmt.DateTime `json:"finish,omitempty"`

	// remark
	Remark string `json:"remark,omitempty"`

	// start
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// timesheet Id
	TimesheetID int64 `json:"timesheetId,omitempty"`
}

// Validate validates this timesheet
func (m *Timesheet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinish(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timesheet) validateFinish(formats strfmt.Registry) error {

	if swag.IsZero(m.Finish) { // not required
		return nil
	}

	if err := validate.FormatOf("finish", "body", "date-time", m.Finish.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Timesheet) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timesheet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timesheet) UnmarshalBinary(b []byte) error {
	var res Timesheet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

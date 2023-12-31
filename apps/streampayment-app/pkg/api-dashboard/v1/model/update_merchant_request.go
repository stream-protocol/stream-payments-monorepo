// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateMerchantRequest update merchant request
//
// swagger:model updateMerchantRequest
type UpdateMerchantRequest struct {

	// Name
	// Example: My Store
	// Required: true
	// Max Length: 128
	// Min Length: 2
	Name string `json:"name"`

	// Website URL
	// Example: https://my-store.com
	// Required: true
	// Max Length: 255
	// Min Length: 2
	Website string `json:"website"`
}

// Validate validates this update merchant request
func (m *UpdateMerchantRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMerchantRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 128); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMerchantRequest) validateWebsite(formats strfmt.Registry) error {

	if err := validate.RequiredString("website", "body", m.Website); err != nil {
		return err
	}

	if err := validate.MinLength("website", "body", m.Website, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("website", "body", m.Website, 255); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update merchant request based on context it is used
func (m *UpdateMerchantRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMerchantRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMerchantRequest) UnmarshalBinary(b []byte) error {
	var res UpdateMerchantRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

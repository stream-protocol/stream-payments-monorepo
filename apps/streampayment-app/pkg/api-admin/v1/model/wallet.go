// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Wallet wallet
//
// swagger:model wallet
type Wallet struct {

	// Address
	// Example: 11256099
	Address string `json:"address,omitempty"`

	// blockchain
	Blockchain Blockchain `json:"blockchain,omitempty"`

	// Created At
	// Example: 1656696522
	CreatedAtUnix int64 `json:"createdAtUnix,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// UUID
	// Example: 123e4567-e89b-12d3-a456-426655440000
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this wallet
func (m *Wallet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockchain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wallet) validateBlockchain(formats strfmt.Registry) error {
	if swag.IsZero(m.Blockchain) { // not required
		return nil
	}

	if err := m.Blockchain.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("blockchain")
		}
		return err
	}

	return nil
}

// ContextValidate validate this wallet based on the context it is used
func (m *Wallet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockchain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wallet) contextValidateBlockchain(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Blockchain.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("blockchain")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Wallet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wallet) UnmarshalBinary(b []byte) error {
	var res Wallet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	models2 "magma/orc8r/cloud/go/models"
	models3 "magma/orc8r/cloud/go/pluginimpl/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LteNetwork LTE Network spec
// swagger:model lte_network
type LteNetwork struct {

	// cellular
	// Required: true
	Cellular *NetworkCellularConfigs `json:"cellular"`

	// description
	// Required: true
	Description models2.NetworkDescription `json:"description"`

	// dns
	// Required: true
	DNS *models3.NetworkDNSConfig `json:"dns"`

	// features
	Features *models3.NetworkFeatures `json:"features,omitempty"`

	// id
	// Required: true
	ID models2.NetworkID `json:"id"`

	// name
	// Required: true
	Name models2.NetworkName `json:"name"`

	// subscriber config
	SubscriberConfig *NetworkSubscriberConfig `json:"subscriber_config,omitempty"`
}

// Validate validates this lte network
func (m *LteNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCellular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LteNetwork) validateCellular(formats strfmt.Registry) error {

	if err := validate.Required("cellular", "body", m.Cellular); err != nil {
		return err
	}

	if m.Cellular != nil {
		if err := m.Cellular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *LteNetwork) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *LteNetwork) validateDNS(formats strfmt.Registry) error {

	if err := validate.Required("dns", "body", m.DNS); err != nil {
		return err
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *LteNetwork) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *LteNetwork) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *LteNetwork) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *LteNetwork) validateSubscriberConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberConfig) { // not required
		return nil
	}

	if m.SubscriberConfig != nil {
		if err := m.SubscriberConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscriber_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LteNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LteNetwork) UnmarshalBinary(b []byte) error {
	var res LteNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

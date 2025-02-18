// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InfraEnvCreateParams infra env create params
//
// swagger:model infra-env-create-params
type InfraEnvCreateParams struct {

	// A comma-separated list of NTP sources (name or IP) going to be added to all the hosts.
	AdditionalNtpSources *string `json:"additional_ntp_sources,omitempty"`

	// Additional SSH public keys for debugging the installation.
	AdditionalSSHAuthorizedKeys *string `json:"additional_ssh_authorized_keys,omitempty"`

	// PEM-encoded X.509 certificate bundle. Hosts discovered by this
	// infra-env will trust the certificates in this bundle. Clusters formed
	// from the hosts discovered by this infra-env will also trust the
	// certificates in this bundle.
	// Max Length: 65535
	AdditionalTrustBundle string `json:"additional_trust_bundle,omitempty"`

	// If set, all hosts that register will be associated with the specified cluster.
	// Format: uuid
	ClusterID *strfmt.UUID `json:"cluster_id,omitempty"`

	// The CPU architecture of the image (x86_64/arm64/etc).
	// Enum: [x86_64 aarch64 arm64 ppc64le s390x]
	CPUArchitecture string `json:"cpu_architecture,omitempty"`

	// JSON formatted string containing the user overrides for the initial ignition config.
	IgnitionConfigOverride string `json:"ignition_config_override,omitempty"`

	// image type
	ImageType ImageType `json:"image_type,omitempty"`

	// kernel arguments
	KernelArguments KernelArguments `json:"kernel_arguments"`

	// Name of the infra-env.
	// Required: true
	Name *string `json:"name"`

	// Version of the OpenShift cluster (used to infer the RHCOS version - temporary until generic logic implemented).
	OpenshiftVersion string `json:"openshift_version,omitempty"`

	// proxy
	Proxy *Proxy `json:"proxy,omitempty" gorm:"embedded;embeddedPrefix:proxy_"`

	// The pull secret obtained from Red Hat OpenShift Cluster Manager at console.redhat.com/openshift/install/pull-secret.
	// Required: true
	PullSecret *string `json:"pull_secret"`

	// SSH public key for debugging the installation.
	SSHAuthorizedKey *string `json:"ssh_authorized_key,omitempty"`

	// static network config
	StaticNetworkConfig []*HostStaticNetworkConfig `json:"static_network_config"`
}

// Validate validates this infra env create params
func (m *InfraEnvCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalTrustBundle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernelArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraEnvCreateParams) validateAdditionalTrustBundle(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalTrustBundle) { // not required
		return nil
	}

	if err := validate.MaxLength("additional_trust_bundle", "body", m.AdditionalTrustBundle, 65535); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

var infraEnvCreateParamsTypeCPUArchitecturePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["x86_64","aarch64","arm64","ppc64le","s390x"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		infraEnvCreateParamsTypeCPUArchitecturePropEnum = append(infraEnvCreateParamsTypeCPUArchitecturePropEnum, v)
	}
}

const (

	// InfraEnvCreateParamsCPUArchitectureX8664 captures enum value "x86_64"
	InfraEnvCreateParamsCPUArchitectureX8664 string = "x86_64"

	// InfraEnvCreateParamsCPUArchitectureAarch64 captures enum value "aarch64"
	InfraEnvCreateParamsCPUArchitectureAarch64 string = "aarch64"

	// InfraEnvCreateParamsCPUArchitectureArm64 captures enum value "arm64"
	InfraEnvCreateParamsCPUArchitectureArm64 string = "arm64"

	// InfraEnvCreateParamsCPUArchitecturePpc64le captures enum value "ppc64le"
	InfraEnvCreateParamsCPUArchitecturePpc64le string = "ppc64le"

	// InfraEnvCreateParamsCPUArchitectureS390x captures enum value "s390x"
	InfraEnvCreateParamsCPUArchitectureS390x string = "s390x"
)

// prop value enum
func (m *InfraEnvCreateParams) validateCPUArchitectureEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, infraEnvCreateParamsTypeCPUArchitecturePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InfraEnvCreateParams) validateCPUArchitecture(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUArchitecture) { // not required
		return nil
	}

	// value enum
	if err := m.validateCPUArchitectureEnum("cpu_architecture", "body", m.CPUArchitecture); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateImageType(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageType) { // not required
		return nil
	}

	if err := m.ImageType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("image_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("image_type")
		}
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateKernelArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.KernelArguments) { // not required
		return nil
	}

	if err := m.KernelArguments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kernel_arguments")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kernel_arguments")
		}
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *InfraEnvCreateParams) validatePullSecret(formats strfmt.Registry) error {

	if err := validate.Required("pull_secret", "body", m.PullSecret); err != nil {
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) validateStaticNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticNetworkConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticNetworkConfig); i++ {
		if swag.IsZero(m.StaticNetworkConfig[i]) { // not required
			continue
		}

		if m.StaticNetworkConfig[i] != nil {
			if err := m.StaticNetworkConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("static_network_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("static_network_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this infra env create params based on the context it is used
func (m *InfraEnvCreateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKernelArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraEnvCreateParams) contextValidateImageType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ImageType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("image_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("image_type")
		}
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) contextValidateKernelArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := m.KernelArguments.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kernel_arguments")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kernel_arguments")
		}
		return err
	}

	return nil
}

func (m *InfraEnvCreateParams) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {
		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *InfraEnvCreateParams) contextValidateStaticNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StaticNetworkConfig); i++ {

		if m.StaticNetworkConfig[i] != nil {
			if err := m.StaticNetworkConfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("static_network_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("static_network_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfraEnvCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraEnvCreateParams) UnmarshalBinary(b []byte) error {
	var res InfraEnvCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

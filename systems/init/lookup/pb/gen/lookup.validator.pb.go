// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lookup.proto

package gen

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddOrgRequest) Validate() error {
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	if this.Certificate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Certificate", fmt.Errorf(`value '%v' must not be an empty string`, this.Certificate))
	}
	if this.Ip == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Ip", fmt.Errorf(`value '%v' must not be an empty string`, this.Ip))
	}
	return nil
}
func (this *AddOrgResponse) Validate() error {
	return nil
}
func (this *UpdateOrgRequest) Validate() error {
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *UpdateOrgResponse) Validate() error {
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *GetOrgRequest) Validate() error {
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *GetOrgResponse) Validate() error {
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *AddNodeRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	return nil
}
func (this *AddNodeResponse) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	return nil
}
func (this *GetNodeForOrgRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *GetNodeResponse) Validate() error {
	return nil
}
func (this *GetNodeRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	return nil
}
func (this *DeleteNodeRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *DeleteNodeResponse) Validate() error {
	return nil
}
func (this *GetSystemRequest) Validate() error {
	if this.SystemName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SystemName", fmt.Errorf(`value '%v' must not be an empty string`, this.SystemName))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *GetSystemResponse) Validate() error {
	return nil
}
func (this *AddSystemRequest) Validate() error {
	if this.SystemName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SystemName", fmt.Errorf(`value '%v' must not be an empty string`, this.SystemName))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	if this.Certificate == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Certificate", fmt.Errorf(`value '%v' must not be an empty string`, this.Certificate))
	}
	if this.Ip == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Ip", fmt.Errorf(`value '%v' must not be an empty string`, this.Ip))
	}
	return nil
}
func (this *AddSystemResponse) Validate() error {
	return nil
}
func (this *UpdateSystemRequest) Validate() error {
	if this.SystemName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SystemName", fmt.Errorf(`value '%v' must not be an empty string`, this.SystemName))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *UpdateSystemResponse) Validate() error {
	return nil
}
func (this *DeleteSystemRequest) Validate() error {
	if this.SystemName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SystemName", fmt.Errorf(`value '%v' must not be an empty string`, this.SystemName))
	}
	if this.OrgName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrgName", fmt.Errorf(`value '%v' must not be an empty string`, this.OrgName))
	}
	return nil
}
func (this *DeleteSystemResponse) Validate() error {
	return nil
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platformpb/platform.proto

package platformpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ConnectRequest) Validate() error {
	if this.ServerName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServerName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServerName))
	}
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	return nil
}
func (this *ConnectResponse) Validate() error {
	return nil
}
func (this *DisconnectRequest) Validate() error {
	return nil
}
func (this *DisconnectResponse) Validate() error {
	return nil
}
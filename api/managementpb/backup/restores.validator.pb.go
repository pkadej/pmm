// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/backup/restores.proto

package backupv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RestoreHistoryItem) Validate() error {
	if this.StartedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartedAt", err)
		}
	}
	if this.FinishedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.FinishedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("FinishedAt", err)
		}
	}
	return nil
}
func (this *ListRestoreHistoryRequest) Validate() error {
	return nil
}
func (this *ListRestoreHistoryResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}

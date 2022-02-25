// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/rds.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DiscoverRDSInstance) Validate() error {
	return nil
}
func (this *DiscoverRDSRequest) Validate() error {
	return nil
}
func (this *DiscoverRDSResponse) Validate() error {
	for _, item := range this.RdsInstances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RdsInstances", err)
			}
		}
	}
	return nil
}
func (this *AddRDSRequest) Validate() error {
	if this.Region == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Region", fmt.Errorf(`value '%v' must not be an empty string`, this.Region))
	}
	if this.InstanceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must not be an empty string`, this.InstanceId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddRDSResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	if this.RdsExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RdsExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RdsExporter", err)
		}
	}
	if this.Mysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
		}
	}
	if this.MysqldExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MysqldExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MysqldExporter", err)
		}
	}
	if this.QanMysqlPerfschema != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanMysqlPerfschema); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanMysqlPerfschema", err)
		}
	}
	if this.Postgresql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Postgresql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
		}
	}
	if this.PostgresqlExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PostgresqlExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PostgresqlExporter", err)
		}
	}
	if this.QanPostgresqlPgstatements != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanPostgresqlPgstatements); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanPostgresqlPgstatements", err)
		}
	}
	return nil
}

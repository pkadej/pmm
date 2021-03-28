// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/managementpb/dbaas/json/client/components"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/kubernetes"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/logs_api"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/psmdb_cluster"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/xtra_db_cluster"
)

// Default PMM d baa s HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new PMM d baa s HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PMMDBaaS {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new PMM d baa s HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PMMDBaaS {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new PMM d baa s client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PMMDBaaS {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PMMDBaaS)
	cli.Transport = transport
	cli.Components = components.New(transport, formats)
	cli.Kubernetes = kubernetes.New(transport, formats)
	cli.LogsAPI = logs_api.New(transport, formats)
	cli.PSMDBCluster = psmdb_cluster.New(transport, formats)
	cli.XtraDBCluster = xtra_db_cluster.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PMMDBaaS is a client for PMM d baa s
type PMMDBaaS struct {
	Components components.ClientService

	Kubernetes kubernetes.ClientService

	LogsAPI logs_api.ClientService

	PSMDBCluster psmdb_cluster.ClientService

	XtraDBCluster xtra_db_cluster.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PMMDBaaS) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Components.SetTransport(transport)
	c.Kubernetes.SetTransport(transport)
	c.LogsAPI.SetTransport(transport)
	c.PSMDBCluster.SetTransport(transport)
	c.XtraDBCluster.SetTransport(transport)
}

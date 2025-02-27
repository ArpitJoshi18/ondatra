/*
Package device is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/gdennis/go/pkg/mod/github.com/openconfig/ygot@v0.23.1/genutil/names.go
using the following YANG input files:
  - models-yang/models/isis/open-traffic-generator-isis.yang
  - models-yang/models/types/open-traffic-generator-types.yang
  - models-yang/models/flow/open-traffic-generator-flow.yang
  - models-yang/models/discovery/open-traffic-generator-discovery.yang
  - models-yang/models/interface/open-traffic-generator-port.yang
  - models-yang/models/bgp/open-traffic-generator-bgp.yang
  - models-yang/models/lag/open-traffic-generator-lag.yang
  - models-yang/models/lacp/open-traffic-generator-lacp.yang

Imported modules were sourced from:
  - models-yang/models/...
*/
package device

import (
	"github.com/openconfig/ondatra/telemetry/otg/bgp"
	"github.com/openconfig/ondatra/telemetry/otg/discovery"
	"github.com/openconfig/ondatra/telemetry/otg/flow"
	"github.com/openconfig/ondatra/telemetry/otg/isis"
	"github.com/openconfig/ondatra/telemetry/otg/lacp"
	"github.com/openconfig/ondatra/telemetry/otg/lag"
	"github.com/openconfig/ondatra/telemetry/otg/port"
	"github.com/openconfig/ygot/ygot"
)

// DevicePath represents the /device YANG schema element.
type DevicePath struct {
	*ygot.DeviceRootBase
}

// DeviceRoot returns a new path object from which YANG paths can be constructed.
func DeviceRoot(id string) *DevicePath {
	return &DevicePath{ygot.NewDeviceRootBase(id)}
}

// BgpPeerAny (list): Each BGP peer is identified by an arbitrary string
// identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "bgp-peers/bgp-peer"
// Path from root: "/bgp-peers/bgp-peer"
// Name (wildcarded): string
func (n *DevicePath) BgpPeerAny() *bgp.BgpPeerPathAny {
	return &bgp.BgpPeerPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"bgp-peers", "bgp-peer"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// BgpPeer (list): Each BGP peer is identified by an arbitrary string
// identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-bgp"
// Instantiating module: "open-traffic-generator-bgp"
// Path from parent: "bgp-peers/bgp-peer"
// Path from root: "/bgp-peers/bgp-peer"
// Name: string
func (n *DevicePath) BgpPeer(Name string) *bgp.BgpPeerPath {
	return &bgp.BgpPeerPath{
		NodePath: ygot.NewNodePath(
			[]string{"bgp-peers", "bgp-peer"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// FlowAny (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "flows/flow"
// Path from root: "/flows/flow"
// Name (wildcarded): string
func (n *DevicePath) FlowAny() *flow.FlowPathAny {
	return &flow.FlowPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Flow (list): A flow of packets between one or more internal and external sources
// and one or more internal and external destinations that the target
// is able to track and report statistics on. Each flow is identified by
// an arbitrary string identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "flows/flow"
// Path from root: "/flows/flow"
// Name: string
func (n *DevicePath) Flow(Name string) *flow.FlowPath {
	return &flow.FlowPath{
		NodePath: ygot.NewNodePath(
			[]string{"flows", "flow"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// InterfaceAny (list): An individual interface defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-discovery-interfaces"
// Instantiating module: "open-traffic-generator-discovery"
// Path from parent: "interfaces/interface"
// Path from root: "/interfaces/interface"
// Name (wildcarded): string
func (n *DevicePath) InterfaceAny() *discovery.InterfacePathAny {
	return &discovery.InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): An individual interface defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-discovery-interfaces"
// Instantiating module: "open-traffic-generator-discovery"
// Path from parent: "interfaces/interface"
// Path from root: "/interfaces/interface"
// Name: string
func (n *DevicePath) Interface(Name string) *discovery.InterfacePath {
	return &discovery.InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// IsisRouterAny (list): Each ISIS router is identified by an arbitrary string
// identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-isis"
// Instantiating module: "open-traffic-generator-isis"
// Path from parent: "isis-routers/isis-router"
// Path from root: "/isis-routers/isis-router"
// Name (wildcarded): string
func (n *DevicePath) IsisRouterAny() *isis.IsisRouterPathAny {
	return &isis.IsisRouterPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"isis-routers", "isis-router"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// IsisRouter (list): Each ISIS router is identified by an arbitrary string
// identifier.
// ----------------------------------------
// Defining module: "open-traffic-generator-isis"
// Instantiating module: "open-traffic-generator-isis"
// Path from parent: "isis-routers/isis-router"
// Path from root: "/isis-routers/isis-router"
// Name: string
func (n *DevicePath) IsisRouter(Name string) *isis.IsisRouterPath {
	return &isis.IsisRouterPath{
		NodePath: ygot.NewNodePath(
			[]string{"isis-routers", "isis-router"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Lacp (container): LACP telemetry collected by the ATE device.
// ----------------------------------------
// Defining module: "open-traffic-generator-lacp"
// Instantiating module: "open-traffic-generator-lacp"
// Path from parent: "lacp"
// Path from root: "/lacp"
func (n *DevicePath) Lacp() *lacp.LacpPath {
	return &lacp.LacpPath{
		NodePath: ygot.NewNodePath(
			[]string{"lacp"},
			map[string]interface{}{},
			n,
		),
	}
}

// LagAny (list): An individual LAG defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-lag"
// Instantiating module: "open-traffic-generator-lag"
// Path from parent: "lags/lag"
// Path from root: "/lags/lag"
// Name (wildcarded): string
func (n *DevicePath) LagAny() *lag.LagPathAny {
	return &lag.LagPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"lags", "lag"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Lag (list): An individual LAG defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-lag"
// Instantiating module: "open-traffic-generator-lag"
// Path from parent: "lags/lag"
// Path from root: "/lags/lag"
// Name: string
func (n *DevicePath) Lag(Name string) *lag.LagPath {
	return &lag.LagPath{
		NodePath: ygot.NewNodePath(
			[]string{"lags", "lag"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// PortAny (list): An individual port defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "ports/port"
// Path from root: "/ports/port"
// Name (wildcarded): string
func (n *DevicePath) PortAny() *port.PortPathAny {
	return &port.PortPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"ports", "port"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Port (list): An individual port defined by an OTG.
// ----------------------------------------
// Defining module: "open-traffic-generator-port"
// Instantiating module: "open-traffic-generator-port"
// Path from parent: "ports/port"
// Path from root: "/ports/port"
// Name: string
func (n *DevicePath) Port(Name string) *port.PortPath {
	return &port.PortPath{
		NodePath: ygot.NewNodePath(
			[]string{"ports", "port"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

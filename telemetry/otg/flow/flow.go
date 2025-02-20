/*
Package flow is a generated package which contains definitions
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
package flow

import (
	"github.com/openconfig/ygot/ygot"
)

// FlowPath represents the /open-traffic-generator-flow/flows/flow YANG schema element.
type FlowPath struct {
	*ygot.NodePath
}

// FlowPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow YANG schema element.
type FlowPathAny struct {
	*ygot.NodePath
}

// Flow_InFrameRatePath represents the /open-traffic-generator-flow/flows/flow/state/in-frame-rate YANG schema element.
type Flow_InFrameRatePath struct {
	*ygot.NodePath
}

// Flow_InFrameRatePathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/in-frame-rate YANG schema element.
type Flow_InFrameRatePathAny struct {
	*ygot.NodePath
}

// Flow_InRatePath represents the /open-traffic-generator-flow/flows/flow/state/in-rate YANG schema element.
type Flow_InRatePath struct {
	*ygot.NodePath
}

// Flow_InRatePathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/in-rate YANG schema element.
type Flow_InRatePathAny struct {
	*ygot.NodePath
}

// Flow_LossPctPath represents the /open-traffic-generator-flow/flows/flow/state/loss-pct YANG schema element.
type Flow_LossPctPath struct {
	*ygot.NodePath
}

// Flow_LossPctPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/loss-pct YANG schema element.
type Flow_LossPctPathAny struct {
	*ygot.NodePath
}

// Flow_NamePath represents the /open-traffic-generator-flow/flows/flow/state/name YANG schema element.
type Flow_NamePath struct {
	*ygot.NodePath
}

// Flow_NamePathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/name YANG schema element.
type Flow_NamePathAny struct {
	*ygot.NodePath
}

// Flow_OutFrameRatePath represents the /open-traffic-generator-flow/flows/flow/state/out-frame-rate YANG schema element.
type Flow_OutFrameRatePath struct {
	*ygot.NodePath
}

// Flow_OutFrameRatePathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/out-frame-rate YANG schema element.
type Flow_OutFrameRatePathAny struct {
	*ygot.NodePath
}

// Flow_OutRatePath represents the /open-traffic-generator-flow/flows/flow/state/out-rate YANG schema element.
type Flow_OutRatePath struct {
	*ygot.NodePath
}

// Flow_OutRatePathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/out-rate YANG schema element.
type Flow_OutRatePathAny struct {
	*ygot.NodePath
}

// Flow_TransmitPath represents the /open-traffic-generator-flow/flows/flow/state/transmit YANG schema element.
type Flow_TransmitPath struct {
	*ygot.NodePath
}

// Flow_TransmitPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/transmit YANG schema element.
type Flow_TransmitPathAny struct {
	*ygot.NodePath
}

// Counters (container): Counters that correspond to the individual flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/counters"
// Path from root: "/flows/flow/state/counters"
func (n *FlowPath) Counters() *Flow_CountersPath {
	return &Flow_CountersPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Counters (container): Counters that correspond to the individual flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/counters"
// Path from root: "/flows/flow/state/counters"
func (n *FlowPathAny) Counters() *Flow_CountersPathAny {
	return &Flow_CountersPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// InFrameRate (leaf): The rate, measured in frames per second, at which frames are being
// received for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/in-frame-rate"
// Path from root: "/flows/flow/state/in-frame-rate"
func (n *FlowPath) InFrameRate() *Flow_InFrameRatePath {
	return &Flow_InFrameRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-frame-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// InFrameRate (leaf): The rate, measured in frames per second, at which frames are being
// received for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/in-frame-rate"
// Path from root: "/flows/flow/state/in-frame-rate"
func (n *FlowPathAny) InFrameRate() *Flow_InFrameRatePathAny {
	return &Flow_InFrameRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-frame-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRate (leaf): The rate, measured in bits per second, at which the flow is being
// received.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/in-rate"
// Path from root: "/flows/flow/state/in-rate"
func (n *FlowPath) InRate() *Flow_InRatePath {
	return &Flow_InRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// InRate (leaf): The rate, measured in bits per second, at which the flow is being
// received.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/in-rate"
// Path from root: "/flows/flow/state/in-rate"
func (n *FlowPathAny) InRate() *Flow_InRatePathAny {
	return &Flow_InRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "in-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// LossPct (leaf): The percentage of transmitted packets that were not received by the
// destinations of the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/loss-pct"
// Path from root: "/flows/flow/state/loss-pct"
func (n *FlowPath) LossPct() *Flow_LossPctPath {
	return &Flow_LossPctPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "loss-pct"},
			map[string]interface{}{},
			n,
		),
	}
}

// LossPct (leaf): The percentage of transmitted packets that were not received by the
// destinations of the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/loss-pct"
// Path from root: "/flows/flow/state/loss-pct"
func (n *FlowPathAny) LossPct() *Flow_LossPctPathAny {
	return &Flow_LossPctPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "loss-pct"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name used for the flow tracked by the system. This
// name must be unique for the flows tracked and exported by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/name"
// Path from root: "/flows/flow/state/name"
func (n *FlowPath) Name() *Flow_NamePath {
	return &Flow_NamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Name (leaf): An arbitary name used for the flow tracked by the system. This
// name must be unique for the flows tracked and exported by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/name"
// Path from root: "/flows/flow/state/name"
func (n *FlowPathAny) Name() *Flow_NamePathAny {
	return &Flow_NamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "name"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutFrameRate (leaf): The rate, measured in frames per second, at which frames are being
// transmitted for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/out-frame-rate"
// Path from root: "/flows/flow/state/out-frame-rate"
func (n *FlowPath) OutFrameRate() *Flow_OutFrameRatePath {
	return &Flow_OutFrameRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-frame-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutFrameRate (leaf): The rate, measured in frames per second, at which frames are being
// transmitted for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/out-frame-rate"
// Path from root: "/flows/flow/state/out-frame-rate"
func (n *FlowPathAny) OutFrameRate() *Flow_OutFrameRatePathAny {
	return &Flow_OutFrameRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-frame-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRate (leaf): The rate, measured in bits per second, at which the flow is being
// transmitted.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/out-rate"
// Path from root: "/flows/flow/state/out-rate"
func (n *FlowPath) OutRate() *Flow_OutRatePath {
	return &Flow_OutRatePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutRate (leaf): The rate, measured in bits per second, at which the flow is being
// transmitted.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/out-rate"
// Path from root: "/flows/flow/state/out-rate"
func (n *FlowPathAny) OutRate() *Flow_OutRatePathAny {
	return &Flow_OutRatePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "out-rate"},
			map[string]interface{}{},
			n,
		),
	}
}

// Transmit (leaf): Whether or not the flow is transmitting
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/transmit"
// Path from root: "/flows/flow/state/transmit"
func (n *FlowPath) Transmit() *Flow_TransmitPath {
	return &Flow_TransmitPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "transmit"},
			map[string]interface{}{},
			n,
		),
	}
}

// Transmit (leaf): Whether or not the flow is transmitting
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "state/transmit"
// Path from root: "/flows/flow/state/transmit"
func (n *FlowPathAny) Transmit() *Flow_TransmitPathAny {
	return &Flow_TransmitPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "transmit"},
			map[string]interface{}{},
			n,
		),
	}
}

// Flow_CountersPath represents the /open-traffic-generator-flow/flows/flow/state/counters YANG schema element.
type Flow_CountersPath struct {
	*ygot.NodePath
}

// Flow_CountersPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/counters YANG schema element.
type Flow_CountersPathAny struct {
	*ygot.NodePath
}

// Flow_Counters_InOctetsPath represents the /open-traffic-generator-flow/flows/flow/state/counters/in-octets YANG schema element.
type Flow_Counters_InOctetsPath struct {
	*ygot.NodePath
}

// Flow_Counters_InOctetsPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/counters/in-octets YANG schema element.
type Flow_Counters_InOctetsPathAny struct {
	*ygot.NodePath
}

// Flow_Counters_InPktsPath represents the /open-traffic-generator-flow/flows/flow/state/counters/in-pkts YANG schema element.
type Flow_Counters_InPktsPath struct {
	*ygot.NodePath
}

// Flow_Counters_InPktsPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/counters/in-pkts YANG schema element.
type Flow_Counters_InPktsPathAny struct {
	*ygot.NodePath
}

// Flow_Counters_OutOctetsPath represents the /open-traffic-generator-flow/flows/flow/state/counters/out-octets YANG schema element.
type Flow_Counters_OutOctetsPath struct {
	*ygot.NodePath
}

// Flow_Counters_OutOctetsPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/counters/out-octets YANG schema element.
type Flow_Counters_OutOctetsPathAny struct {
	*ygot.NodePath
}

// Flow_Counters_OutPktsPath represents the /open-traffic-generator-flow/flows/flow/state/counters/out-pkts YANG schema element.
type Flow_Counters_OutPktsPath struct {
	*ygot.NodePath
}

// Flow_Counters_OutPktsPathAny represents the wildcard version of the /open-traffic-generator-flow/flows/flow/state/counters/out-pkts YANG schema element.
type Flow_Counters_OutPktsPathAny struct {
	*ygot.NodePath
}

// InOctets (leaf): The total number of bytes received by the target for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "in-octets"
// Path from root: "/flows/flow/state/counters/in-octets"
func (n *Flow_CountersPath) InOctets() *Flow_Counters_InOctetsPath {
	return &Flow_Counters_InOctetsPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// InOctets (leaf): The total number of bytes received by the target for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "in-octets"
// Path from root: "/flows/flow/state/counters/in-octets"
func (n *Flow_CountersPathAny) InOctets() *Flow_Counters_InOctetsPathAny {
	return &Flow_Counters_InOctetsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// InPkts (leaf): The total number of packets received by the target for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "in-pkts"
// Path from root: "/flows/flow/state/counters/in-pkts"
func (n *Flow_CountersPath) InPkts() *Flow_Counters_InPktsPath {
	return &Flow_Counters_InPktsPath{
		NodePath: ygot.NewNodePath(
			[]string{"in-pkts"},
			map[string]interface{}{},
			n,
		),
	}
}

// InPkts (leaf): The total number of packets received by the target for the flow.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "in-pkts"
// Path from root: "/flows/flow/state/counters/in-pkts"
func (n *Flow_CountersPathAny) InPkts() *Flow_Counters_InPktsPathAny {
	return &Flow_Counters_InPktsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"in-pkts"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOctets (leaf): The total number of bytes sent by the target for the flow. These
// packets may be generated or forwarded by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "out-octets"
// Path from root: "/flows/flow/state/counters/out-octets"
func (n *Flow_CountersPath) OutOctets() *Flow_Counters_OutOctetsPath {
	return &Flow_Counters_OutOctetsPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutOctets (leaf): The total number of bytes sent by the target for the flow. These
// packets may be generated or forwarded by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "out-octets"
// Path from root: "/flows/flow/state/counters/out-octets"
func (n *Flow_CountersPathAny) OutOctets() *Flow_Counters_OutOctetsPathAny {
	return &Flow_Counters_OutOctetsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-octets"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutPkts (leaf): The total number of packets sent by the target for the flow. These
// packets may be generated or forwarded by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "out-pkts"
// Path from root: "/flows/flow/state/counters/out-pkts"
func (n *Flow_CountersPath) OutPkts() *Flow_Counters_OutPktsPath {
	return &Flow_Counters_OutPktsPath{
		NodePath: ygot.NewNodePath(
			[]string{"out-pkts"},
			map[string]interface{}{},
			n,
		),
	}
}

// OutPkts (leaf): The total number of packets sent by the target for the flow. These
// packets may be generated or forwarded by the target.
// ----------------------------------------
// Defining module: "open-traffic-generator-flow"
// Instantiating module: "open-traffic-generator-flow"
// Path from parent: "out-pkts"
// Path from root: "/flows/flow/state/counters/out-pkts"
func (n *Flow_CountersPathAny) OutPkts() *Flow_Counters_OutPktsPathAny {
	return &Flow_Counters_OutPktsPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"out-pkts"},
			map[string]interface{}{},
			n,
		),
	}
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AudioRoutingGroup undocumented
type AudioRoutingGroup struct {
	// Entity is the base model of AudioRoutingGroup
	Entity
	// RoutingMode undocumented
	RoutingMode *RoutingMode `json:"routingMode,omitempty"`
	// Sources undocumented
	Sources []string `json:"sources,omitempty"`
	// Receivers undocumented
	Receivers []string `json:"receivers,omitempty"`
}
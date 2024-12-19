// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package vswitchd

import "github.com/ovn-org/libovsdb/model"

const NetFlowTable = "NetFlow"

// NetFlow defines an object in NetFlow table
type NetFlow struct {
	UUID             string            `ovsdb:"_uuid"`
	ActiveTimeout    int               `ovsdb:"active_timeout"`
	AddIDToInterface bool              `ovsdb:"add_id_to_interface"`
	EngineID         *int              `ovsdb:"engine_id"`
	EngineType       *int              `ovsdb:"engine_type"`
	ExternalIDs      map[string]string `ovsdb:"external_ids"`
	Targets          []string          `ovsdb:"targets"`
}

func (a *NetFlow) GetUUID() string {
	return a.UUID
}

func (a *NetFlow) GetActiveTimeout() int {
	return a.ActiveTimeout
}

func (a *NetFlow) GetAddIDToInterface() bool {
	return a.AddIDToInterface
}

func (a *NetFlow) GetEngineID() *int {
	return a.EngineID
}

func copyNetFlowEngineID(a *int) *int {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalNetFlowEngineID(a, b *int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *NetFlow) GetEngineType() *int {
	return a.EngineType
}

func copyNetFlowEngineType(a *int) *int {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalNetFlowEngineType(a, b *int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *NetFlow) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyNetFlowExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalNetFlowExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *NetFlow) GetTargets() []string {
	return a.Targets
}

func copyNetFlowTargets(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalNetFlowTargets(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *NetFlow) DeepCopyInto(b *NetFlow) {
	*b = *a
	b.EngineID = copyNetFlowEngineID(a.EngineID)
	b.EngineType = copyNetFlowEngineType(a.EngineType)
	b.ExternalIDs = copyNetFlowExternalIDs(a.ExternalIDs)
	b.Targets = copyNetFlowTargets(a.Targets)
}

func (a *NetFlow) DeepCopy() *NetFlow {
	b := new(NetFlow)
	a.DeepCopyInto(b)
	return b
}

func (a *NetFlow) CloneModelInto(b model.Model) {
	c := b.(*NetFlow)
	a.DeepCopyInto(c)
}

func (a *NetFlow) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *NetFlow) Equals(b *NetFlow) bool {
	return a.UUID == b.UUID &&
		a.ActiveTimeout == b.ActiveTimeout &&
		a.AddIDToInterface == b.AddIDToInterface &&
		equalNetFlowEngineID(a.EngineID, b.EngineID) &&
		equalNetFlowEngineType(a.EngineType, b.EngineType) &&
		equalNetFlowExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		equalNetFlowTargets(a.Targets, b.Targets)
}

func (a *NetFlow) EqualsModel(b model.Model) bool {
	c := b.(*NetFlow)
	return a.Equals(c)
}

var _ model.CloneableModel = &NetFlow{}
var _ model.ComparableModel = &NetFlow{}
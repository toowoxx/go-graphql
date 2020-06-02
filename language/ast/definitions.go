package ast

import (
	"github.com/graphql-go/graphql/language/kinds"
)

type Definition interface {
	GetOperation() string
	GetVariableDefinitions() []*VariableDefinition
	GetSelectionSet() *SelectionSet
	GetKind() string
	GetLoc() *Location
}

// Ensure that all definition types implements Definition interface
var _ Definition = (*OperationDefinition)(nil)
var _ Definition = (*FragmentDefinition)(nil)
var _ Definition = (TypeSystemDefinition)(nil) // experimental non-spec addition.

// Note: subscription is an experimental non-spec addition.
const (
	OperationTypeQuery        = "query"
	OperationTypeMutation     = "mutation"
	OperationTypeSubscription = "subscription"
)

// OperationDefinition implements Node, Definition
type OperationDefinition struct {
	Kind                string
	Loc                 *Location
	Operation           string
	Name                *Name
	VariableDefinitions []*VariableDefinition
	Directives          []*Directive
	SelectionSet        *SelectionSet
}

func NewOperationDefinition(op *OperationDefinition) *OperationDefinition {
	if op == nil {
		op = &OperationDefinition{}
	}
	op.Kind = kinds.OperationDefinition
	return op
}

func (op *OperationDefinition) GetKind() string {
	return op.Kind
}

func (op *OperationDefinition) GetLoc() *Location {
	return op.Loc
}

func (op *OperationDefinition) GetOperation() string {
	return op.Operation
}

func (op *OperationDefinition) GetName() *Name {
	return op.Name
}

func (op *OperationDefinition) GetVariableDefinitions() []*VariableDefinition {
	return op.VariableDefinitions
}

func (op *OperationDefinition) GetDirectives() []*Directive {
	return op.Directives
}

func (op *OperationDefinition) GetSelectionSet() *SelectionSet {
	return op.SelectionSet
}

// FragmentDefinition implements Node, Definition
type FragmentDefinition struct {
	Kind                string
	Loc                 *Location
	Operation           string
	Name                *Name
	VariableDefinitions []*VariableDefinition
	TypeCondition       *Named
	Directives          []*Directive
	SelectionSet        *SelectionSet
}

func NewFragmentDefinition(fd *FragmentDefinition) *FragmentDefinition {
	if fd == nil {
		fd = &FragmentDefinition{}
	}
	return &FragmentDefinition{
		Kind:                kinds.FragmentDefinition,
		Loc:                 fd.Loc,
		Operation:           fd.Operation,
		Name:                fd.Name,
		VariableDefinitions: fd.VariableDefinitions,
		TypeCondition:       fd.TypeCondition,
		Directives:          fd.Directives,
		SelectionSet:        fd.SelectionSet,
	}
}

func (fd *FragmentDefinition) GetKind() string {
	return fd.Kind
}

func (fd *FragmentDefinition) GetLoc() *Location {
	return fd.Loc
}

func (fd *FragmentDefinition) GetOperation() string {
	return fd.Operation
}

func (fd *FragmentDefinition) GetName() *Name {
	return fd.Name
}

func (fd *FragmentDefinition) GetVariableDefinitions() []*VariableDefinition {
	return fd.VariableDefinitions
}

func (fd *FragmentDefinition) GetSelectionSet() *SelectionSet {
	return fd.SelectionSet
}

// VariableDefinition implements Node
type VariableDefinition struct {
	Kind         string
	Loc          *Location
	Variable     *Variable
	Type         Type
	DefaultValue Value
}

func NewVariableDefinition(vd *VariableDefinition) *VariableDefinition {
	if vd == nil {
		vd = &VariableDefinition{}
	}
	vd.Kind = kinds.VariableDefinition
	return vd
}

func (vd *VariableDefinition) GetKind() string {
	return vd.Kind
}

func (vd *VariableDefinition) GetLoc() *Location {
	return vd.Loc
}

// ScalarExtensionDefinition implements Node, Definition
type ScalarExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *ScalarDefinition
}

func NewScalarExtensionDefinition(def *ScalarExtensionDefinition) *ScalarExtensionDefinition {
	if def == nil {
		def = &ScalarExtensionDefinition{}
	}
	return &ScalarExtensionDefinition{
		Kind:       kinds.ScalarExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *ScalarExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *ScalarExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *ScalarExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *ScalarExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *ScalarExtensionDefinition) GetOperation() string {
	return ""
}

// ObjectExtensionDefinition implements Node, Definition
type ObjectExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *ObjectDefinition
}

func NewObjectExtensionDefinition(def *ObjectExtensionDefinition) *ObjectExtensionDefinition {
	if def == nil {
		def = &ObjectExtensionDefinition{}
	}
	return &ObjectExtensionDefinition{
		Kind:       kinds.ObjectExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *ObjectExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *ObjectExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *ObjectExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *ObjectExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *ObjectExtensionDefinition) GetOperation() string {
	return ""
}

// InterfaceExtensionDefinition implements Node, Definition
type InterfaceExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *InterfaceDefinition
}

func NewInterfaceExtensionDefinition(def *InterfaceExtensionDefinition) *InterfaceExtensionDefinition {
	if def == nil {
		def = &InterfaceExtensionDefinition{}
	}
	return &InterfaceExtensionDefinition{
		Kind:       kinds.InterfaceExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *InterfaceExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *InterfaceExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *InterfaceExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *InterfaceExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *InterfaceExtensionDefinition) GetOperation() string {
	return ""
}

// UnionExtensionDefinition implements Node, Definition
type UnionExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *UnionDefinition
}

func NewUnionExtensionDefinition(def *UnionExtensionDefinition) *UnionExtensionDefinition {
	if def == nil {
		def = &UnionExtensionDefinition{}
	}
	return &UnionExtensionDefinition{
		Kind:       kinds.UnionExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *UnionExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *UnionExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *UnionExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *UnionExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *UnionExtensionDefinition) GetOperation() string {
	return ""
}

// EnumExtensionDefinition implements Node, Definition
type EnumExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *EnumDefinition
}

func NewEnumExtensionDefinition(def *EnumExtensionDefinition) *EnumExtensionDefinition {
	if def == nil {
		def = &EnumExtensionDefinition{}
	}
	return &EnumExtensionDefinition{
		Kind:       kinds.EnumExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *EnumExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *EnumExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *EnumExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *EnumExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *EnumExtensionDefinition) GetOperation() string {
	return ""
}

// InputObjectExtensionDefinition implements Node, Definition
type InputObjectExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *InputObjectDefinition
}

func NewInputObjectExtensionDefinition(def *InputObjectExtensionDefinition) *InputObjectExtensionDefinition {
	if def == nil {
		def = &InputObjectExtensionDefinition{}
	}
	return &InputObjectExtensionDefinition{
		Kind:       kinds.InputObjectExtensionDefinition,
		Loc:        def.Loc,
		Definition: def.Definition,
	}
}

func (def *InputObjectExtensionDefinition) GetKind() string {
	return def.Kind
}

func (def *InputObjectExtensionDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *InputObjectExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *InputObjectExtensionDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *InputObjectExtensionDefinition) GetOperation() string {
	return ""
}

// DirectiveDefinition implements Node, Definition
type DirectiveDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Arguments   []*InputValueDefinition
	Locations   []*Name
}

func NewDirectiveDefinition(def *DirectiveDefinition) *DirectiveDefinition {
	if def == nil {
		def = &DirectiveDefinition{}
	}
	return &DirectiveDefinition{
		Kind:        kinds.DirectiveDefinition,
		Loc:         def.Loc,
		Name:        def.Name,
		Description: def.Description,
		Arguments:   def.Arguments,
		Locations:   def.Locations,
	}
}

func (def *DirectiveDefinition) GetKind() string {
	return def.Kind
}

func (def *DirectiveDefinition) GetLoc() *Location {
	return def.Loc
}

func (def *DirectiveDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *DirectiveDefinition) GetSelectionSet() *SelectionSet {
	return &SelectionSet{}
}

func (def *DirectiveDefinition) GetOperation() string {
	return ""
}

func (def *DirectiveDefinition) GetDescription() *StringValue {
	return def.Description
}

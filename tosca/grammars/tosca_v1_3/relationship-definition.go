package tosca_v1_3

import (
	"github.com/tliron/puccini/tosca"
)

//
// RelationshipDefinition
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.7.3.2.3
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.7.3.2.3
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.6.3.2.3
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.6.2.2.3
//

type RelationshipDefinition struct {
	*Entity `name:"relationship definition"`

	RelationshipTypeName *string              `read:"type"` // required only if cannot be inherited
	InterfaceDefinitions InterfaceDefinitions `read:"interfaces,InterfaceDefinition" inherit:"interfaces,RelationshipType"`

	RelationshipType *RelationshipType `lookup:"type,RelationshipTypeName" json:"-" yaml:"-"`

	typeMissingProblemReported bool
}

func NewRelationshipDefinition(context *tosca.Context) *RelationshipDefinition {
	return &RelationshipDefinition{
		Entity:               NewEntity(context),
		InterfaceDefinitions: make(InterfaceDefinitions),
	}
}

// tosca.Reader signature
func ReadRelationshipDefinition(context *tosca.Context) tosca.EntityPtr {
	self := NewRelationshipDefinition(context)

	if context.Is("!!map") {
		// Long notation
		context.ValidateUnsupportedFields(context.ReadFields(self))
	} else if context.ValidateType("!!map", "!!str") {
		// Short notation
		self.RelationshipTypeName = context.FieldChild("type", context.Data).ReadString()
	}

	return self
}

func (self *RelationshipDefinition) NewDefaultAssignment(context *tosca.Context) *RelationshipAssignment {
	assignment := NewRelationshipAssignment(context)
	assignment.RelationshipTemplateNameOrTypeName = self.RelationshipTypeName
	assignment.RelationshipType = self.RelationshipType
	return assignment
}

func (self *RelationshipDefinition) Inherit(parentDefinition *RelationshipDefinition) {
	log.Debug("{inherit} relationship definition")

	if (self.RelationshipTypeName == nil) && (parentDefinition.RelationshipTypeName != nil) {
		self.RelationshipTypeName = parentDefinition.RelationshipTypeName
	}
	if (self.RelationshipType == nil) && (parentDefinition.RelationshipType != nil) {
		self.RelationshipType = parentDefinition.RelationshipType
	}

	// Validate type compatibility
	if (self.RelationshipType != nil) && (parentDefinition.RelationshipType != nil) && !self.Context.Hierarchy.IsCompatible(parentDefinition.RelationshipType, self.RelationshipType) {
		self.Context.ReportIncompatibleType(self.RelationshipType, parentDefinition.RelationshipType)
		return
	}

	self.InterfaceDefinitions.Inherit(parentDefinition.InterfaceDefinitions)
}

// parser.Renderable interface
func (self *RelationshipDefinition) Render() {
	log.Debug("{render} relationship definition")

	if self.RelationshipTypeName == nil {
		// Avoid reporting more than once
		if !self.typeMissingProblemReported {
			self.Context.FieldChild("type", nil).ReportFieldMissing()
			self.typeMissingProblemReported = true
		}
	}
}

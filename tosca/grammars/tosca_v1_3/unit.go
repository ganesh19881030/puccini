package tosca_v1_3

import (
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
)

//
// Unit
//
// See ServiceTemplate
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.10
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.10
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.9
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.9
//

type Unit struct {
	*Entity `name:"unit"`

	ToscaDefinitionsVersion *string           `read:"tosca_definitions_version" require:"tosca_definitions_version"`
	Namespace               *string           `read:"namespace"` // introduced in TOSCA 1.2
	Metadata                Metadata          `read:"metadata,!Metadata"`
	Repositories            Repositories      `read:"repositories,Repository"`
	Imports                 Imports           `read:"imports,[]Import"`
	ArtifactTypes           ArtifactTypes     `read:"artifact_types,ArtifactType" hierarchy:""`
	CapabilityTypes         CapabilityTypes   `read:"capability_types,CapabilityType" hierarchy:""`
	DataTypes               DataTypes         `read:"data_types,DataType" hierarchy:""`
	GroupTypes              GroupTypes        `read:"group_types,GroupType" hierarchy:""`
	InterfaceTypes          InterfaceTypes    `read:"interface_types,InterfaceType" hierarchy:""`
	NodeTypes               NodeTypes         `read:"node_types,NodeType" hierarchy:""`
	PolicyTypes             PolicyTypes       `read:"policy_types,PolicyType" hierarchy:""`
	RelationshipTypes       RelationshipTypes `read:"relationship_types,RelationshipType" hierarchy:""`
}

func NewUnit(context *tosca.Context) *Unit {
	return &Unit{Entity: NewEntity(context)}
}

// tosca.Reader signature
func ReadUnit(context *tosca.Context) tosca.EntityPtr {
	self := NewUnit(context)
	context.ScriptletNamespace.Merge(DefaultScriptletNamespace)
	context.ValidateUnsupportedFields(append(context.ReadFields(self), "dsl_definitions"))
	if self.Namespace != nil {
		context.CanonicalNamespace = self.Namespace
	}
	return self
}

// parser.Importer interface
func (self *Unit) GetImportSpecs() []*tosca.ImportSpec {
	// TODO: importing should also import repositories

	var importSpecs = make([]*tosca.ImportSpec, 0, len(self.Imports))
	for _, import_ := range self.Imports {
		if importSpec, ok := import_.NewImportSpec(self); ok {
			importSpecs = append(importSpecs, importSpec)
		}
	}
	return importSpecs
}

func (self *Unit) Normalize(normalServiceTemplate *normal.ServiceTemplate) {
	log.Debug("{normalize} unit")

	if self.Metadata != nil {
		for k, v := range self.Metadata {
			normalServiceTemplate.Metadata[k] = v
		}
	}
}

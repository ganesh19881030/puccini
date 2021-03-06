package normal

//
// Relationship
//

type Relationship struct {
	Requirement *Requirement `json:"-" yaml:"-"`
	Name        string       `json:"-" yaml:"-"`

	Description string         `json:"description" yaml:"description"`
	Types       Types          `json:"types" yaml:"types"`
	Properties  Constrainables `json:"properties" yaml:"properties"`
	Attributes  Constrainables `json:"attributes" yaml:"attributes"`
	Interfaces  Interfaces     `json:"interfaces" yaml:"interfaces"`
}

func (self *Requirement) NewRelationship() *Relationship {
	relationship := &Relationship{
		Requirement: self,
		Types:       make(Types),
		Properties:  make(Constrainables),
		Attributes:  make(Constrainables),
		Interfaces:  make(Interfaces),
	}
	self.Relationship = relationship
	return relationship
}

//
// Relationships
//

type Relationships []*Relationship

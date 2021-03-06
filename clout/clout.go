package clout

import (
	"encoding/json"
	"errors"
	"fmt"

	"strings"

	"github.com/tliron/puccini/ard"
	"github.com/tliron/puccini/common/format"
)

const Version = "1.0"

//
// Clout
//

type Clout struct {
	Version    string        `yaml:"version"`
	Metadata   ard.StringMap `yaml:"metadata"`
	Properties ard.StringMap `yaml:"properties"`
	Vertexes   Vertexes      `yaml:"vertexes"`
}

func NewClout() *Clout {
	return &Clout{
		Version:    Version,
		Metadata:   make(ard.StringMap),
		Properties: make(ard.StringMap),
		Vertexes:   make(Vertexes),
	}
}

type MarshalableCloutStringMaps Clout

func (self *Clout) MarshalableStringMaps() interface{} {
	return &MarshalableCloutStringMaps{
		Version:    self.Version,
		Metadata:   ard.EnsureStringMaps(self.Metadata),
		Properties: ard.EnsureStringMaps(self.Properties),
		Vertexes:   self.Vertexes,
	}
}

// json.Marshaler interface
func (self *Clout) MarshalJSON() ([]byte, error) {
	// JavaScript requires keys to be strings, so we would lose complex keys
	return json.Marshal(self.MarshalableStringMaps())
}

func (self *Clout) Resolve() error {
	if self.Version == "" {
		return errors.New("no Clout \"Version\"")
	}
	if self.Version != Version {
		return fmt.Errorf("unsupported Clout version: \"%s\"", self.Version)
	}

	for key, v := range self.Vertexes {
		v.Clout = self
		v.ID = key

		for _, e := range v.EdgesOut {
			var ok bool
			if e.Target, ok = self.Vertexes[e.TargetID]; !ok {
				return fmt.Errorf("could not resolve Clout, bad TargetID: \"%s\"", e.TargetID)
			}

			e.Source = v
			e.Target.EdgesIn = append(e.Target.EdgesIn, e)
		}
	}
	return nil
}

func (self *Clout) Normalize() (*Clout, error) {
	// TODO: there must be a more efficient way to do this
	if s, err := format.EncodeYAML(self, " ", false); err == nil {
		return ReadYAML(strings.NewReader(s))
	} else {
		return nil, err
	}
}

func (self *Clout) ARD() (ard.Map, error) {
	if s, err := format.EncodeYAML(self, " ", false); err == nil {
		map_, _, err := ard.ReadYAML(strings.NewReader(s), false)
		return map_, err
	} else {
		return nil, err
	}
}

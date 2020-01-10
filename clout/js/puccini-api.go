package js

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/fatih/color"
	"github.com/tliron/puccini/common"
	"github.com/tliron/puccini/common/format"
)

//
// PucciniAPI
//

type PucciniAPI struct {
	Log    *Log
	Stdout *os.File
	Stderr *os.File
	Stdin  *os.File
	Format string
	Pretty bool

	context *Context
}

func (self *Context) NewPucciniAPI() *PucciniAPI {
	format := self.Format
	if format == "" {
		format = "yaml"
	}
	return &PucciniAPI{
		Log:     self.Log,
		Stdout:  self.Stdout,
		Stderr:  self.Stderr,
		Stdin:   self.Stdin,
		Format:  format,
		Pretty:  self.Pretty,
		context: self,
	}
}

func (entry *PucciniAPI) Sprintf(f string, a ...interface{}) string {
	return fmt.Sprintf(f, a...)
}

func (self *PucciniAPI) Timestamp() string {
	return common.Timestamp()
}

func (self *PucciniAPI) NewXMLDocument() *etree.Document {
	return etree.NewDocument()
}

func (self *PucciniAPI) Write(data interface{}, path string, dontOverwrite bool) {
	output := self.context.Output
	if path != "" {
		output = filepath.Join(output, path)
		var err error
		output, err = filepath.Abs(output)
		self.context.FailOnError(err)
	}

	if output == "" {
		if self.context.Quiet {
			return
		}
	} else {
		_, err := os.Stat(output)
		var message string
		var skip bool
		if (err == nil) || os.IsExist(err) {
			if dontOverwrite {
				message = color.RedString("skippping:  ")
				skip = true
			} else {
				message = color.YellowString("overwriting:")
			}
		} else {
			message = color.GreenString("writing:    ")
		}
		if !self.context.Quiet {
			fmt.Fprintln(self.Stdout, fmt.Sprintf("%s %s", message, output))
		}
		if skip {
			return
		}
	}

	err := format.WriteOrPrint(data, self.Format, self.Pretty, output)
	self.context.FailOnError(err)
}

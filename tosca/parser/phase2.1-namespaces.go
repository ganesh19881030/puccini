package parser

import (
	"fmt"

	"github.com/tliron/puccini/common/terminal"
	"github.com/tliron/puccini/tosca"
)

func (self *Context) AddNamespaces() {
	self.Root.MergeNamespaces(self.NamespacesWork)
}

func (self *Unit) MergeNamespaces(work *ContextualWork) {
	context := self.GetContext()

	if promise, ok := work.Start(context); ok {
		defer promise.Release()

		for _, import_ := range self.Imports {
			import_.MergeNamespaces(work)
			context.Namespace.Merge(import_.GetContext().Namespace, import_.NameTransformer)
			context.ScriptletNamespace.Merge(import_.GetContext().ScriptletNamespace)
		}

		log.Debugf("{namespaces} create: %s", context.URL.String())
		namespace := tosca.NewNamespaceFor(self.EntityPtr)
		context.Namespace.Merge(namespace, nil)
	}
}

// Print

func (self *Context) PrintNamespaces(indent int) {
	childIndent := indent + 1
	for _, import_ := range self.Units {
		context := import_.GetContext()
		if !context.Namespace.Empty() {
			terminal.PrintIndent(indent)
			fmt.Fprintf(terminal.Stdout, "%s\n", terminal.ColorValue(context.URL.String()))
			context.Namespace.Print(childIndent)
		}
	}
}

package commands

import (
	"github.com/op/go-logging"
	"github.com/tliron/puccini/clout"
	urlpkg "github.com/tliron/puccini/url"
)

const toolName = "puccini-js"

var log = logging.MustGetLogger(toolName)

var output string

func ReadClout(path string) (*clout.Clout, error) {
	var url urlpkg.URL

	var err error
	if path != "" {
		if url, err = urlpkg.NewValidURL(path, nil); err != nil {
			return nil, err
		}
	} else {
		if url, err = urlpkg.ReadToInternalURLFromStdin("yaml"); err != nil {
			return nil, err
		}
	}
	defer url.Release()

	if reader, err := url.Open(); err == nil {
		defer reader.Close()
		return clout.Read(reader, url.Format())
	} else {
		return nil, err
	}
}

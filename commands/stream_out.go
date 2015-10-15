package commands

import (
	"io"
	"os"

	"github.com/cloudfoundry-incubator/garden"
)

type StreamOut struct {
	Source string `short:"s" long:"source" description:"source path in the container" required:"true"`
}

func (command *StreamOut) Execute(maybeHandle []string) error {
	container, err := globalClient().Lookup(handle(maybeHandle))
	failIf(err)

	streamOutSpec := garden.StreamOutSpec{Path: command.Source}
	output, err := container.StreamOut(streamOutSpec)
	failIf(err)

	io.Copy(os.Stdout, output)

	return nil
}
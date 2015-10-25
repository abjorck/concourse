package commands

import (
	"fmt"
	"log"
	"runtime"

	"github.com/inconshreveable/go-update"

	"github.com/concourse/fly/atcclient"
	"github.com/concourse/fly/rc"
)

type SyncCommand struct{}

var syncCommand SyncCommand

func init() {
	sync, err := Parser.AddCommand(
		"sync",
		"Download and replace the current fly from the target",
		"",
		&syncCommand,
	)
	if err != nil {
		panic(err)
	}

	sync.Aliases = []string{"s"}
}

func (command *SyncCommand) Execute(args []string) error {
	client, err := rc.TargetClient(globalOptions.Target)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	handler := atcclient.NewAtcHandler(client)
	body, err := handler.GetCLIReader(runtime.GOARCH, runtime.GOOS)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("downloading fly from %s... ", client.URL())

	err = update.Apply(body, update.Options{})
	if err != nil {
		failf("update failed: %s", err)
	}

	fmt.Println("update successful!")
	return nil
}

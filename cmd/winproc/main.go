package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var cli struct {
		List  ListCmd  `kong:"cmd,help='Provides a list view of the windows process list.'"`
		Tree  TreeCmd  `kong:"cmd,help='Provides a tree view of the windows process list.'"`
		Watch WatchCmd `kong:"cmd,help='Watches the windows process list.'"`
	}

	parser := kong.Must(&cli,
		kong.Description("Shows information about running windows processes."),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.UsageOnError())

	app, parseErr := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(parseErr)

	appErr := app.Run()
	app.FatalIfErrorf(appErr)
}

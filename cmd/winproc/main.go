package main

import (
	"os"
	"syscall"

	"github.com/gentlemanautomaton/signaler"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app                     = kingpin.New("winproc", "Shows information about running windows processes.")
		listCommand             = app.Command("list", "Provides a list view of the windows process list.")
		listIncludePIDs         = listCommand.Flag("pid", "include processes with a particular ID").Ints()
		listIncludeNames        = listCommand.Flag("name", "include processes with a particular name").Strings()
		listIncludeAncestors    = listCommand.Flag("ancestors", "include ancestors of matching processes").Short('a').Bool()
		listIncludeDescendants  = listCommand.Flag("descendants", "include descendants of matching processes").Short('d').Bool()
		treeCommand             = app.Command("tree", "Provides a tree view of the windows process list.")
		treeIncludePIDs         = treeCommand.Flag("pid", "include processes with a particular ID").Ints()
		treeIncludeNames        = treeCommand.Flag("name", "include processes with a particular name").Strings()
		treeIncludeAncestors    = treeCommand.Flag("ancestors", "include ancestors of matching processes").Short('a').Bool()
		treeIncludeDescendants  = treeCommand.Flag("descendants", "include descendants of matching processes").Short('d').Bool()
		watchCommand            = app.Command("watch", "Watches the windows process list.")
		watchIncludePIDs        = watchCommand.Flag("pid", "include processes with a particular ID").Ints()
		watchIncludeNames       = watchCommand.Flag("name", "include processes with a particular name").Strings()
		watchIncludeAncestors   = watchCommand.Flag("ancestors", "include ancestors of matching processes").Short('a').Bool()
		watchIncludeDescendants = watchCommand.Flag("descendants", "include descendants of matching processes").Short('d').Bool()
		watchInterval           = watchCommand.Flag("interval", "interval between updates").Short('i').Default("1s").Duration()
	)

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	shutdown := signaler.New().Capture(os.Interrupt, syscall.SIGTERM)
	ctx := shutdown.Context()
	defer shutdown.Trigger()

	if ctx.Err() != nil {
		return
	}

	switch command {
	case listCommand.FullCommand():
		opts := makeOptions(*listIncludePIDs, *listIncludeNames, *listIncludeAncestors, *listIncludeDescendants)
		list(ctx, opts...)
	case treeCommand.FullCommand():
		opts := makeOptions(*treeIncludePIDs, *treeIncludeNames, *treeIncludeAncestors, *treeIncludeDescendants)
		tree(ctx, opts...)
	case watchCommand.FullCommand():
		opts := makeOptions(*watchIncludePIDs, *watchIncludeNames, *watchIncludeAncestors, *watchIncludeDescendants)
		watch(ctx, *watchInterval, opts...)
	}
}

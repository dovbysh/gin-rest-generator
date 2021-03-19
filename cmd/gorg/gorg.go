package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dovbysh/gin-rest-generator/gorg"
	"github.com/dovbysh/gin-rest-generator/loader"
)

func init() {
	ldr := loader.New(nil)

	gorg.RegisterCommand("gen", gorg.NewGenerateCommand(ldr))
	gorg.RegisterCommand("template", gorg.NewTemplateCommand(ldr))
}

func main() {
	if len(os.Args) < 2 {
		if err := gorg.Usage(os.Stderr); err != nil {
			die(1, err.Error())
		}
		os.Exit(2)
	}

	flag.CommandLine.Usage = func() {
		die(2, "Run 'gorg help' for usage.")
	}

	flag.Parse()
	args := flag.Args()

	if args[0] == "help" {
		if err := help(args[1:], os.Stdout); err != nil {
			die(2, err.Error())
		}
		return
	}

	command := gorg.GetCommand(args[0])
	if command == nil {
		die(2, "gorg: unknown subcommand %q\nRun 'gorg help' for usage.", args[0])
	}

	if err := command.Run(args[1:], os.Stdout); err != nil {
		if _, ok := err.(gorg.CommandLineError); ok {
			die(2, "%s\nRun 'gorg help %s' for usage.\n", err.Error(), args[0])
		}
		die(1, err.Error())
	}
}

func die(exitCode int, format string, args ...interface{}) {
	if _, err := fmt.Fprintf(os.Stderr, format+"\n", args...); err != nil {
		os.Exit(1)
	}
	os.Exit(exitCode)
}

func help(args []string, w io.Writer) error {
	if len(args) > 1 {
		return errors.New("usage: gorg help command\n\nToo many arguments given")
	}

	if len(args) == 0 {
		return gorg.Usage(w)
	}

	command := gorg.GetCommand(args[0])
	if command == nil {
		return fmt.Errorf(fmt.Sprintf("gounit: unknown subcommand %q\nRun 'gounit help' for usage", args[0]))
	}

	if _, err := fmt.Fprintf(w, "Usage: gorg %s %s\n", args[0], command.UsageLine()); err != nil {
		return err
	}

	if fs := command.FlagSet(); fs != nil {
		fs.SetOutput(w)
		fs.PrintDefaults()
	}

	return command.HelpMessage(w)
}

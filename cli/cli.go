// Code generated by "pkgimport -p cli -i github.com/urfave/cli/v2 -o cli.go"; DO NOT EDIT.
// Install by "go get -u -v github.com/wzshiming/gotype/cmd/pkgimport";
//go:generate pkgimport -p cli -i github.com/urfave/cli/v2 -o cli.go

package cli

import (
	origin "github.com/urfave/cli/v2"
)

// type
type (

	// UintFlag is a flag with type uint
	UintFlag = origin.UintFlag

	// Uint64Flag is a flag with type uint64
	Uint64Flag = origin.Uint64Flag

	// StringSliceFlag is a flag with type *StringSlice
	StringSliceFlag = origin.StringSliceFlag

	// StringSlice wraps a []string to satisfy flag.Value
	StringSlice = origin.StringSlice

	// StringFlag is a flag with type string
	StringFlag = origin.StringFlag

	// ShellCompleteFunc is an action to execute when the shell completion flag is set
	//ShellCompleteFunc = origin.ShellCompleteFunc

	// Serializeder is used to circumvent the limitations of flag.FlagSet.Set
	//Serializeder = origin.Serializeder

	// PathFlag is a flag with type string
	PathFlag = origin.PathFlag

	// OnUsageErrorFunc is executed if an usage error occurs. This is useful for displaying
	// customized usage error messages.  This function is able to replace the
	// original error messages.  If this function is not set, the "Incorrect usage"
	// is displayed and the execution is interrupted.
	OnUsageErrorFunc = origin.OnUsageErrorFunc

	// MultiError is an error that wraps multiple errors.
	MultiError = origin.MultiError

	// IntSliceFlag is a flag with type *IntSlice
	IntSliceFlag = origin.IntSliceFlag

	// IntSlice wraps an []int to satisfy flag.Value
	IntSlice = origin.IntSlice

	// IntFlag is a flag with type int
	IntFlag = origin.IntFlag

	// Int64SliceFlag is a flag with type *Int64Slice
	Int64SliceFlag = origin.Int64SliceFlag

	// Int64Slice is an opaque type for []int to satisfy flag.Value
	Int64Slice = origin.Int64Slice

	// Int64Flag is a flag with type int64
	Int64Flag = origin.Int64Flag

	// GenericFlag is a flag with type Generic
	GenericFlag = origin.GenericFlag

	// Generic is a generic parseable type identified by a specific flag
	Generic = origin.Generic

	// Float64SliceFlag is a flag with type *Float64Slice
	Float64SliceFlag = origin.Float64SliceFlag

	// Float64Slice is an opaque type for []float64 to satisfy flag.Value
	Float64Slice = origin.Float64Slice

	// Float64Flag is a flag with type float64
	Float64Flag = origin.Float64Flag

	// FlagsByName is a slice of Flag.
	FlagsByName = origin.FlagsByName

	// FlagStringFunc is used by the help generation to display a flag, which is
	// expected to be a single line.
	FlagStringFunc = origin.FlagStringFunc

	// Flag is a common interface related to parsing flags in cli.
	// For more advanced flag parsing techniques, it is recommended that
	// this interface be implemented.
	Flag = origin.Flag

	// ExitCoder is the interface checked by `App` and `Command` for a custom exit
	// code
	ExitCoder = origin.ExitCoder

	ErrorFormatter = origin.ErrorFormatter

	// DurationFlag is a flag with type time.Duration (see https://golang.org/pkg/time/#ParseDuration)
	DurationFlag = origin.DurationFlag

	// Context is a type that is passed through to
	// each Handler action in a cli application. Context
	// can be used to retrieve context-specific args and
	// parsed command-line options.
	Context = origin.Context

	CommandsByName = origin.CommandsByName

	// CommandNotFoundFunc is executed if the proper command cannot be found
	CommandNotFoundFunc = origin.CommandNotFoundFunc

	// CommandCategory is a category containing commands.
	CommandCategory = origin.CommandCategory

	CommandCategories = origin.CommandCategories

	// Command is a subcommand for a cli.App.
	Command = origin.Command

	// BoolFlag is a flag with type bool
	BoolFlag = origin.BoolFlag

	// BeforeFunc is an action to execute before any subcommands are run, but after
	// the context is ready if a non-nil error is returned, no subcommands are run
	BeforeFunc = origin.BeforeFunc

	// Author represents someone who has contributed to a cli project.
	Author = origin.Author

	Args = origin.Args

	// App is the main structure of a cli application.
	App = origin.App

	// AfterFunc is an action to execute after any subcommands are run, but after the
	// subcommand has finished it is run even if Action() panics
	AfterFunc = origin.AfterFunc

	// ActionFunc is the action to execute when no subcommands are specified
	ActionFunc = origin.ActionFunc
)

// Declaration
var (

	// VersionPrinter prints the version for the App
	VersionPrinter = origin.VersionPrinter

	VersionFlag = origin.VersionFlag

	// SubcommandHelpTemplate is the text template for the subcommand help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	SubcommandHelpTemplate = origin.SubcommandHelpTemplate

	// ShowVersion prints the version number of the App
	ShowVersion = origin.ShowVersion

	// ShowSubcommandHelp prints help for the given subcommand
	ShowSubcommandHelp = origin.ShowSubcommandHelp

	// ShowCompletions prints the lists of commands within a given context
	ShowCompletions = origin.ShowCompletions

	// ShowCommandHelpAndExit - exits with code after showing help
	ShowCommandHelpAndExit = origin.ShowCommandHelpAndExit

	// ShowCommandHelp prints help for the given command
	ShowCommandHelp = origin.ShowCommandHelp

	// ShowCommandCompletions prints the custom completions for a given command
	ShowCommandCompletions = origin.ShowCommandCompletions

	// ShowAppHelpAndExit - Prints the list of subcommands for the app and exits with exit code.
	ShowAppHelpAndExit = origin.ShowAppHelpAndExit

	// ShowAppHelp is an action that displays the help.
	ShowAppHelp = origin.ShowAppHelp

	// OsExiter is the function used when the app exits. If not set defaults to os.Exit.
	OsExiter = origin.OsExiter

	// NewStringSlice creates a *StringSlice with default values
	NewStringSlice = origin.NewStringSlice

	// NewIntSlice makes an *IntSlice with default values
	NewIntSlice = origin.NewIntSlice

	// NewInt64Slice makes an *Int64Slice with default values
	NewInt64Slice = origin.NewInt64Slice

	// NewFloat64Slice makes a *Float64Slice with default values
	NewFloat64Slice = origin.NewFloat64Slice

	// NewContext creates a new context. For use in when invoking an App or Command action.
	NewContext = origin.NewContext

	// InitCompletionFlag generates completion code
	//InitCompletionFlag = origin.InitCompletionFlag

	HelpPrinterCustom = origin.HelpPrinterCustom

	HelpPrinter = origin.HelpPrinter

	HelpFlag = origin.HelpFlag

	// HandleExitCoder checks if the error fulfills the ExitCoder interface, and if
	// so prints the error to stderr (if it is non-empty) and calls OsExiter with the
	// given exit code.  If the given error is a MultiError, then this func is
	// called on all members of the Errors slice and calls OsExiter with the last exit code.
	HandleExitCoder = origin.HandleExitCoder

	//GenerateCompletionFlag = origin.GenerateCompletionFlag

	FlagStringer = origin.FlagStringer

	// Exit wraps a message and exit code into an ExitCoder suitable for handling by
	// HandleExitCoder
	Exit = origin.Exit

	ErrWriter = origin.ErrWriter

	// DefaultAppComplete returns an ActionFunc to run a default command if non were passed.
	// Usage: `app.Action = cli.DefaultCommand("command")`
	//DefaultCommand = origin.DefaultCommand

	// DefaultAppComplete prints the list of subcommands as the default app completion method
	DefaultAppComplete = origin.DefaultAppComplete

	// CommandHelpTemplate is the text template for the command help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	CommandHelpTemplate = origin.CommandHelpTemplate

	// AppHelpTemplate is the text template for the Default help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	AppHelpTemplate = origin.AppHelpTemplate
)

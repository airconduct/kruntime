package app

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootLogger logr.Logger

func New() *cobra.Command {

	rootLogger = zapr.NewLogger(zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout), zap.InfoLevel,
	))).WithName("golett")

	rootCommand := &cobra.Command{
		Use: "golett",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	rootCommand.AddCommand(
		serverCommand(),
	)
	return rootCommand
}

type options interface {
	Complete(cmd *cobra.Command, args []string) error
	Run() error
}

func optionsRun(opts options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := opts.Complete(cmd, args); err != nil {
			return err
		}
		return opts.Run()
	}
}

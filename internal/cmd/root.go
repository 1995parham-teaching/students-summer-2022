package cmd

import (
	"log"
	"os"

	"github.com/1995parham-teaching/students/internal/cmd/migrate"
	"github.com/1995parham-teaching/students/internal/cmd/serve"
	"github.com/1995parham-teaching/students/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg := config.New()

	var (
		logger *zap.Logger
		err    error
	)

	if cfg.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}

	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "students",
		Short: "sample application for snapp-summer-2022",
	}

	root.AddCommand(serve.New(logger, cfg))
	root.AddCommand(migrate.New(logger, cfg))

	if err := root.Execute(); err != nil {
		logger.Error("failed to execute root command", zap.Error(err))
		os.Exit(ExitFailure)
	}
}

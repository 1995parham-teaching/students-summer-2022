package serve

import (
	"github.com/1995parham-teaching/students/internal/config"
	"github.com/1995parham-teaching/students/internal/db"
	"github.com/1995parham-teaching/students/internal/handler"
	"github.com/1995parham-teaching/students/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main(logger *zap.Logger, cfg config.Config) {
	app := echo.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		logger.Named("db").Fatal("cannot create a db instance", zap.Error(err))
	}

	var studentStore store.Student

	{
		logger := logger.Named("store")

		studentStore = store.NewStudentMongoDB(
			db, logger.Named("student"),
		)
	}

	{
		logger := logger.Named("http")

		ha := handler.Auth{
			Logger:   logger.Named("auth"),
			Name:     cfg.Admin.Name,
			Username: cfg.Admin.Username,
			Password: cfg.Admin.Password,
			Key:      []byte(cfg.Secret),
		}

		ha.Register(app.Group(""))

		hs := handler.Student{
			Store:  studentStore,
			Logger: logger.Named("student"),
		}

		hs.Register(app.Group("/api/students", ha.Auth))
	}

	app.Debug = cfg.Debug

	if err := app.Start(":1234"); err != nil {
		logger.Error("cannot start the http server", zap.Error(err))
	}
}

func New(logger *zap.Logger, cfg config.Config) *cobra.Command {
	// nolint: exhaustruct
	return &cobra.Command{
		Use:   "serve",
		Short: "runs http server for students api",
		Run: func(_ *cobra.Command, _ []string) {
			main(logger, cfg)
		},
	}
}

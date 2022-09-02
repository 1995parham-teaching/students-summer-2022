package migrate

import (
	"context"

	"github.com/1995parham-teaching/students/internal/config"
	"github.com/1995parham-teaching/students/internal/db"
	"github.com/1995parham-teaching/students/internal/store"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main(logger *zap.Logger, cfg config.Config) {
	db, err := db.New(cfg.Database)
	if err != nil {
		logger.Named("db").Fatal("cannot create a db instance", zap.Error(err))
	}

	idx, err := db.Collection(store.StudentCollection).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		logger.Fatal("cannot create an index", zap.Error(err))
	}

	logger.Info("database index", zap.Any("index", idx))
}

func New(logger *zap.Logger, cfg config.Config) *cobra.Command {
	// nolint: exhaustruct
	return &cobra.Command{
		Use:   "migrate",
		Short: "create indexes on mongodb",
		Run: func(cmd *cobra.Command, args []string) {
			main(logger, cfg)
		},
	}
}

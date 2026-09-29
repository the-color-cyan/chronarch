package main

import (
	"chronarch/internal/cli"
	"chronarch/internal/config"
	"chronarch/internal/storage"
	"context"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	ctx := context.Background()
	db, err := storage.Open(ctx, cfg.GetString("storage.path"))
	if err != nil {
		return err
	}
	defer db.Close()

	if err := storage.InitSchema(ctx, db); err != nil {
		return err
	}

	return cli.NewRootCmd().Execute()
}

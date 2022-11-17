package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := lint(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func lint(ctx context.Context) error {
	// daggerクライアントの初期化
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// `node` コンテナイメージを持ってくる
	node := client.Container().From("node:18.11.0")

	// `node` コンテナでnodeのバージョンを出す
	node = node.Exec(dagger.ContainerExecOpts{
		Args: []string{"node", "-v"},
	})

	if _, err := node.ExitCode(ctx); err != nil {
		panic(err)
	}

	return nil
}

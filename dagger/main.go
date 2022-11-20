package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
		if err != nil {
		panic("Error loading .env file")
	}

	linterTargetDir := os.Getenv("LINTER_TARGET_DIR")

	if err := lint(context.Background(), linterTargetDir); err != nil {
		fmt.Println(err)
	}
}

func lint(ctx context.Context, targetDir string) error {
	// daggerクライアントの初期化
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout), dagger.WithWorkdir("../"))
	if err != nil {
		return err
	}
	defer client.Close()

	// workdirのパスを取得
	projectRoot := client.Host().Workdir()

	// `node` コンテナイメージを持ってくる
	node := client.Container().From("node:18.11.0-alpine3.15")

	// `node` コンテナの/srcにworkdirをマウントする
	node = node.WithMountedDirectory("/src", projectRoot).WithWorkdir("/src")

	// `node` コンテナ内で必要なnpmパッケージをインストールする
	node = node.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"npm", "ci"},
		})

	// `node` コンテナ内でtextlintを流す
	node = node.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"npx", "textlint", "--fix", "--dry-run", targetDir},
		})

	if _, err := node.ExitCode(ctx); err != nil {
		panic(err) // TODO: panicすべきではなさそうなので、あとで直す
	}

	return nil
}

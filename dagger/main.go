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
			Args: []string{"npx", "textlint", "--fix", "--dry-run", "README.md"},
		})

	if _, err := node.ExitCode(ctx); err != nil {
		panic(err) // TODO: panicすべきではなさそうなので、あとで直す
	}

	return nil
}

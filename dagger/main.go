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
	node := client.Container().From("node:18.11.0")

	// `node` コンテナの/srcにworkdirをマウントする
	node = node.WithMountedDirectory("/src", projectRoot).WithWorkdir("/src")

	// `node` コンテナでworkdirをls(これは消す)
	node = node.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"ls"},
		})

	if _, err := node.ExitCode(ctx); err != nil {
		panic(err)
	}

	return nil
}

package main

import (
	"context"

	"github.com/CompaniesInfoStore/cmd/app"
)

func main() {

	app := app.Application{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Init(ctx)
	app.Start()

}

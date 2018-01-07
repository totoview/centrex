package main

//go:generate protoc pb/CentrexMsg.proto --go_out=plugins=grpc:.
//go:generate go run generate.go

import (
	"fmt"
	"os"

	"github.com/totoview/centrex/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

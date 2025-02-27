//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	goPath := os.Getenv("GOPATH")

	if err := entc.Generate("./schema",
		&gen.Config{
			Target:  fmt.Sprintf("%v/src/github.com/grpc-kit/pkg/lion", goPath),
			Package: "github.com/grpc-kit/pkg/lion",
		},
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

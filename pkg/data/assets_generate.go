// +build ignore

package main

import (
	"log"
	"github.com/shurcooL/vfsgen"

	"github.com/gargath/vfsgen-example/pkg/data"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatalln(err)
	}
}

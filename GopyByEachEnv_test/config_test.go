package GopyByEachEnv_test

import (
	"fmt"

	GopyByEachEnv "github.com/sky0621/work-go-GopyByEachEnv"
)

func ExampleReadConfig() {
	config := GopyByEachEnv.ParseConfig("../cmd/config.json")
	fmt.Println(len(config.CopySpecs))
	// Output:
	// 2
}

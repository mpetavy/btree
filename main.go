package main

import (
	"fmt"
	"github.com/mpetavy/common"
)

func init() {
	common.Init(false, "1.0.0", "", "", "2022", "Btree tryout", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

func run() error {
	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}

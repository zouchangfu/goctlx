package env

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zouchangfu/goctlx/pkg/env"
)

func write(_ *cobra.Command, _ []string) error {
	if len(sliceVarWriteValue) > 0 {
		return env.WriteEnv(sliceVarWriteValue)
	}
	fmt.Println(env.Print())
	return nil
}

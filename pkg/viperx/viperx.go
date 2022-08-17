package viperx

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CobraToViperName(cobraName string) (viperName string) {
	return strings.ReplaceAll(cobraName, "-", ".")
}

func GetString(name string) string {
	return viper.GetString(CobraToViperName(name))
}

func GetBool(name string) bool {
	return viper.GetBool(CobraToViperName(name))
}

func GetInt(name string) int {
	return viper.GetInt(CobraToViperName(name))
}

func BindCobraFlag(flag string, cmd *cobra.Command) error {
	return viper.BindPFlag(CobraToViperName(flag), cmd.Flags().Lookup(flag))
}

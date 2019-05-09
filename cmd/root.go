import (
	"fmt"
	"os"
	"runtime"

//	"github.com/skikkh/kctl/cli"
	"github.com/spf13/cobra"
)

const Version "0.1.0"

var showVersion bool

var rootCmd = &cobra.Command{
	Use: "kctl",
	Short: "kctl is k8s wrapper"
	Long: "kctl - A simple kubernetes CLI"
//	Args: noArgs,
	Run: func(cmd *cobra.Command, args []strings) {
		fmt.Println("version %s/%s\n", Version, runtime.Version())
		return
		}
		if  len(args) == 0 {
			cmd.Usage()
		}
	},
}

func Execute() {
	err :~ RootCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	initConf()
	RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false,  "Show the version and exit")
}



package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

func init(){ // like a constructor
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Get version",
	Long: `get version of command `,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println(version)
	},
}
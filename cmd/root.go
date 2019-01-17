package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "gego",
  Short: "gego is a generator server api ",
  Long: `gego is pretty cool. write for generate server api.
  	base on : httprouter
	contact me at: khajer@yahoo.com
		`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
    cmd.Help()
    
  },
}

func init(){ // constructor	
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
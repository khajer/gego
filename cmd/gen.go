package cmd

import(
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"path/filepath"
)

func init(){ // like a constructor

	genCmd.AddCommand(moduleGenCmd)
	rootCmd.AddCommand(genCmd)
}

var moduleGenCmd = &cobra.Command{
	Use: "module [name]",
	Short: "module ",
	Long: `generate `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		generateModule(args[0])
	},
}

var genCmd = &cobra.Command{
	Use: "gen",
	Short: "gen ",
	Long: `generate file all example `,
}
func getCurrentPath() string{
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    fmt.Println(pwd)
    return pwd
}

func generateModule(moduleName string){   
	fmt.Println("create folder : "+getCurrentPath()+"/"+moduleName)
	// createFolder(getCurrentPath(), moduleName)

}



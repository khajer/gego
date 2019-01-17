package cmd
func init(){ // like a constructor
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use: "gen [name]",
	Short: "gen project",
	Long: `gen module `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		// moduleName := args[0]
		println(args)		
	},
}

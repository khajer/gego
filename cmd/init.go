package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	// "errors"
	"os"
	"path/filepath"	
	"text/template"
	"bytes"
)

func init(){ // like a constructor
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use: "init [name]",
	Short: "init project",
	Long: `create new project with name `,
	Args: cobra.MinimumNArgs(1),
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires at least one arg")
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string){
		projectName := args[0]
		fmt.Println("init project : "+projectName)
		initProject(projectName)		
	},
}


func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}


func initProject(projectName string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newDir := filepath.Join(pwd, projectName)
	found, _ := exists(newDir)
	if found == true{
		fmt.Println("non empty directory :"+newDir)
		return
	}
		
	generateFile(newDir, projectName)
}

func writefile(path string, strData string) bool{
	f, err := os.Create(path)
	if err != nil{
		panic(err)
		return false
	}
	defer f.Close()

	_, err = f.WriteString(strData)
	if err != nil{
		panic(err)
		return false
	}
	return true
}

func generateFile(pathProject string, projectName string){
	os.MkdirAll(pathProject, os.ModePerm)
	if writefile(filepath.Join(pathProject, "main.go") , genFileMain(projectName)) {
		println("generate file :"+filepath.Join(pathProject, "main.go"))
	}
	if writefile(filepath.Join(pathProject, "go.mod") , genFileGoMod(projectName)){
		println("generate file :"+filepath.Join(pathProject, "go.mod"))
	}
	
	pathRoute := filepath.Join(pathProject, "route") 	
	os.MkdirAll(pathRoute, os.ModePerm)
	if writefile(filepath.Join(pathRoute, "route.go") , genFileRoute(projectName)){
		println("generate file :"+filepath.Join(pathRoute, "route.go"))		
	}

	pathHandler := filepath.Join(pathProject, "handler")
	os.MkdirAll(pathHandler, os.ModePerm)
	if writefile(filepath.Join(pathHandler, "handler.go") , genFileHandler(projectName)){
		println("generate file :"+filepath.Join(pathHandler, "handler.go"))		
	}

	pathApp := filepath.Join(pathProject, "app")
	os.MkdirAll(pathApp, os.ModePerm)
	if writefile(filepath.Join(pathApp, "app.go") , genFileApp(projectName)){
		println("generate file :"+filepath.Join(pathApp, "app.go"))		
	}

}

func genFileApp(projectName string) string{
	tmp := `package app
import (

)

`
	return tmp	
}

func genFileHandler(projectName string) string{
	tmp := `package handler
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)
type Handler struct{
	Method 	string
	Url 	string
	Func    func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
`
	return tmp		
}

func genFileRoute(projectName string) string{
	tmp := `package route
import (
	"github.com/julienschmidt/httprouter"	
	"{{ .}}/handler"
	"fmt"
	"strings"
)
var arHandler []*handler.Handler

func AddHandle(handler ...*handler.Handler){	
	for _, h := range handler{
		arHandler = append(arHandler, h)	
	}
}

func Routes() *httprouter.Router{	
	r := httprouter.New()
	for _ ,v := range arHandler{
		if strings.ToUpper(v.Method) == "GET" {
			fmt.Println("[GET] URL: "+v.Url)
			r.GET(v.Url, v.Func)
		}else if strings.ToUpper(v.Method) == "POST" {
			fmt.Println("[POST] URL: "+v.Url)
			r.POST(v.Url, v.Func)
		}
	}
	return r
}`
	funcMap := template.FuncMap{}
	var tpl bytes.Buffer
	t := template.Must(template.New("templateFile").Funcs(funcMap).Parse(tmp))
	err := t.Execute(&tpl, projectName)
	if err!= nil{
		panic(err)
	}
	return tpl.String()	

}

func genFileGoMod(projectName string) string{
	return "module "+projectName
}

func genFileMain(projectName string) string{
	tmp :=`package main
import (	
	"net/http"
	"log"
	"{{ .}}/route"		
	_ "{{ .}}/app"
)
func main(){	
    port := ":8080"	
    println("start server at localhost "+port)
	log.Fatal(http.ListenAndServe(port, route.Routes()))
}
`
	funcMap := template.FuncMap{}
	var tpl bytes.Buffer
	t := template.Must(template.New("templateFile").Funcs(funcMap).Parse(tmp))
	err := t.Execute(&tpl, projectName)
	if err!= nil{
		panic(err)
	}
	return tpl.String()	
}




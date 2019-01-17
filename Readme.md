
# how to build
run build.sh or manual install
`
	go build -o gego . 
	sudo cp gego /usr/local/bin/
`

# how to use 
## create project
`
gego new example.com
`

# structure folder
```
/app
	app.go
	feature1 
		feature1.go
		feature1_test.go
	version
		version.go
		version_test.go 
/conf
	conf.go

/handler
	handler.go
/route
	route.go

main.go
config.yml
go.mod
go.sum
```

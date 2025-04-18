`go run $(go list -m -f {{.Dir}}  github.com/hopeio/protobuf)/tools/install_tools.go`
`protogen.exe go -e -w -v -p cherrywithinit/proto -o cherrywithinit/proto`
补全local.toml的配置
`go run cherrywithinit/main.go -c cherrywithinit/config.toml`
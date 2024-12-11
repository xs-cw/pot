module github.com/go-pot/pot

go 1.23.1

replace github.com/wzshiming/go-swagger => github.com/xs-cw/go-swagger v0.0.0-20241209121728-71a0d4beb73f
replace (
	github.com/go-pot/sessions => ../sessions
)
require (
	github.com/beego/beego v1.12.14
	github.com/fsnotify/fsnotify v1.6.0
	github.com/garyburd/redigo v1.6.4
	github.com/go-pot/sessions v1.4.3
	github.com/gorilla/mux v1.8.1
	github.com/unrolled/render v1.7.0
	github.com/urfave/cli/v2 v2.27.5
	github.com/urfave/negroni v1.0.0
	github.com/wzshiming/cache v0.0.0-20171220122851-ae83cd3c4aa0
	github.com/wzshiming/go-swagger v1.2.1
	gopkg.in/configer.v1 v1.1.0
	gopkg.in/ffmt.v1 v1.5.6
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.4.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/wzshiming/fork v0.0.0-20191212125603-a884e312d3a0 // indirect
	github.com/wzshiming/llrb v1.0.1 // indirect
	github.com/wzshiming/task v2.2.1+incompatible // indirect
	github.com/wzshiming/walk v1.1.1 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/go-pot/sessions.v1 v1.4.3 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/pot.v1 v1.0.0-20180719094846-de5d4864a366 // indirect
)

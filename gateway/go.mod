module gateway

go 1.18

require (
	github.com/99designs/gqlgen v0.17.10
	github.com/go-redis/redis/v9 v9.0.0-beta.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/rs/cors v1.8.2
	github.com/vektah/gqlparser/v2 v2.4.5
	google.golang.org/protobuf v1.28.0
	services/comments v1.0.0
)

replace services/comments => ../comments

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/matryer/moq v0.2.7 // indirect
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.8.1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20220707150051-590a5ac7bee1 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

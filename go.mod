module github.com/liubaninc/m0

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/alexandrevicenzi/unchained v1.3.0
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/cosmos/go-bip39 v1.0.0
	github.com/dchest/captcha v0.0.0-20200903113550-03f5f0333e1f
	github.com/ddliu/motto v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/go-units v0.4.0
	github.com/fsouza/go-dockerclient v1.7.2
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.2
	github.com/go-kit/kit v0.10.0
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/prometheus/client_golang v1.10.0
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	go.opencensus.io v0.22.6 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/cosmos/cosmos-sdk => ./cosmos-sdk

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

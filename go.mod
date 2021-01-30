module github.com/shunail2029/SecretDB

go 1.15

// TODO: change
replace github.com/shunail2029/SecretDB-master => ../SecretDB-master

require (
	github.com/CosmWasm/wasmd v0.11.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gorilla/mux v1.8.0
	github.com/shunail2029/SecretDB-master v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.7
	github.com/tendermint/tm-db v0.5.1
	go.mongodb.org/mongo-driver v1.4.2
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/yaml.v2 v2.4.0
)

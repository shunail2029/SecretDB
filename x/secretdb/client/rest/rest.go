package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers secretdb-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/secretdb/item", createItemHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/secretdb/item/{key}", getItemHandler(cliCtx, "secretdb")).Methods("GET")
	r.HandleFunc("/secretdb/item/{key}", getItemsHandler(cliCtx, "secretdb")).Methods("GET")
	r.HandleFunc("/secretdb/item", setItemHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/secretdb/item", setItemsHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/secretdb/item", deleteItemHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc("/secretdb/item", deleteItemsHandler(cliCtx)).Methods("DELETE")
}

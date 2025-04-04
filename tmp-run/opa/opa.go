package opa

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
)

func OPA() {

	_ = genPolicy("admin")
	ctx := context.Background()
	input_data := map[string]interface{}{
		"username": "admin",
		"action":   "GET",
		"domain":   "global",
		"object":   "/api/v1/users",
	}
	fmt.Printf("input data: %v\n", input_data)
	fmt.Printf("policies: %v\n", policies)
	//store := inmem.NewFromReader(bytes.NewBufferString(policies))
	store := inmem.NewFromObject(policies)

	r := rego.New(
		rego.Query("data.authz.allow"),
		rego.Module("authz.rego", module),
		rego.Store(store),
		rego.Input(input_data),
	)
	res, err := r.Eval(ctx)
	if err != nil || len(res) != 1 || len(res[0].Expressions) != 1 {
		fmt.Println("opa error: 500")
	}
	fmt.Println("opa result: ", res[0].Expressions[0].Value)
	if res[0].Expressions[0].Value == true {
		fmt.Println("opa result: allow")
	}

}

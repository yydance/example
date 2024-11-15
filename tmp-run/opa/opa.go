package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
)

func OPA() {
	ctx := context.Background()
	module, err := os.ReadFile("./opa/models.rego")
	if err != nil {
		panic("Read model failed")
	}
	data, err := os.ReadFile("./opa/example_data.json")
	if err != nil {
		panic("Read data failed")
	}
	var opaData map[string]interface{}
	if err := json.Unmarshal(data, &opaData); err != nil {
		panic("Unmarshal data failed")
	}
	input_data := map[string]interface{}{
		"username": "damon",
		"action":   "GET",
		"domain":   "global",
		"object":   "/api/v1/users",
	}
	store := inmem.NewFromObject(opaData)
	r := rego.New(
		rego.Query("data.authz.allow"),
		rego.Module("example.rego", string(module)),
		rego.Store(store),
		rego.Input(input_data),
	)
	res, err := r.Eval(ctx)
	if err != nil || len(res) != 1 || len(res[0].Expressions) != 1 {
		panic("Eval failed")
	}
	fmt.Println("res: ", res)
	if res[0].Expressions[0].Value == true {
		fmt.Println("Allow")
	}

}

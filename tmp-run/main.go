package main

import (
	"fmt"
)

var (
	apisix_admin = "http://10.252.9.198:45975/apisix/admin"
	apisix_token = "5b54e554ed45426d9af01528b33661f1"
	upstream_url = apisix_admin + "/upstreams"
)

type Upstream struct {
	Id          string
	Name        string
	Type        string
	Desc        string
	CreatedTime int64
	UpdatedTime int64
}

func main() {
	/*
		client := resty.New()
		resp, err := client.R().SetHeader("X-API-KEY", apisix_token).Get(upstream_url)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", resp)

		//ids, _, _, err := jsonparser.Get(resp.Body(), "list")
		//if err != nil {
		//	panic(err)
		//}
		fmt.Println("********************")

		//此处，取到数据直接插入MySQL，每次一行
		_, _ = jsonparser.ArrayEach(resp.Body(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			//fmt.Printf("each, value: %s\t Type: %s\n", string(value), dataType)
			fmt.Println("---------------")
			id, _ := jsonparser.GetString(value, "value", "id")
			name, _ := jsonparser.GetString(value, "value", "name")
			upstream_type, _ := jsonparser.GetString(value, "value", "type")
			desc, _ := jsonparser.GetString(value, "value", "desc")
			create_time, _ := jsonparser.GetInt(value, "value", "create_time")

			fmt.Printf("id: %s, name: %s, upstream_type: %s, desc: %s, create_time: %d\n", id, name, upstream_type, desc, create_time)
			//fmt.Printf("each, value: %s\t Type: %s\n", string(value), dataType)
		}, "list")
	*/
	errMsgs := make([]string, 0)
	if len(errMsgs) == 0 {
		fmt.Print("nil\n")
	}
	fmt.Printf("%v\n", errMsgs)
	test := fmt.Sprintf("%%%s%%", "name")
	fmt.Println(test)
	//fmt.Println(string(ids))
}

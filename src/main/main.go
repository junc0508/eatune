package main

import (
	"fmt"
	"github.com/junc0508/eatune/src/price"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Extra_vars struct {
	Account_id int
	End_point  string
	Token      string
}

func main() {
	buf, err := ioutil.ReadFile("./src/extra_vars/extra_vars.yml")
	if err != nil {
		panic(err)
	}

	var extra_vars Extra_vars
	err = yaml.Unmarshal(buf, &extra_vars)
	fmt.Println(extra_vars)
	body := price.Get_current_rate(extra_vars.End_point, extra_vars.Token)
	fmt.Println(string(body))
	
}

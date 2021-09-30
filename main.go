package main

import (
	"fmt"
	"strconv"
	"strings"

	"tudai2021.com/model"
)

func main() {
	res := createResult("TX03ABC")
	fmt.Println(res)
}

func createResult(param string) model.Result {
	values := strings.Split(param, "")
	t := values[0] + values[1]
	l := values[2] + values[3]
	v := values[4:]
	lengthFinal, _ := strconv.Atoi(l)
	return model.NewResult(t, lengthFinal, strings.Join(v, ""))
}

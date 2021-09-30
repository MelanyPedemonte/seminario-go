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

	res2 := createResult("NN0512345")
	fmt.Println(res2)

	res3 := createResult("NN101234567890")
	fmt.Println(res3)
}

func createResult(param string) model.Result {
	values := strings.Split(param, "")
	t := values[:2]
	l := values[2:4]
	v := values[4:]
	lengthFinal, _ := strconv.Atoi(strings.Join(l, ""))
	return model.NewResult(strings.Join(t, ""), lengthFinal, strings.Join(v, ""))
}

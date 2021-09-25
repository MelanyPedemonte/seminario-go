package main 
import (
	"fmt"
	"tudai2021.com/model"
)  

func main(){
	res := model.NewResult("AX", 3, "ABC")
	fmt.Println(res)
}
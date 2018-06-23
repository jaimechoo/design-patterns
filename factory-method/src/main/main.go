package main

import (
	"fmt"

	"product"
)

func main() {
	var phonefatory product.Phonefatory

	phonefatory = new(product.XiaomiFactory)
	xiaomiphone, err := phonefatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	xiaomiphone.Call()

	phonefatory = new(product.HuaweiFactory)
	huaweiphone, err := phonefatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	huaweiphone.Call()

	phonefatory = new(product.AppleFactory)
	applephone, err := phonefatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	applephone.Call()
}

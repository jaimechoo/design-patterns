package main

import (
	"fmt"

	"product"
)

func main() {
	xiaomiphone, err := product.CreatPhone(product.Xiaomiphone)
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	xiaomiphone.Call()

	huaweiphone, err := product.CreatPhone(product.Huaweiphone)
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	huaweiphone.Call()

	applephone, err := product.CreatPhone(product.Applephone)
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	applephone.Call()
}

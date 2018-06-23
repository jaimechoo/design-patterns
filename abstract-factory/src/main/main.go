package main

import (
	"fmt"

	"product"
)

func main() {
	var electronicfatory product.Electronicfatory

	electronicfatory = new(product.XiaomiFactory)
	xiaomiphone, err := electronicfatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	xiaomiphone.Call()

	xiaomicomputer, err := electronicfatory.CreateComputer()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	xiaomicomputer.Work()

	electronicfatory = new(product.HuaweiFactory)
	huaweiphone, err := electronicfatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	huaweiphone.Call()

	huaweicomputer, err := electronicfatory.CreateComputer()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	huaweicomputer.Work()

	electronicfatory = new(product.AppleFactory)
	applephone, err := electronicfatory.CreatePhone()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	applephone.Call()

	applecomputer, err := electronicfatory.CreateComputer()
	if err != nil {
		fmt.Println("CreatPhone failed :", err.Error())
		return
	}
	applecomputer.Work()
}

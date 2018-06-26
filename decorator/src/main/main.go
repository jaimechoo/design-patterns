package main

import (
	"decorator"
	"fmt"
)

func main() {
	b1 := &decorator.StringDisplayStu{Str: "Hello,world."}
	b2 := &decorator.SideBorderStu{Display: b1, BorderChar: "#"}
	b3 := &decorator.FullBorderStu{Display: b2}
	fmt.Println(b1.GetColums())
	fmt.Println(b2.GetColums())
	fmt.Println(b3.GetColums())
	b4 := &decorator.SideBorderStu{
		Display: &decorator.FullBorderStu{
			Display: &decorator.FullBorderStu{
				Display: &decorator.SideBorderStu{
					Display: &decorator.FullBorderStu{
						Display: &decorator.StringDisplayStu{
							Str: "你好，世界。",
						},
					},
					BorderChar: "*",
				},
			},
		},
		BorderChar: "/",
	}
	fmt.Println(b4.GetColums())
}

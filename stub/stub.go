package stub

import "fmt"

// 桩，或称桩代码，是指用来代替关联代码或者未实现代码的代码。如果函数 B 用 B1 来代替，那么，B 称为原函数，B1 称为桩函数。打桩就是编写或生成桩代码。
var myName = "yanshuifa"

func ShowAndGetMyName() string {
	fmt.Println(myName)

	return myName
}

var show = func(name string) string {
	fmt.Println(name)
	return name
}
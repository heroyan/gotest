package mock

import "fmt"

// 全局变量
var global1 = 100

// 使用GoMonkey打桩测试

type MyCls struct {
	name string
	age  int
}

func (cls *MyCls) SayHi() {
	fmt.Println("Hi, i am " + cls.name)
}

func (cls *MyCls) Add(a, b int) int {
	return a + b
}

func (cls *MyCls) GetName() string {
	return cls.name
}

func (cls *MyCls) GetAge() int {
	return cls.age
}

func NewMyCls(name string, age int) *MyCls {
	return &MyCls{
		name: name,
		age: age,
	}
}

func MonkeyFunc(a, b int) int {
	return a*b
}
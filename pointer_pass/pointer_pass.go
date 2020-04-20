package main

type mytype struct {
	name string
}

func newMytype(name string) *mytype {
	t := mytype{name: name}
	return &t
}

func function1(obj *mytype) {
	println(obj.name)
}

func function2(obj mytype) {
	println(obj.name)
}

func main() {
	myobj := *newMytype("hello")
	println(myobj.name)
	println(&myobj.name)
	println("Function 1")
	function1(&myobj)
	println("Function 2")
	function2(myobj)
}

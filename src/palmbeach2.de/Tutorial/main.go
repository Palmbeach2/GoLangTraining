package main

import (
	"fmt"
	"net/http"
)

//Für 7 (Webserver)
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello From Server")
}

func webserver() {

	fmt.Println("Webserver")
	http.HandleFunc("/hello", hello)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}

//Für 8 (Pointer)
func pointer() {
	i := 2
	fmt.Println(i)
	fmt.Println(&i)
	pointer := &i
	fmt.Println(pointer)
	fmt.Println(*pointer)
}

func editNumberCopy(j int) {
	j = 27
}

func editNumberPointer(j *int) {
	*j = 27
}

func pointer2() {
	i := 2
	fmt.Println(i)
	editNumberCopy(i)
	fmt.Println(i)
	editNumberPointer(&i)
	fmt.Println(i)

	value1 := 27
	value2 := 42
	p1 := &value1
	p2 := &value2
	fmt.Println(*p1, *p2)
	swapPointers(&p1, &p2)
	fmt.Println(*p1, *p2)
}

func swapPointers(p1 **int, p2 **int) {
	tmp := *p1
	*p1 = *p2
	*p2 = tmp
}

func main() {
	fmt.Println("Go Tutorial 8")
	pointer2()
}

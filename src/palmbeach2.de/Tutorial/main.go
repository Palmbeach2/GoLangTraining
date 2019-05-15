package main

import (
	"fmt"
	"math"
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

//Für 9 (Arrays and slice)
func tut9() {
	fmt.Println("Go Tutorial 9")
	var buffer [256]int
	for i := 0; i < len(buffer); i++ {
		buffer[i] = i
	}

	slice := buffer[27:72]
	fmt.Println(cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	copy(newSlice, slice)
	slice = newSlice
	fmt.Println(cap(slice))
}

//Für 10 (Variable ANzahl an Argumenten)
func printWithNewLines(lines ...string) {
	fmt.Println("Go Tutorial 10")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func tut10() {
	printWithNewLines("a", "b", "c", "d")
	slice := []string{"string1", "string2", "string3", "string4"}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	printWithNewLines(slice...)
}

//Für 11 (Strukturen)Vector3
func tut11() {
	fmt.Println("Go Tutorial 11")
	position := Vector3{1, 2, 3}
	position2 := Vector3{y: 42}
	position.x = 27
	pointer := &position
	pointer.z = 1
	fmt.Println(*pointer)
	fmt.Println(position2)
}

//Für 12 (Hashmap)
func tut12() {
	fmt.Println("Go Tutorial 12")
	hashmap := make(map[string]int)
	hashmap["Walter"] = 5
	hashmap["Gertrude"] = 27
	highscore := hashmap["Walter"]
	highscore2, exists := hashmap["Gertrude"]
	if !exists {
		fmt.Println("Highscore of Gertrude: doesn't exists")
	} else {
		fmt.Println("Highscore of Gertrude:", highscore2)
	}
	fmt.Println("Highscore of Gertrude:", hashmap["Gertrude"])
	fmt.Println("Highscore of Walter:", highscore)

	for key, value := range hashmap {
		fmt.Println("Key:", key, "Value:", value)
	}

	delete(hashmap, "Gertrude")
	highscore3, exists := hashmap["Gertrude"]
	if !exists {
		fmt.Println("Highscore of Gertrude: doesn't exists")
	} else {
		fmt.Println("Highscore of Gertrude:", highscore3)
	}

	cheaters := map[string]bool{
		"Pilzschaf":  false,
		"cheaterboy": true,
		"Palmbeach2": false,
	}
	fmt.Println("Cheaters:", cheaters)
}

//Vector3 a 3d Vector
type Vector3 struct {
	x float64
	y float64
	z float64
}

//Für 13 (Methoden)
func tut13() {
	fmt.Println("Go Tutorial 13")
	v := Vector3{1, 2, 3}
	v.SetNull()
	fmt.Println(v.Length())
}

// Length Returns length of Vector
func (v Vector3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

// SetNull Sets all componets of Vector to 0
func (v *Vector3) SetNull() {
	v.x = 0
	v.y = 0
	v.z = 0
}

func main() {
	tut13()

}

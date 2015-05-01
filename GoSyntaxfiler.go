package main

import (
	"fmt"
)

func main(){
	fmt.Println("text")
}



PACKAGE-er
	TIME
	time.Now()
	time.Sleep(100 * time.Millisecond)
	MATH/RAND
	rand.Intn(23)
	MATH
	math.Nextafter(2,3)

GENERELT
	for å bruke funksjoner i andre pakker bruk stor bokstav. Eksempel math.Pi for å bruke pi fra math package-en. 
	defer kan brukes for å utsette returkallet fra en funksjon
	Kan ha funksjoner inne funksjoner. Se Go tour 20/23. 
TYPER
	int
	string
	bool
	int64
	float64
	uint64
	const PI=3.14
	
	

	var x=10 eller x:=10	Dete må gjøres når du oppretter en variabel. 
	x=10					Etter x er opprettet. 
	x,y int  tilsvarer: x int, y int 		:x kan kun brukes inni en funksjon
	var i int				Dette tilsvarer at i=0
	int(f)					Konverter til int
	type MyFloat float64			Lager typen MyFloat

	

CHANNELS
	c := make(chan int)
	ch <- v    // Send v to channel
	v := <-ch  // Receive from ch, and assign value to v.
	x, y := <-c, <-c // receive from c
	c := make(chan int, 2)		Buffret channel. 
	close(c)			Gjøres fra sender. 
	for i:=range c{}		Kan iterere over channel, channel må være closed. 			
	
	
	

FUNKSJONER
func add(x int, y int) int {
	return x + y
}

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func main() {
	var state int=10
	switch state{
	case 10:
		fmt.Println(300)
	default:
		fmt.Println(1010)
	}
}

//DEFER
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

PEKERE
func main() {
	i:= 42
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
}

STRUCT
type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

ARRAY
func main() {
	p := [][]int{{2, 3, 5},{ 7, 11, 13}}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		for j:=0;j<len(p[i]);j++{
			fmt.Println( i, j, p[i][j])
		}
	}
}
p := []int{2, 3, 5, 7, 11, 13}
	u:=p[:4]
	v:=p[3:6]
a := make([]int, 0,5)					//make oppretter en slice med 0 elementer og kan maks ha 5. 

ITERERE OVER LISTE
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
	for i, v := range pow {					// i index, v value
		fmt.Printf("2**%d = %d\n", i, v)
	}
}


CLOSURE: fancy greier. 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


METHODS
type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}


CHANNALS and GORUTINES
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {		
		fmt.Println(i)
	}
}


func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

INHERITANCE GO
type Car struct {
    wheelCount int
}

func (car Car) numberOfWheels() int {
    return car.wheelCount
}

type Ferrari struct {
    Car
}

// a behavior only available for the Ferrari
func (f Ferrari) sayHiToSchumacher() {
    fmt.Println("Hi Schumacher!")
}

type AstonMartin struct {
    Car
}

// a behavior only available for the AstonMartin
func (a AstonMartin) sayHiToBond() {
    fmt.Println("Hi Bond, James Bond!")
}

func main() {
    f := Ferrari{Car{4}}
    fmt.Println("A Ferrari has this many wheels: ", f.numberOfWheels()) //has car behavior
    f.sayHiToSchumacher() //has Ferrari behavior

    a := AstonMartin{Car{4}}
    fmt.Println("An Aston Martin has this many wheels: ", a.numberOfWheels()) //has car behavior
    a.sayHiToBond() //has AstonMartin behavior


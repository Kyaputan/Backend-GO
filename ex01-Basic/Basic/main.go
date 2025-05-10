package main

import (
	"fmt"
	"time"
)

// ------------- ส่วนที่ 1: ตัวแปร พื้นฐาน ฟังก์ชัน -------------
func main() {
	fmt.Println("เริ่มเรียน Go")

	// การประกาศตัวแปร
	var a int = 10
	b := 20
	c := a + b
	fmt.Println("ผลรวม:", c)

	// ฟังก์ชันทั่วไป
	result := add(5, 15)
	fmt.Println("add(5,15):", result)

	// การใช้ if-else
	if c > 20 {
		fmt.Println("มากกว่า 20")
	} else {
		fmt.Println("ไม่เกิน 20")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("รอบ:", i)
	}

	// array
	nums := [3]int{1, 2, 3}
	for i, v := range nums {
		fmt.Printf("nums[%d] = %d\n", i, v)
	}

	// slice
	langs := []string{"Go", "Python", "C++"}
	langs = append(langs, "Rust")
	for index, lang := range langs {
		fmt.Println(index, "ภาษา:", lang)
	}

	// map
	ages := map[string]int{"Alice": 30, "Bob": 25}
	ages["Charlie"] = 35
	for name, age := range ages {
		fmt.Printf("%s อายุ %d\n", name, age)
	}

	// struct และ method
	p := Person{"Caption", 22}
	fmt.Println(p.Greet())

	// switch
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("ยอดเยี่ยม")
	case "B":
		fmt.Println("ดีมาก")
	default:
		fmt.Println("ปรับปรุงได้")
	}

	// Goroutine
	go printHello()
	time.Sleep(1 * time.Second)

	// Pointer
	x := 100
	incnot(x)
	fmt.Println("ก่อนเพิ่มค่า:", x)
	inc(&x)
	fmt.Println("หลังเพิ่มค่า:", x)

	// Interface
	var s Shape = Rectangle{width: 5, height: 10}
	fmt.Println("พื้นที่:", s.Area())

	// Error Handling
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ผลลัพธ์การหาร:", res)
	}

	// Channel
	ch := make(chan string)
	go func() {
		ch <- "ข้อความจาก channel"
	}()
	msg := <-ch
	fmt.Println(msg)

	// Select
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "จาก ch1"
	}()
	go func() {
		time.Sleep(700 * time.Millisecond)
		ch2 <- "จาก ch2"
	}()

	select {
	case m1 := <-ch1:
		fmt.Println("รับจาก ch1:", m1)
	case m2 := <-ch2:
		fmt.Println("รับจาก ch2:", m2)
	}

	// การใช้ defer
	defer fmt.Println("สิ้นสุดโปรแกรม")

	// Anonymous Function
	square := func(x int) int {
		return x * x
	}
	fmt.Println("square(4):", square(4))

	// Closure
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	// Recursion
	fmt.Println("factorial(5):", factorial(5))
}

func add(x, y int) int {
	return x + y
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("สวัสดี ผมชื่อ %s อายุ %d ปี", p.Name, p.Age)
}

func printHello() {
	fmt.Println("สวัสดีจาก goroutine")
}

func inc(x *int) {
	*x += 1
}

func incnot(x int) {
	x += 1
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("หารด้วยศูนย์ไม่ได้")
	}
	return a / b, nil
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

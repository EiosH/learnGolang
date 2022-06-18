package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// "log"

// "example.com/greetings"
//Animal 动物
type Animal struct {
	name string
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func test() {

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}

	data, _ := json.Marshal(d1)

	fmt.Printf("json:%s\n", data)

	i := 0
	i++
	fmt.Println(i)
}

type Test struct {
	test2 []int
	test4 struct {
		aa string
	}
}

func main() {

	// test()

	// 	log.SetPrefix("greetings: ")
	// 	log.SetFlags(0)
	//
	// 	sentences, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
	// 	// If an error was returned, print it to the console and
	// 	// exit the program.
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// the Hello function to get a message for each name.
	// 	for key, sentence := range sentences {
	// 		fmt.Println(key + ": " + sentence)
	// 	}

	go func() {

		func() {
			fmt.Println("goroutine")
			runtime.Goexit()

		}()
	}()

	fmt.Println("go")

	time.Sleep(time.Millisecond)

}

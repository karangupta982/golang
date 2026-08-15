package main

// import "fmt"


func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) {
	fn(1)
}

func funcReturningFn() func(a int) int{
	return func(a int) int{
		return 2
	}
}

func main(){
	// res := add(3,4)
	// fmt.Println(res)

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int{
		return 2
	}

	processIt(fn)

	fn2 := funcReturningFn()
	fn2(3)

}
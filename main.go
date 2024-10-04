package main

import(
	"fmt"
	"gocache/pkg/gocache"
)


func main() {

	gostore := new(gocache.Store)
	gostore.Make()
	
	
	// setting function to storage
	gostore.Set("testfunc", func(){
		fmt.Println("Hello from storage")
	})
	// calling function from storage
	gostore.Get("testfunc").(func())()

	// getting count of items in store
	fmt.Println(gostore.Count())

	gostore.Set("some_a", "value_a")
	fmt.Println(gostore.Get("some_a"))

}
package main

func main() {
	err := Root.Execute()
	if err != nil {
		panic(err)
	}

}

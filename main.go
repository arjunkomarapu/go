package main

func main() {
	todos := newDeck()
	todos.shuffle()
	todos.print()
	// hand, remainingTodos := deal(todos, 3)
	// hand.print()
	// remainingTodos.print()
	// greeting := "Arjunkomarapu"
	// fmt.Println([]byte(greeting))
	// todos := newDeck()
	// fmt.Println(todos.toString())
	// todos.saveToFile("sample_file")
	// fmt.Println("fileCreated")
	// todos := newDeckFromFile("sample_file")
	// todos.print()
	// fmt.Println(todos)

}

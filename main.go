package main


func main() {
	todos := Todos{}
	storage:=NewStorage[Todos]("todos.json")
	storage.Load(&todos)	
	Commands:=NewCommandFlags()
	Commands.Execute(&todos)
	storage.Save(todos)
	

}
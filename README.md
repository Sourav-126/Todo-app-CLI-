Todo CLI Application in Go

📌 Project Description

This is a simple Todo List application that runs in the command-line interface (CLI), built using Go. It allows users to add, edit, delete, toggle, and list todos with easy-to-use commands.

🚀 Features

Add new todos

Edit existing todos by index

Delete todos by index

Toggle the completion status of todos

List all todos

🛠 Installation

Ensure you have Go installed. Then, clone the repository and build the project:

# Clone the repository
git clone https://github.com/yourusername/todo-cli.git
cd todo-cli



Now, you can run the application using:

go run ./ [command]

📜 Available Commands

➕ Add a new Todo

go run ./ -add "This is a new task"

Example Output:

Todo added: "This is a new task"

📝 Edit an existing Todo

go run ./ -edit 2:"Updated task title"

Example Output:

Todo at index 2 updated: "Updated task title"

❌ Delete a Todo

go run ./ -delete 1

Example Output:

Todo at index 1 deleted.

✅ Toggle completion status

go run ./ -toggle 3

Example Output:

Todo at index 3 marked as completed.

📋 List all Todos

go run ./ -list

Example Output:

1. [ ] Buy groceries
2. [✔] Finish project
3. [ ] Read a book

🏗 How It Works

This application uses Go's flag package to parse command-line arguments and manage a list of todos stored in memory or a file (if implemented). The application follows a simple CRUD pattern to manipulate todos effectively.



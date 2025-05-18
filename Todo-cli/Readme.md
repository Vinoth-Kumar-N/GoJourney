# Todo CLI (GoJourney)

A simple, efficient, and interactive command-line Todo application written in Go. Manage your tasks directly from the terminal with intuitive commands and persistent storage.

---

## Features

- Add new todos with a single command
- List all your todos in a clean table format
- Mark tasks as completed or incomplete
- Edit the titles of your todos
- Delete completed or unwanted todos
- All data is stored locally in a JSON file (`todos.json`)

---

## Getting Started

### Prerequisites

- Go 1.18+ installed on your machine

### Installation

Clone the repository and navigate to the `todo_cli` directory:

```bash
git clone https://github.com/Vinoth-Kumar-N/GoJourney.git
cd GoJourney/todo_cli
```

Install dependencies:

```bash
go mod tidy
```

Build the application:

```bash
go build -o todo
```

---

## Usage

Run the executable with the following command-line flags:

### Add a new todo

```bash
./todo -add "Buy groceries"
```

### List all todos

```bash
./todo -list
```

### Mark a todo as completed/incomplete

```bash
./todo -toggle 0
```

*(Here, `0` is the index of the todo you want to toggle)*

### Edit a todo

```bash
./todo -edit 1:"Read a Go tutorial"
```

*(Here, `1` is the index of the todo you want to edit)*

### Delete a todo

```bash
./todo -del 2
```

*(Here, `2` is the index of the todo you want to delete)*

---

## Command-Line Flags

| Flag         | Description                                            |
|--------------|-------------------------------------------------------|
| `-add`       | Add a new todo (specify the title as argument)        |
| `-del int`   | Delete a todo by its index (default: -1)              |
| `-edit`      | Edit a todo: provide `index:new_title`                |
| `-toggle int`| Mark a todo as completed/incomplete by its index      |
| `-list`      | List all todos                                        |

---

## Example

```bash
# Add todos
./todo -add "Learn Go"
./todo -add "Build a CLI app"

# List todos
./todo -list

# Mark first todo as completed
./todo -toggle 0

# Edit the second todo's title
./todo -edit 1:"Explore Go modules"

# Delete the first todo
./todo -del 0
```

---

## Data Persistence

All todos are stored in a `todos.json` file in the project directory. You can back up or modify this file manually if needed.

---

## Dependencies

- [github.com/aquasecurity/table](https://github.com/aquasecurity/table) (for pretty table output)

---

## Author

- Vinoth Kumar N

---

> Happy tasking! 🚀

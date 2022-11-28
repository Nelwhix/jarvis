# JARVIS
    A CLI tool to record my programming tasks for the 
    day.

## USAGE
    Build the tool with:
```bash
    go build
```
    then to add a new task you write:
```bash
    ./todo -add "Learn how to use Docker and Laravel"
```
    To list out all your tasks you write
```bash
    ./todo -list
```
    To delete a task you write the "-del" flag and the task's number:
```bash
    ./todo -del 1
```
    To complete a task you write the "-finish" flag and the task's number:
```bash
    ./todo -finish 1
```

## Tests
    To test the todo api, from the root directory you run
```bash
    go test -v
```
    To run the integration tests on the cli tool, navigate to the cmd/todo directory and run
```bash
    go test -v
```
    NB: Ensure that no todos.json folder resides in the cmd/todo directory as it could affect tests.
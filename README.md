# To-Do App

This is a simple **CLI To-Do Application** built with **Go**. It allows you to manage your tasks efficiently through a command-line interface. You can **add**, **list**, **complete**, and **delete** tasks with ease.

Under the hood, a **CSV file** is used as a storage of the records.

## Features

- **Add Tasks**: Add new to-do items to your list.
- **List Tasks**: View all uncompleted tasks or optionally list all tasks, including completed ones.
- **Complete Tasks**: Mark tasks as completed.
- **Delete Tasks**: Remove tasks from your list.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/misxka/go-cli-todoapp.git
    ```
2. Navigate to the project directory:
    ```sh
    cd todoapp
    ```
3. Build the app executable:
    ```sh
    go build -o <your_folder>/todoapp
    ```

## Usage

Run the application using the following command:

  ```sh
  <your_folder>/todoapp
  ```

### Commands

- Add a task:
  ```sh
  <your_folder>/todoapp add "Task description"
  ```

- List uncompleted tasks:
  ```sh
  <your_folder>/todoapp list
  ```

  To list all the tasks, including comleted ones:

  ```sh
  <your_folder>/todoapp list --all
  ```

- Mark a task as completed:

  ```sh
  <your_folder>/todoapp complete <task-id>
  ```

- Delete a task:

  ```sh
  ./bin/todoapp delete <task-id>
  ```

## Dependencies

This project uses the following Go libraries:

- [Cobra](https://github.com/spf13/cobra): For building the CLI.
- [timediff](https://github.com/mergestat/timediff): For displaying human-readable time differences.

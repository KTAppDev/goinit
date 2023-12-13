## goinit: A Simple CLI for Initializing Go Projects

**goinit** is a command-line tool that simplifies the process of setting up a new Go project with a basic directory structure and a "Hello World" main file. 

**What it does:**

* Creates a new directory for your project.
* Initializes a `go.mod` file with the provided project name.
* Creates the following directory structure:
    * `cmd`: For your main application files.
    * `api`: For any API-related code.
    * `bin`: For compiled binaries.
    * `lib`: For reusable libraries and packages.

**How to use it:**

```
goinit github.com/USERNAME/simple_app
```

This command will create a new project named `simple_app` in your current directory with the default structure mentioned above.

**Features:**

* **Simple and concise:** Just one command to get your project started.
* **Flexible:** Customizable directory structure and main file content.
* **Fast:** Saves you time and effort compared to manual setup.

**Requirements:**

* Go 1.17+ installed.
* A terminal with access to your command line.

**Installation:**

get it from the releases section

**Contributing:**

We welcome contributions! Feel free to fork the repository and submit pull requests with bug fixes, improvements, or new features.

**License:**

MIT

**Example:**

```
goinit my-amazing-app
```

This will create a directory named `my-amazing-app` with the standard structure and a basic "Hello World" message in the `cmd/main.go` file.


I hope this README provides a clear and concise overview of `goinit`. If you have any questions or suggestions, feel free to create an issue on the GitHub repository.

**Bonus:**

You can also add a table to the README showcasing the available directories and their purposes to further enhance the documentation.

Feel free to customize this template further and make it your own!


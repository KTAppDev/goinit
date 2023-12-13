package main

// created by lazy man Ken Taylor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <project_name>\n", os.Args[0])
		fmt.Println("Not enough arguments. Usage: goinit <project_name>")
		os.Exit(1)
	}

	projectName := os.Args[1]

	err = runCommand("go", "mod", "init", projectName)
	if err != nil {
		fmt.Printf("Error running 'go mod init': %v\n", err)
		os.Exit(1)
	}

	err = createProjectStructure(currentDir, projectName)
	if err != nil {
		fmt.Printf("Error creating project structure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Project '%s' created successfully in %s\n", projectName, currentDir)
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func createProjectStructure(baseDir, projectName string) error {
	directories := []string{
		"cmd",
		"lib",
		"bin",
		"api",
	}

	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(baseDir, dir), os.ModePerm)
		if err != nil {
			return err
		}

		switch dir {
		case "cmd":
			err := createMainFile(filepath.Join(baseDir, dir), projectName)
			if err != nil {
				return err
			}
			// not sure what else i want to add right now
		}
	}

	return nil
}

func createMainFile(dir, projectName string) error {
	filePath := filepath.Join(dir, "main.go")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`
package main

import "fmt"

func main() {
    fmt.Println("Created project, %s")
}
`, projectName)

	_, err = file.WriteString(content)
	return err
}

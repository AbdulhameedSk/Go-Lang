// Project or Workspace Name
package main

// The `main` package is used for executable programs in Go.
// There are two types of packages in Go:

// 1. Executable Packages:
//    - When the package name is `main`, Go generates an executable file.
//    - It must include a `main` function, which serves as the entry point of the program.
//    - The `main` package is meant to be run, not imported by other packages.

// 2. Reusable Packages:
//    - These are library packages meant to be imported and used by other Go programs.
//    - Go’s standard library and third-party packages fall into this category.

// Importing standard or custom packages is done using the `import` statement.
import "fmt" // The `fmt` package provides functions for formatted I/O.

// The `main` function is the entry point when the program is executed.
func main() {
	fmt.Println("Hello, World!")
}

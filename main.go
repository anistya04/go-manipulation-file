package main

import append_file "github.com/anistya04/go-append-file"

func main() {
	append_file.AppendFile("test.txt", "Hello World \n")
}

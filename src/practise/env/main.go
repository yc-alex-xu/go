package main

/*
sample of env in Ubuntu

$ unset port
$ env |grep port
$ set port=80
$ env |grep port
$ export port=80
$ env |grep port
port=80
*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//os.Clearenv()

	os.Setenv("FOO", "1")

	fmt.Println("FOO:", os.Getenv("FOO"))

	os.Unsetenv("FOO")

	fmt.Println("FOO:", os.Getenv("FOO"))

	os.Setenv("hello", "1")

	os.Setenv("world", "2")

	fmt.Println()
	// fetcha all env variables as shown in $ env
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		fmt.Println(variable[0], "=>", variable[1])
	}

}

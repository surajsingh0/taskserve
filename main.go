package main

import "github.com/surajsingh0/taskserve/cmd"

func main() {
	defer cmd.Close()
	cmd.Execute()
}

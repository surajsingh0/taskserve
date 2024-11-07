package main

import "todo/cmd"

func main() {
	defer cmd.Close()
	cmd.Execute()
}

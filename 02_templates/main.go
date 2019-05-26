package main

import "fmt"

func main() {
	name := "Banobe Pascal"

	tpl := `
	<!DOCTYPE html>
 	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <title>Document</title>
	</head>
	<body>
    <h1>` + name + `</h1?
	</body>
	</html>
	`
	fmt.Println(tpl)
}

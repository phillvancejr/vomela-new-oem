package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

const root = "/volumes/phill_backup/"
var (
	company string
	year string
	id string
	number string
	sub_folders = []string{"JPG", "LEGACY", "PRINT", "RESOURCES", "UPLOAD"}
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	print("project company: ")
	company, _ := stdin.ReadString('\n')

	company = strings.Title(company)
	company = strings.Replace(company, " ", ".", -1)
	company = strings.TrimSuffix(company, "\n")

	print("project year: ")
	year, _ := stdin.ReadString('\n')
	year = strings.TrimSuffix(year, "\n")
	if len(year) == 2 {
		year = "20" + year
	}

	print("project id: ")
	id, _ := stdin.ReadString('\n')
	id = strings.TrimSuffix(id, "\n")
	id = fmt.Sprintf("%04s", id)

	print("design number: ")
	number, _ := stdin.ReadString('\n')
	number = strings.TrimSuffix(number, "\n")

	number = strings.Title(number)
	if !strings.HasPrefix(number, "D") {
		number = "D" + number
	}

	project_name := fmt.Sprintf("%s-%s-%s", company, year[len(year)-2:],id)

	for _, dir := range sub_folders {
		os.MkdirAll(root + year + "/" + company + "/" + project_name + "/" + number + "/" + dir, 0777)
	}

	println("finished")
}
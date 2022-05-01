package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Go", "Python", "Java"}
	fmt.Println(courseName)
	courseName = append(courseName, "C++", "C#", "Javascript", "PHP") //สามารถเพิ่มได้หลายตัว
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
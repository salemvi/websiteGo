package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name                     string
	Age                      uint16
	Money                    int16
	AverageGrades, Happiness float64
	Hobbies                  []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is %s He is %d "+
		"and he has money equal %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	vadim := User{ // объект
		Name:          "Vadim",
		Age:           22,
		Money:         -100,
		AverageGrades: 4.1,
		Happiness:     1.1,
		Hobbies:       []string{"JavaScript", "Typescript", "GO", "React"},
	}
	//fmt.Println(w, "User name is: "+vadim.name)
	vadim.setNewName("Salemvi")
	//fmt.Fprintf(w, `"<b>Main text</b> ${}`)
	template, _ := template.ParseFiles("template/home.html")
	template.Execute(w, vadim)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about/", aboutPage)
	http.ListenAndServe(":3008", nil)
}

func main() {
	handleRequest()
}

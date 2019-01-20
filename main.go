package main

import "fmt"

type contactInfo struct {
	Details  string
	Handle   string
	Work     string
	Twitter  string
	Npm      string
	Github   string
	Linkedin string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	nik := person{
		firstName: " Danny \n",
		lastName:  " Vaidyar \n",
		contactInfo: contactInfo{
			Details:  "\n\n",
			Handle:   " Doing stuff in ReactJS|GO|ReactNative|Vue|NodeJS \n",
			Work:     " Software Developer at Yoghurt \n",
			Twitter:  " https://twitter.com/nikhilvaidyar \n",
			Npm:      " https://npmjs.com/~zaynian \n",
			Github:   " https://github.com/nikzayn \n",
			Linkedin: " https://linkedin.com/in/nikhil-vaidyar-2ab005124 \n",
		},
	}

	nik.updateName(" Nikhil \n")
	nik.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

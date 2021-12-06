package commands

import (
	"strings"
)

func Attendance(campaign string) (msg strings.Builder) {
	// Define the slice(array) of bean urls
	transmutation := []string{"Tisi", "Ptrosk", "Baldrick", "Ikol", "Avu", "Red Staché"}
	tropolis := []string{"Adelvir", "Akta", "Ayayron", "Duvu", "Gisli", "Krasus", "Wrench"}

	msg.WriteString("Hey, MyDudes! Who of ")
	switch campaign {
	case "transmutation", "noodles", "noods", "noodles'", "noods'":
		for i, person := range transmutation {
			msg.WriteString(person)
			if i < len(tropolis) - 2 {
				msg.WriteString(", ")
			} else if i == len(tropolis) - 2 {
				msg.WriteString(", and ")
			}
		}
	case "tropolis", "revengers", "mesegea", "jones'", "joneses", "jonesin's":
		for i, person := range tropolis {
			msg.WriteString(person)
			if i < len(tropolis) - 2 {
				msg.WriteString(", ")
			} else if i == len(tropolis) - 2 {
				msg.WriteString(", and ")
			}
		}
	}
	msg.WriteString(" will come to the call this week?")
	return msg
}

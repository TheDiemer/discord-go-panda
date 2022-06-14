package commands

import (
	"strconv"
	"strings"
)

func Attendance(campaign string, author string) (msg strings.Builder, dm strings.Builder) {
	// Define the slice(array) of bean urls
	transmutation := []string{"Tisi", "Ptrosk", "Baldrick", "Ikol", "Avu", "Red Staché"}
	tropolis := []string{"Adelvir", "Akta", "Ayayron", "Duvu", "Gisli", "Krasus", "Wrench"}

	msg.WriteString("Hey, <@&654763072828997642>! Who of ")
	var attendance int
	attendance = 0
	switch campaign {
	case "transmutation", "noodles", "noods", "noodles'", "noods'":
		for i, person := range transmutation {
			msg.WriteString(person)
			if i == len(transmutation)-2 {
				msg.WriteString(", and ")
			} else if i < len(transmutation)-2 {
				msg.WriteString(", ")
			}
		}
		attendance = len(transmutation) + 1
	case "tropolis", "revengers", "mesegea", "jones'", "joneses", "jonesin's":
		for i, person := range tropolis {
			msg.WriteString(person)
			if i == len(tropolis)-2 {
				msg.WriteString(", and ")
			} else if i < len(tropolis)-2 {
				msg.WriteString(", ")
			}
		}
		attendance = len(tropolis) + 1
	default:
		attendance = 0
	}
	tmp := strconv.Itoa(attendance)
	msg.WriteString(" will come to the call this week?")
	dm.WriteString("Hey, ")
	dm.WriteString(author)
	dm.WriteString("! **Thanks for triggering the rollcall!**\nPlease take a moment to create a thread named `attendance` and copy/paste the following chart into the thread between two sets of triple back tics (`):\n\n")
	dm.WriteString("\nYee | ProbablYee | MaybYee | NYee | AFKYee\n 0  |     0      |    0    |   0  |   ")
	dm.WriteString(tmp)

	return msg, dm
}

package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + fmt.Sprint(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return "Welcome to my party, " + name + "!\nYou have been assigned to table " + fmt.Sprintf("%.3d", table) + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\nYou will be sitting next to " + neighbor + "."
}

// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.

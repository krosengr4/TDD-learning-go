package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	roster := Roster{99: "Aaron Judge", 2: "Derek Jeter", 27: "Mike Trout"}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n-----WHAT DO YOU WANT TO DO-----\n1 - Search Player By Number\n2 - Add Player\n3 - Update Player Name\n4 - Delete Player\n0 - Exit")
		fmt.Println("Enter Option:")

		scanner.Scan()
		userOption, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		switch userOption {
		case 1:
			roster.searchPlayer()
		case 2:
			roster.addPlayer()
		case 3:
			roster.updatePlayer()
		case 4:
			roster.deletePlayer()
		case 0:
			fmt.Println("Thanks for visiting!")
			return
		default:
			fmt.Println("ERROR: please enter a number that is listed!")
		}
	}
}

func (r Roster) searchPlayer() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the player's number:")
	scanner.Scan()
	searchNumb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		number, player, err := r.Search(searchNumb)
		if err != nil {
			fmt.Println(strings.Repeat("_", 50))
			fmt.Println(err)
			fmt.Println(strings.Repeat("_", 50))
		} else {
			fmt.Println(strings.Repeat("_", 50))
			fmt.Println("Number:", number, "Player:", player)
			fmt.Println(strings.Repeat("_", 50))
		}
	}
}

func (r Roster) addPlayer() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the player number:")
	scanner.Scan()
	addNumb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("Enter the player name:")
	scanner.Scan()
	addName := scanner.Text()

	err = r.Add(addNumb, addName)
	if err != nil {
		fmt.Println(strings.Repeat("_", 50))
		fmt.Println(err)
		fmt.Println(strings.Repeat("_", 50))
	} else {

	}
}

func (r Roster) updatePlayer() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the player number:")
	scanner.Scan()
	updateNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter the updated name:")
	scanner.Scan()
	newName := scanner.Text()

	updatedName, err := r.Update(updateNum, newName)
	if err != nil {
		fmt.Println(strings.Repeat("_", 50))
		fmt.Println(err)
		fmt.Println(strings.Repeat("_", 50))
	} else {
		fmt.Println("Successfully updated player")
		fmt.Println("Number:", updateNum, "New name:", updatedName)
	}
}

func (r Roster) deletePlayer() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Player Number:")
	scanner.Scan()
	deleteNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	err = r.Delete(deleteNum)
	if err != nil {
		fmt.Println(strings.Repeat("_", 50))
		fmt.Println(err)
		fmt.Println(strings.Repeat("_", 50))
	} else {
		fmt.Println("Successfully deleted player")
	}
}

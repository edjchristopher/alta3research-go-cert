/* Alta3 Research Certification
   reading a JSON file asking a question, taking input, formatting
   it with a struct and displayoing the information requested - EdC */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of websites links
type User struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Age      int    `json:"Age"`
	Websites `json:"websites"`
}

// Favorite struct which contains a
// list of links
type Websites struct {
	LinkedIn string `json:"LinkedIn"`
	Favorite string `json:"Favorite"`
}

func main() {
loop:
	a := "author"
	g := "gamer"
	// ################################
	// declaring the variable using the var keyword
	var fromUser string
	fmt.Println("")
	fmt.Println("To end this program type in CTRL-C. To see a list of Authors type in author, to see a list of Gamers type in gamer:\n")
	// scanning the input by the user
	fmt.Scanln(&fromUser)
	fromUser = strings.ToLower(fromUser)
	if strings.Compare(fromUser, a) == 0 {
		// fmt.Println("\nSorry try typing that again!\n")
		goto finish
	} else if strings.Compare(fromUser, g) == 0 {
		// fmt.Println("\nSorry try typing that again!\n")
		goto finish
	} else {
		fmt.Println("\nSorry try typing that again!\n")
		goto loop
	}

	// ################################

finish:
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(byteValue, &users); err != nil {
		log.Println(err)
	}

	// iterate through each user within our users array and
	// print out only the user Type, of the fromUser that was typed in
	for i := 0; i < len(users.Users); i++ {
		// ########################################
		fromUser = strings.ToLower(fromUser)
		author := strings.ToLower(users.Users[i].Type)
		if fromUser == author {
			fmt.Println("User Type: " + users.Users[i].Type)
			fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
			fmt.Println("User Name: " + users.Users[i].Name)
			fmt.Println("LinkedIn Url: " + users.Users[i].Websites.LinkedIn)
			fmt.Println("Favorite Url: " + users.Users[i].Websites.Favorite)
			fmt.Println("\n")
		}
	}
	goto loop
}

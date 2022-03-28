package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Provide data type explicitly, package level variables can be accessed in any function within this package
// Best Practic : Define variable as "local" as possible
const conferenceTickets int = 50

var conferenceName string = "Go Conference"

var remainingTickets uint = 50

// Slice, It is dynamic in size, It grows as we add elements, we have declared slice of string here
//var bookings []string

// We have declared slice of map here, map is also string key and string value type
// We have provided the initial size as 0 (required when we initialize the slice) and slice is dynamic so it will keep
// increasing as we add more elements
// var bookings = make([]map[string]string, 0)

// We have declared slice of UserData here
var bookings = make([]UserData, 0)

//bookings = []string{} // We cannot use shortlang for package level variables

// Structure can have mixed datatypes, you can relate it to classes in Java
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait group is used to make goroutine synchronized so that main thread will wait for other threads to complete before
// terminating itself
var wg = sync.WaitGroup{}

func main() {
	//testCode()
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userTickets, email, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// call bookTicket
			bookTicket(userTickets, email, firstName, lastName)

			// "go..." - starts a new goroutine
			// A goroutine is a lightweight thread managed by the Go runtime
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// print all bookings
			// fmt.Printf("These are all our bookings: %v\n", bookings)

			// call getFirstNames function
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

	// call swtichCodeExample for switch statement code
	//switchCodeExample()
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{} //shorthand to declare a slice

	// for index, booking := range bookings { // below line works same as this line and also ignore index
	// for _, booking := range bookings { // _ is used to ignore the variable , we are iterating the list of strings
	// 	var names = strings.Fields(booking) // split string with space
	// 	firstNames = append(firstNames, names[0])
	// }

	// for _, booking := range bookings { // _ is used to ignore the variable, we are iterating the list of maps
	// 	firstNames = append(firstNames, booking["firstName"])
	// }

	for _, booking := range bookings { // _ is used to ignore the variable, we are iterating the list of struct UserData here
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, email string, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	//create a map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // convert the unint value to string

	// create UserData
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Working with slice of string, below line add element to slice
	// bookings = append(bookings, firstName+" "+lastName)

	// Working with slice of map, below line add map to slice
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending %v to email address %v\n", ticket, email)
	fmt.Println("#######################")
	wg.Done()
}

func switchCodeExample() {
	//Switch statement
	city := "London"
	switch city {
	case "New York":
		// execute code for booking New York conference tickets
	case "Singapore", "Hong Kong":
		// execute code for booking Singapore & Hong Kong conference tickets
	case "London", "Berlin":
		// execute code for booking London & Berlin conference tickets
	case "Maxico City":
		// some code here
	default:
		fmt.Println("No valid city selected")
	}
}

func testCode() {
	// fmt.Print("Hello World")

	// Go will automatically identify the data type on bases of values if you have not specified
	// var conferenceName = "Go Conference"
	// conferenceName := "Go Conference" //It also works same as above line
	// const conferenceTickets = 50
	// var remainingTickets = 50

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	// Array with assignment, array has fixed size which you provide
	// var bookings = [50]string{"LK", "Nana", "Mike"}
	// fmt.Printf("bookings =>%v", bookings)

	// Array example
	// var bookings [50]string
	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	// It will print the pointer address
	// fmt.Println(&conferenceName)

}

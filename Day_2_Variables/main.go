package main
import "fmt"

const Login_token string = "hdvbcnh" // public

// variable without var can't br declare outside of the main function
// number_of_users := 30000.0
// var number_of_users int = 30000.0
// fmt.Println(number_of_users)

func main() {
	var username string = "Kebba"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var is_logged_in bool = true
	fmt.Println(is_logged_in)
	fmt.Printf("Variable is of type: %T \n", is_logged_in)

	var small_val uint8 = 255
	fmt.Println(small_val)
	fmt.Printf("Variable is of type: %T \n", small_val)

	var small_float float64 = 255.45544567777
	fmt.Println(small_float)
	fmt.Printf("Variable is of type: %T \n", small_float)

	default values and some aliases
	var another_variable int
	fmt.Println(another_variable)
	fmt.Printf("Variable is of type: %T \n", another_variable)

	// implicit value
	var website = "disruptiveX.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style to declare variables
	number_of_users := 30000.0
	fmt.Println(number_of_users)

	fmt.Println("Login_token")
	fmt.Printf("Variable is of type: %T \n", Login_token)

}


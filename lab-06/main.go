package main

type User struct {
	Email string

	FirstName string
	LastName  string

	AddressFirstLine  string
	AddressSecondLine string
	City              string
	Postcode          string
}

func main() {

	// using the var keyword to create the variable
	var myUser User

	// we can then use dot notation to set fields within the variable or print them out
	myUser.Email = "foo@bar.com"
	println(myUser.Email) // foo@bar.com

	// we can also use the := notation for creation and assignment like other variables
	myUserTwo := User{
		Email:     "john.doe@bar.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	// We can interact with the fields inside by using var.field notation
	println(myUserTwo.Email) // john.doe@bar.com

}

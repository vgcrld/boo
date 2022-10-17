package common

type User struct {
	FirstName string
	LastName  string
	Email     string
	Tickets   uint16
}

// Create a new user
func NewUser(firstName, lastName, email string, tickets uint16) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Tickets:   tickets,
	}
}

type Event struct {
	Class    string
	Location string
	Students []User
}

// Create a new event
func NewEvent(class, location string) Event {
	return Event{
		Class:    class,
		Location: location,
	}
}

package deck

import "math/rand"

type Dealer struct{ FirstName, LastName string }

func generateDealer() Dealer {
	dealer := Dealer{FirstName: "", LastName: ""}
	randomIndex := rand.Intn(len(FIRST_NAMES))
	dealer.FirstName = FIRST_NAMES[randomIndex]
	randomIndex = rand.Intn(len(LAST_NAMES))
	dealer.LastName = LAST_NAMES[randomIndex]
	return dealer
}

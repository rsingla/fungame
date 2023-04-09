package main

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Customer struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	TransactionId string `json:"transaction_id"`
}

package main

import "github.com/brianvoe/gofakeit/v6"

func randomContractData() string {

	faker := gofakeit.New(0)

	contractData := faker.Sentence(100)

	return contractData

}

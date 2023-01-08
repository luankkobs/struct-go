package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

func (c Client) Display() {
	fmt.Println("Displaying Client by method:", c.Name)
}

type InternationalClient struct {
	Client
	Country string `json:"country"`
}

func main() {
	client := Client{
		Name:  "Luan",
		Email: "luan@gmail.com",
		CPF:   11079245960,
	}
	fmt.Println(client)

	client2 := Client{"João", "joão@gmail.com", 751760969653}
	fmt.Printf("Name: %s. Email: %s CPF: %d\n", client2.Name, client2.Email, client2.CPF)

	client3 := InternationalClient{
		Client: Client{
			Name:  "Robert",
			Email: "robert@gmail.com",
			CPF:   10315890908,
		},
		Country: "Korea",
	}
	fmt.Printf("Name: %s. Email: %s CPF: %d Country: %s\n", client3.Name, client3.Email, client3.CPF, client3.Country)
	client.Display()
	client2.Display()
	client3.Display()

	clientJson, err := json.Marshal(client3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clientJson))

	jsonClient4 := `{"name":"Robert","email":"robert@gmail.com","cpf":10315890908,"country":"Korea"}`
	client4 := InternationalClient{}
	json.Unmarshal([]byte(jsonClient4), &client4)
	fmt.Println(client4)
}

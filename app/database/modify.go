package database

import (
	"fmt"
	"log"
	"time"
)

func Init() {
	db := GetCredentials()
	defer db.Close()

	err := create_schema(db)
	if err != nil {
		fmt.Println(err)
	}
}

func AddMessage(author int64, content string, ip_addr string) {
	db := GetCredentials()
	defer db.Close()
	message := &Message{Content: content, Author: author, IpAddr: ip_addr, DateTime: time.Now()}
	fmt.Printf("%s", message)
	_, err := db.Model(message).Insert()
	if err != nil {
		log.Fatal(err)
	}
}
func GetMessage(id int64) (Message, error) {
	db := GetCredentials()
	defer db.Close()

	var message Message
	err := db.Model(&message).Where("message.id = ?", id).Select()
	return message, err
}

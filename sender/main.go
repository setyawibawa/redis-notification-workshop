package main

import (
	"bufio"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("====== Sender Window ======")
	username := prompt(reader, "username: ")

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	defer func(c redis.Conn) {
		_ = c.Close()
	}(c)

	go pingOnline(username, c)

	for {
		message := prompt(reader, "-> ")

		err := sendMessage(username, message, c)
		if err != nil {
			panic(err)
		}

	}
}

func pingOnline(username string, c redis.Conn) {
	for {
		_, err := c.Do("SETEX", fmt.Sprintf("online:%s", username), "10", "true")
		if err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}

func sendMessage(senderUsername string, message string, c redis.Conn) error {
	_, err := c.Do("SETEX", fmt.Sprintf("message:%s", senderUsername), "60", message)
	if err != nil {
		return err
	}
	return nil
}

func prompt(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	username := readString(reader)
	return username
}

func readString(reader *bufio.Reader) string {
	username, _ := reader.ReadString('\n')
	// convert CRLF to LF
	username = strings.Replace(username, "\n", "", -1)
	return username
}

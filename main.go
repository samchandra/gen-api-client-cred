package main

import (
  "flag"
  "fmt"
  "time"
  "crypto/md5"
  "encoding/hex"

  "github.com/satori/go.uuid"
)

func main() {

	clientName := flag.String("n", fmt.Sprintf("Client#%d", time.Now().Unix()), "Client name unique string")
	flag.Parse()

	fmt.Println("Creating API Client Credentials\n")
	fmt.Printf("Client Name: %s\n\n", *clientName)

	DisplayCred("DEV", *clientName)
	DisplayCred("PRD", *clientName)

}

func DisplayCred(env string, clientName string) {
	salt := "my-awesome-go-client-generator"

	fmt.Println(env)
	fmt.Println("===")

	md5:= GetMD5Hash(fmt.Sprintf("%s-%s-%s", clientName, env, salt))
	fmt.Printf("Client ID: %s\n", md5)

	// Using V3 with namespace and string
	// const nsString = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
        // ns, err := uuid.FromString(nsString)
	// if err != nil {
        //   fmt.Printf("Something went wrong: %s", err)
	// }
	
	// u1 := uuid.NewV3(ns, clientName)
	// fmt.Printf("Client Secret: %s\n", u1)

	// Using V4 with random
	u1 := uuid.NewV4()
	fmt.Printf("Client Secret: %s\n\n", u1)
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

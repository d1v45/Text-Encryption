//#####################################################################################################
// Title: TEXT ENCRYPTION                                                                             #
// Author: d1v45                                                                                      #
// Github : https://github.com/d1v45                                                                  #
//#####################################################################################################


package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating keys: %v", err)
	}
	publicKey = &privateKey.PublicKey
}

func banner() {
	fmt.Println(`
"    █▀▀█    ┏━━┓░░░░░░░░░┏┓┏┓░┏┳┓      ░░
"    █░      ┃━━┫┏━┓┏┳┓┏┳┓┣┫┃┗┓┃┃┃    █▀▀█
"    █▄▄█    ┣━━┃┃━┫┃┃┃┃┏┛┃┃┃┏┫┣┓┃      ░█
"    ░░      ┗━━┛┗━┛┗━┛┗┛░┗┛┗━┛┗━┛    █▄▄█
"    				-d1v45 
`)
}

func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func DecodeBase64(base64Str string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	base64EncodedMessage := EncodeBase64(requestData.Message)

	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(base64EncodedMessage))
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedMessage)

	json.NewEncoder(w).Encode(map[string]string{"encrypted_message": encryptedBase64})
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		EncryptedMessage string `json:"encrypted_message"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	encryptedMessage, err := base64.StdEncoding.DecodeString(requestData.EncryptedMessage)
	if err != nil {
		http.Error(w, "Invalid base64 encoded message", http.StatusBadRequest)
		return
	}

	decryptedBase64Message, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessage)
	if err != nil {
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}

	finalMessage, err := DecodeBase64(string(decryptedBase64Message))
	if err != nil {
		http.Error(w, "Base64 decode failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"decrypted_message": finalMessage})
}

func main() {
	banner()
	http.HandleFunc("/encrypt", EncryptHandler)
	http.HandleFunc("/decrypt", DecryptHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

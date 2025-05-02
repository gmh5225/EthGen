package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Get the number of wallets to generate from command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: ethgen <number_of_wallets>")
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 {
		fmt.Println("Please provide a valid positive number")
		os.Exit(1)
	}

	// Create the directory to store wallet files
	outputDir := "wallets"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", outputDir, err)
		os.Exit(1)
	}

	fmt.Printf("Generating %d Ethereum wallets and saving to '%s' directory...\n", count, outputDir)

	// Generate wallets and save them to files
	for i := 0; i < count; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			fmt.Printf("Error generating key %d: %v\n", i+1, err)
			continue
		}

		privateKeyBytes := crypto.FromECDSA(privateKey)
		address := crypto.PubkeyToAddress(privateKey.PublicKey)

		// Construct the filename
		fileName := fmt.Sprintf("wallet_%d.txt", i+1)
		filePath := filepath.Join(outputDir, fileName)

		// Create and open the file
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filePath, err)
			continue
		}
		defer file.Close()

		// Format and write the content
		content := fmt.Sprintf("Address: %s\nPrivate Key: %x\n", address.Hex(), privateKeyBytes)
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
		}

		fmt.Printf("Saved wallet %d to %s\n", i+1, filePath)
	}

	fmt.Printf("Wallet generation complete. Files saved in '%s' directory.\n", outputDir)
}

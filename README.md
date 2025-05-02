# EthGen

A simple command-line tool written in Go to generate Ethereum wallet addresses and private keys in bulk.

## Features

*   Generates a specified number of Ethereum wallets.
*   Saves each wallet's address and private key to a separate `.txt` file.
*   Organizes generated wallet files into a `wallets` directory.

## Usage

Run the compiled executable from your terminal, providing the number of wallets you want to generate as a command-line argument:

```bash
./EthGen <number_of_wallets>
```

**Example:**

To generate 5 wallets:

```bash
./EthGen 5
```

This command will:

1.  Create a `wallets` directory if it doesn't exist.
2.  Generate 5 new Ethereum wallets.
3.  Save each wallet's details into files named `wallets/wallet_1.txt`, `wallets/wallet_2.txt`, ..., `wallets/wallet_5.txt`.
4.  Print the progress to the console.

Each `.txt` file will contain:

```
Address: <0x... Ethereum Address>
Private Key: <your private key in hex>
```

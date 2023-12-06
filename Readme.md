# Ethereum Mnemonic Permutation Checker

This Go program is designed to find a specific Ethereum address by permuting the words of a given mnemonic. It utilizes goroutines for concurrent processing, making the search process faster and more efficient.

## Overview

The program takes a mnemonic phrase (a list of words) and a desired Ethereum address as input. It then generates all possible permutations of the mnemonic and checks each one to see if it derives the specified Ethereum address. The program uses the `go-ethereum-hdwallet` library to derive addresses from mnemonics.

## Features

- **Permutation Generation**: Generates all possible permutations of the provided mnemonic.
- **Concurrent Processing**: Utilizes Go's goroutines to check multiple permutations in parallel.
- **Ethereum Address Derivation**: Uses the HD Wallet derivation standard to generate Ethereum addresses from mnemonics.

## Requirements

- Go (version 1.13 or later recommended)
- `go-ethereum-hdwallet` library

## Installation

Before running the program, ensure you have Go installed on your system. You can then install the required Go package using:

```bash
go get -u github.com/miguelmota/go-ethereum-hdwallet
```

## Usage

1. Clone the repository or copy the source code into a new .go file.

2. Modify the mnemonic and desiredAddress variables in the main function to your desired mnemonic and Ethereum address.

3. Run the program using:

```bash
go run <filename>.go
```

## How It Works

The program starts by sorting the words in the mnemonic to ensure permutations are generated in a systematic order.
It then enters a loop, generating the next permutation of the mnemonic words on each iteration.
For each generated mnemonic, a new goroutine is spawned to check if the derived Ethereum address matches the desired address.
If a match is found, the program prints the successful mnemonic and address, then exits.

## Limitations

The number of permutations grows factorially with the number of words in the mnemonic, which can make the process time-consuming for long mnemonics.
This tool should be used responsibly and ethically, considering the security implications of handling mnemonics.

## Contributing

Contributions to the project are welcome. Please ensure that your code adheres to Go best practices and includes appropriate tests.

## License

This project is open-source and available under **Apache License 2.0**.

## Disclaimer

This tool is provided "as is", without warranty of any kind. Use it at your own risk.
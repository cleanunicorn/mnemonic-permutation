package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func nextPermutation(words []string) bool {
    // Find non-increasing suffix
    i := len(words) - 1
    for i > 0 && words[i-1] >= words[i] {
        i--
    }
    if i <= 0 {
        return false
    }

    // Find successor to pivot
    j := len(words) - 1
    for words[j] <= words[i-1] {
        j--
    }

    words[i-1], words[j] = words[j], words[i-1]

    // Reverse suffix
    j = len(words) - 1
    for i < j {
        words[i], words[j] = words[j], words[i]
        i++
        j--
    }

    return true
}

func checkMnemonic(mnemonic, desiredAddress string, wg *sync.WaitGroup, found *bool) {
	defer wg.Done()

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return
	}

	derivedAddress := account.Address.Hex()

	if strings.EqualFold(derivedAddress, desiredAddress) {
		fmt.Println("Found it!")
		fmt.Println(derivedAddress)
		fmt.Println(mnemonic)
		*found = true
	}
}

func main() {
	mnemonic := "banana alien boat bone cat cloud dog dolphin hospital kiwi lion pizza"
	desiredAddress := "0x67d37A4E1674e4CBf13d66B839010B31f63dd844"

	fmt.Println(mnemonic, desiredAddress)

	// Split mnemonic into words
	words := strings.Fields(mnemonic)
	sort.Strings(words)

	var wg sync.WaitGroup
	found := false
	permutationCount := 0	

	wg.Add(1)
	currentMnemonic := strings.Join(words, " ")
	fmt.Println("Trying ", currentMnemonic)
	go checkMnemonic(currentMnemonic, desiredAddress, &wg, &found)

	for !found && nextPermutation(words) {
		permutationCount++
		if permutationCount%100000 == 0 {
			fmt.Printf("Tried %d permutations\n", permutationCount)
		}

		currentMnemonic := strings.Join(words, " ")

		wg.Add(1)
		go checkMnemonic(currentMnemonic, desiredAddress, &wg, &found)

		if found {
			break
		}
	}

	wg.Wait()
}

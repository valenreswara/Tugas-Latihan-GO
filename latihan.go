package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Pilih program yang akan dijalankan: (Tekan 'q' untuk keluar)")
	fmt.Println("1. Print Universe")
	fmt.Println("2. Reverse Word")

	var input string
	fmt.Scanln(&input)

	if input == "q" {
		fmt.Println("Keluar dari program.")
		return
	}

	switch input {
	case "1":
		fmt.Println("Masukkan angka:")
		var number int
		fmt.Scanln(&number)
		if number == 42 {
			fmt.Println("Hello Universe!")
		} else {
			fmt.Println(number)
		}
	case "2":
		fmt.Println("Berikan input kalimat!: (minimal 3 kata)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var kalimat string = scanner.Text()
		terpisah := strings.Split(kalimat, " ")
		if len(terpisah) < 3 {
			fmt.Println("Kalimat minimal harus 3 kata!")
			return
		}
		reversed := make([]string, len(terpisah))
		for i, kata := range terpisah {
			reversed[len(terpisah)-1-i] = kata
		}
		fmt.Printf("Hasil reverse word: ")
		fmt.Println(reversed)
	default:
		fmt.Println("Pilihan tidak ada")
	}
}

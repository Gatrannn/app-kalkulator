package app_kalkulator

import "fmt"

func Kalkulator(){
	var pertama int
	var operator string
	var kedua int
	var tanya string

	for{
		fmt.Print("Apakah anda ingin melanjutkan ke kalkulator (y/n) :")
		fmt.Scan(&tanya)

		if tanya == "n"{
			break
		}


		fmt.Print("Masukkan angka pertama :")
		fmt.Scan(&pertama)
		hasil := pertama


		for{
			fmt.Print("Masukkan operator (+,-,*,/) dan (c) untuk reset :")
			fmt.Scan(&operator)

			if operator == "c"{
				break
			}

			

			fmt.Print("Masukkan angka kedua :")
			fmt.Scan(&kedua)

			switch operator {
			case "+":
				hasil += kedua
				fmt.Println(hasil)
			case "-":
				hasil -= kedua
				fmt.Println(hasil)
			case "*":
				hasil *= kedua
				fmt.Println(hasil)
			case "/":
				if kedua == 0{
					fmt.Println("Maaf tidak dapat melakukan pembagiand dengan 0")
					continue
				}
				hasil /=kedua
				fmt.Println(hasil)
		
			}


		}
	}
	

}

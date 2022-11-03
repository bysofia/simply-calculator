package main

import "fmt"

func main() {
	var sisi, alas, tinggi, jari, depan, samping, miring float64
	var s, t, v0, a, m, h, v, f, n, v1, F, A float64
	var celcius, reamur, fahrenheit, kelvin float64
	var pilihan1, pilihan2 string

	fmt.Println("dibuat oleh Rizky Abdulah")
	fmt.Println("04/11/2022")
	fmt.Println("Pilih Program yang akan digunakan dengan memasukan input huruf program")
	fmt.Println("a. Menghitung Luas Persegi")
	fmt.Println("b. Menghitung Luas Segitiga")
	fmt.Println("c. Menghitung Luas Lingkaran")
	fmt.Println("d. Menghitung Sudut sinus, cosinus, tangen")
	fmt.Println("e. Menghitung gerak lurus beraturan")
	fmt.Println("f. Menghitung gerak lurus berubah beraturan")
	fmt.Println("g. Menghitung energi potensial, kinetik")
	fmt.Println("h. Menghitung frekuensi atau getaran")
	fmt.Println("i. Menghitung masa jenis")
	fmt.Println("j. Menghitung daya, tekanan, usaha dan gaya")
	fmt.Println("k. Konversi untuk semua satuan suhu")
	fmt.Scanln(&pilihan1)

	switch pilihan1 {
	case "a":
		fmt.Println("Masukan nilai sisi panjang Persegi")
		fmt.Scanln(&sisi)
		fmt.Printf("%.2f*%.2f = %.2f", sisi, sisi, sisi*sisi)

	case "b":
		fmt.Println("Masukan nilai alas")
		fmt.Scanln(&alas)
		fmt.Println("Masukan nilai Tinggi")
		fmt.Scanln(&tinggi)
		fmt.Printf("%.2f * %.2f = %.2f", alas, tinggi, alas*tinggi)

	case "c":
		fmt.Println("Masukan nilai panjang jari jari Lingkaran")
		fmt.Scanln(&jari)
		fmt.Printf("phi * %.2f * %.2f = %.2f", jari, jari, 3.14*jari*jari)

	case "d":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukan input huruf")
		fmt.Println("a. Nilai Sinus")
		fmt.Println("b. Nilai Cosinus")
		fmt.Println("c. Nilai Tangen")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "a":
			fmt.Println("Masukan nilai panjang sisi depan")
			fmt.Scanln(&depan)
			fmt.Println("Masukan nilai panjang sisi miring")
			fmt.Scanln(&miring)
			fmt.Printf("Sin %.2f/%.2f = %.2f", depan, miring, depan/miring)

		case "b":
			fmt.Println("Masukan nilai panjang sisi samping")
			fmt.Scanln(&samping)
			fmt.Println("Masukan nilai panjang sisi miring")
			fmt.Scanln(&miring)
			fmt.Printf("Cos %.2f/%.2f = %.2f", samping, miring, samping/miring)

		case "c":
			fmt.Println("Masukan nilai panjang sisi depan")
			fmt.Scanln(&depan)
			fmt.Println("Masukan nilai panjang sisi samping")
			fmt.Scanln(&samping)
			fmt.Printf("Tan %.2f/%.2f = %.2f", depan, samping, depan/samping)

		}

	case "e":
		fmt.Println("Masukan nilai jarak")
		fmt.Scanln(&s)
		fmt.Println("Masukan nilai waktu")
		fmt.Scanln(&t)
		fmt.Printf("V = %.2f/%.2f = %.2f", s, t, s/t)

	case "f":
		fmt.Println("Masukan nilai kecepatan awal")
		fmt.Scanln(&v0)
		fmt.Println("Masukan nilai percepatan")
		fmt.Scanln(&a)
		fmt.Println("Masukan nilai waktu")
		fmt.Scanln(&t)
		fmt.Printf("Vt =  %.2f +  %.2f*/%.2f = %.2f", v0, a, t, v0+(a*t))

	case "g":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukan input huruf")
		fmt.Println("a. Energi Potensial")
		fmt.Println("b. Energi Kinetik")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "a":
			fmt.Println("Masukan nilai Massa")
			fmt.Scanln(&m)
			fmt.Println("MAsukan nilai Ketinggian")
			fmt.Scanln(&h)
			fmt.Printf("Ep = m * g * h = %.2f", m*10*v)

		case "b":

			fmt.Println("Masukan nilai Masa")
			fmt.Scanln(&m)
			fmt.Println("Masukan nilai Kecepatan")
			fmt.Scanln(&v)
			fmt.Printf("Et = 1/2 * m * v^2 = %.2f", 0.5*m*v*v)

		}

	case "h":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukan input huruf")
		fmt.Println("a. Frekuensi")
		fmt.Println("b. Getaran")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "a":
			fmt.Println("Masukan nilai getaran")
			fmt.Scanln(&n)
			fmt.Println("Masukan nilai waktu")
			fmt.Scanln(&t)
			fmt.Printf("f = %.2f/%.2f = %.2f", n, t, n/t)

		case "b":
			fmt.Println("Masukan nilai frekuensi")
			fmt.Scanln(&f)
			fmt.Println("Masukan nilai waktu")
			fmt.Scanln(&t)
			fmt.Printf("f = %.2f * %.2f = %.2f", f, t, f*t)
		}

	case "i":
		fmt.Println("Masukan nilai Masa")
		fmt.Scanln(&m)
		fmt.Println("Masukan nilai Volume")
		fmt.Scanln(&v1)
		fmt.Printf("rho = %.2f/%.2f = %2f", m, v1, m/v1)

	case "j":
		fmt.Println("Pilih nilai yang ingin dicari dengan memasukan input huruf")
		fmt.Println("a. Daya")
		fmt.Println("b. Tekanan")
		fmt.Println("c. Usaha")
		fmt.Println("d. Gaya")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "a":
			fmt.Println("Masukan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukan nilai kecepatan")
			fmt.Scanln(&v)
			fmt.Printf("P = %.2f/%.2f = %.2f", F, v, F/v)

		case "b":
			fmt.Println("Masukan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukan nilai Luas Permukaan (m^2)")
			fmt.Scanln(&A)
			fmt.Printf("p = %.2f * %.2f = %.2f", F, A, F/A)

		case "c":
			fmt.Println("Masukan nilai gaya")
			fmt.Scanln(&F)
			fmt.Println("Masukan nilai perpindahan")
			fmt.Scanln(&s)
			fmt.Printf("W = %.2f * %.2f = %.2f", F, s, F*s)

		case "d":
			fmt.Println("Masukan nilai masa")
			fmt.Scanln(&m)
			fmt.Println("Masukan nilai kecepatan")
			fmt.Scanln(&a)
			fmt.Printf("F = %.2f * %.2f = %.2f", m, a, m*a)

		}

	case "k":
		fmt.Println("Pilih nilai suhu yang diketahui")
		fmt.Println("a. Celsius")
		fmt.Println("b. Reamur")
		fmt.Println("c. Fahrenheit")
		fmt.Println("d. Kelvin")
		fmt.Scanln(&pilihan2)

		switch pilihan2 {
		case "a":
			fmt.Println("Pilih nilai suhuh dalam °C")
			fmt.Scanln(&celcius)
			fmt.Printf("%.2f °R\n", celcius*4/5)
			fmt.Printf("%.2f °F\n", (celcius*9/5)+32)
			fmt.Printf("%.2f °K\n", celcius+273)

		case "b":
			fmt.Println("Pilih nilai suhuh dalam °R")
			fmt.Scanln(&reamur)
			fmt.Printf("%.2f °C\n", reamur*5/4)
			fmt.Printf("%.2f °F\n", (reamur*9/4)+32)
			fmt.Printf("%.2f °K\n", reamur+273)

		case "c":
			fmt.Println("Pilih nilai suhuh dalam °F")
			fmt.Scanln(&fahrenheit)
			fmt.Printf("%.2f °C\n", 5/9*(fahrenheit-32))
			fmt.Printf("%.2f °R\n", 4/9*(fahrenheit+32))
			fmt.Printf("%.2f °K\n", (5/9*(fahrenheit-32))+273)

		case "d":
			fmt.Println("Pilih nilai suhuh dalam °K")
			fmt.Scanln(&kelvin)
			fmt.Printf("%.2f °C\n", kelvin-273)
			fmt.Printf("%.2f °R\n", (kelvin-273)*4/5)
			fmt.Printf("%.2f °F\n", ((kelvin-273)*9/5)+32)

		}
	}

}

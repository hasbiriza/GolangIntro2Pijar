// Ini pake function biasa , tanpa struct --------------------------------------------------------//
// package main

// import (
// 	"fmt"
// )

// // type mahasiswa struct {
// // 	width  float64
// // 	height float64
// // }

// // func calculateArea(width float64, height float64) float64 {
// // 	return width * height
// // }

// // func main() {
// // 	cek := mahasiswa{
// // 		width:  10,
// // 		height: 20,
// // 	}

// // 	fmt.Println(calculateArea(cek.width, cek.height))
// // }

//Ini Dibawah Menggunakan Metode Struct Methode---------------------------------------------------------//

package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func (cek Rectangle) kalkulasiArea() float64 {
	return cek.width * cek.height

}
func cekFloat(value float64) bool {
	// Validasi Float
	return true
}

func main() {
	coba := Rectangle{
		width:  12123,
		height: 20,
	}

	// Verifikasi bahwa nilai yang dimasukkan adalah float64
	if cekFloat(coba.width) && cekFloat(coba.height) {
		fmt.Println(coba.kalkulasiArea())
	} else {
		fmt.Println("Nilai yang dimasukkan bukan float64.")
	}
}

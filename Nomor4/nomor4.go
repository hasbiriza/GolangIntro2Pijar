package main

import (
	"fmt"
	"sort"
)

type SortData struct {
	Index  int
	Number int
}

func SeleksiNilai(nilaiAwal, nilaiAkhir int, dataArray []int) {
	if nilaiAwal >= nilaiAkhir {
		fmt.Println("Nilai akhir harus lebih besar dari nilai awal")
		return
	}

	if len(dataArray) <= 5 {
		fmt.Println("Jumlah angka dalam dataArray harus lebih dari 5")
		return
	}

	var foundData []SortData

	for i, num := range dataArray {
		if num >= nilaiAwal && num <= nilaiAkhir {
			foundData = append(foundData, SortData{i, num})
		}
	}

	if len(foundData) == 0 {
		fmt.Println("Nilai tidak ditemukan")
		return
	}

	sort.Slice(foundData, func(i, j int) bool {
		return foundData[i].Number < foundData[j].Number
	})

	for _, data := range foundData {
		fmt.Println(data.Number)
	}
}
func main() {
	//berhasil
	numbers := []int{2, 25, 4, 14, 17, 30, 8}
	SeleksiNilai(5, 20, numbers)

	//Jumlah angka dalam array harus lebih dari 5
	lebihBesar := []int{1, 1, 3}
	SeleksiNilai(5, 20, lebihBesar)

	//nilai akhir harus lebih besar dari nilai awal
	nilaiAkhir := []int{1, 1, 3, 2, 4}
	SeleksiNilai(20, 5, nilaiAkhir)

	//Nilai Tidak Ditemukan
	nilaiTidakDitemukan := []int{1, 1, 3, 2, 4, 1}
	SeleksiNilai(5, 15, nilaiTidakDitemukan)
}

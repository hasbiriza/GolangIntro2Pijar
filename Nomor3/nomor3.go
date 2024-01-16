package main

import (
	"fmt"
	"sort"
	"strings"
)

// Fungsi untuk mencari nama dengan memperhitungkan huruf besar/kecil
func searchName(daftarNama []string, searchName string, maxOutput int, callback func([]string)) {
	var found []string

	// Mengubah searchName menjadi lowercase agar pencarian tidak bersifat case-sensitive
	searchName = strings.ToLower(searchName)

	for _, name := range daftarNama {
		// Mengubah nama menjadi lowercase agar pembandingan tidak bersifat case-sensitive
		lowercaseName := strings.ToLower(name)
		if strings.Contains(lowercaseName, searchName) {
			found = append(found, name)
		}
	}

	sort.Strings(found)

	if len(found) > maxOutput {
		found = found[:maxOutput]
	}

	callback(found)
}

func callback(names []string) {
	fmt.Println(names)
}

func main() {

	daftarNama := []string{
		"Abigail", "Alexandra", "Alison",
		"Amanda", "Angela", "Bella",
		"Carol", "Caroline", "Carolyn",
		"Deirdre", "Diana", "Elizabeth",
		"Ella", "Faith", "Olivia", "Penelope"}

	searchName(daftarNama, "AN", 3, callback)
}

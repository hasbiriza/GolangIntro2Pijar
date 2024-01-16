package main

import (
	"fmt"
	"sort"
)

func main() {
	//1
	Slice := []int{1, 2, 3, 4, 5}
	length := len(Slice) //1.Len berfungsi untuk mendapatkan panjang dari array,slice,map,string,atau channel
	fmt.Println("Panjang slice:", length)
	//2
	capacity := cap(Slice) //2.Capacity berfungsi untuk mendapatkan kapasitas dari array,slice,map,string,atau channel
	fmt.Println("Kapasitas slice:", capacity)
	//3
	makeSlice := make([]int, 3, 5) //3.Digunakan untuk membuat dan menginisialisasi slice, map, atau channel.
	fmt.Println("Slice baru:", makeSlice)
	//4
	appendSlice := []int{1, 2, 3}
	appendSlice = append(appendSlice, 4) //4.Menambahkan elemen ke dalam slice.
	fmt.Println("Slice setelah ditambahkan:", appendSlice)
	//5
	sumber := []int{1, 2, 3}
	tujuan := make([]int, 3)
	copy(tujuan, sumber) //5.Menyalin elemen dari satu slice ke slice lainnya.
	fmt.Println("Slice tujuan setelah disalin:", tujuan)
	//6
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	delete(map1, "a") //6. Method delete() digunakan untuk menghapus elemen dari map
	fmt.Println(map1)
	//7
	print("Hello, world!") //7.Method print() digunakan untuk mencetak nilai ke layar
	//8
	println("Hello, world!") //8.Method println() digunakan untuk mencetak nilai ke layar diikuti dengan karakter newline.
	//9
	slice := []int{5, 2, 1, 4, 3}
	sort.Ints(slice) //9 Method sort() digunakan untuk mengurutkan slice.
	fmt.Println(slice)

	//10
	panic("Ini adalah suatu panic!") //6.Digunakan untuk menimbulkan panic (error yang tidak dapat diatasi) dalam program.
}

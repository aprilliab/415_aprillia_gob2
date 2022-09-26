package main

// import package
import (
	"fmt"
)
type student struct {
	number 		int
	nama 		string
	alamat 		string
	pekerjaan 	string
	alasan 		string
}
func main () {
	var students = []Student{
		{number : 1, nama: "Salsabila", alamat: "jakarta selatan", pekerjaan: "mahasiswa", alasan: "ingin mencoba hal baru"},
		{number : 2, nama: "Anindita", alamat: "pasar minggu", pekerjaan: "sekretaris", alasan: "ingin memperdalam go"},
		{number : 3, nama: "Adnan Husaen", alamat: "ciledug", pekerjaan: "freelance", alasan: "ingin memperdalan backend"},
		{number : 4, nama: "Salsabila", alamat: "jakarta selatan", pekerjaan: "supervisor", alasan: "tertarik dengan web service"},
		{number : 5, nama: "Salsabila", alamat: "jakarta selatan", pekerjaan: "admin", alasan: "mengumpulkan portofolio"},
    {number : 6, nama: "Ardana", alamat: "cikeusal", pekerjaan: "operator"m alasan:"ingin lebih tau"},
	 }
}
fmt.Print("Masukkan nomor absen: ")
var input int
fmt.Scanln(&input)
findStudent(input, students)
}
func findStudent(absenNumber int ,listStudent []Student){
	for _, v := range listStudent {
		 if v.number == absenNumber {
			fmt.Println("Nama :", v.nama)
			fmt.Println("alamat :", v.alamat)
			fmt.Println("pekerjaan :", v.pekerjaan)
			fmt.Println("alasan :", v.alasan)

			}	
	}
}
if absenNumber > 7 {
	fmt.Println ("biodata tidak ditemukan")
}

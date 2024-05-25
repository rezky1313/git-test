package git_test

//jika ada perubahan  sedikit selallu update versi modul

//sebelum major update
// func Hello() string {
// 	return "Hello"
// }

//major update
//melakukan  perubahan besar pada struktur program/function
//jika melakukan perubahan besar lakukan perubahan pada modul
//contoh module github.com/rezky1313/git-test/v2
//ada penambahan pharse v2
//penambahan pharse pada modul tidak bedampak pada github
//setelah itu lakukan update pada tag versi contoh v2.0.0

//setelah perubahan besar
func Hello(name string) string {
	return "Hello" + name
}

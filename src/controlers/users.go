package controlers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}
func SearchAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}

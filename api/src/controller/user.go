package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recuperando usuario"))
}

func RetrieveUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recuoerando usuarios"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))

}

package controllers

import (
	"./models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
	"math/rand"
	"./config"
	"log"
)

var tmpl = template.Must(template.ParseGlob("views/*"))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()


	dataGender := map[string]bool{
		"laki-laki": true,
		"perempuan": true,
	}

	dataRole := map[string]bool{
		"admin": true,
		"user": true,
	}
	request := models.AddUserRequest{}
	response := models.AddUserResponse{}
	reqbyte, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(reqbyte, &request); err != nil {
		fmt.Println("Request tidak sesuai")

		response.Code = "02"
		response.Message = "Request tidak sesuai"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	if request.Name == "" {
		response.Code = "02"
		response.Message = "Request tidak sesuai, Name Harus Diisi"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	if request.Age == "" {
		response.Code = "02"
		response.Message = "Request tidak sesuai, Age Harus Diisi"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	if request.Email == "" {
		response.Code = "02"
		response.Message = "Request tidak sesuai, Email Harus Diisi"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	if _, ok := dataGender[request.Gender]; !ok {
		response.Code = "02"
		response.Message = "Request tidak sesuai"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}
	//Cek role data
	if _, ok := dataRole[request.Role]; !ok {
		response.Code = "02"
		response.Message = "Request tidak sesuai"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	rand.Seed(time.Now().UnixNano())
	request.Password = randSeq(10)

	if r.Method == "POST" {
		name := request.Name
		age := request.Age
		email := request.Email
		password := request.Password
		gender := request.Gender
		role := request.Role
		insForm, err := db.Prepare("INSERT INTO users (name, age, email, password, gender, role) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, age, email, password, gender, role)
		log.Println("Insert Data: name " + name + " | age " + age + " | email " + email + " | gender " + gender + " | role " + role)

		response.Code = "00"
		response.Message = "Berhasil"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func ViewEditUser(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id, name, age, email, gender, role FROM users WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	user := models.AddUserRequest{}

	for selDB.Next() {
		var id, name, age, email, gender, role string
		err := selDB.Scan(&id, &name, &age, &email, &gender, &role)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.Age = age
		user.Email = email
		user.Gender = gender
		user.Role = role
	}

	tmpl.ExecuteTemplate(w, "Edit", user)
	defer db.Close()
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		email := r.FormValue("email")
		gender := r.FormValue("gender")
		role := r.FormValue("role")
		id := r.FormValue("id")
		insForm, err := db.Prepare("UPDATE users SET name=?, age=?, email=?, gender=?, role=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, age, email, gender, role, id)
		// log.Println("UPDATE Data: name " + name + " | category " + category + " | url " + url + " | rating " + rating + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func ViewDetail(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM users WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	user := models.AddUserRequest{}

	for selDB.Next() {
		var id, name, age, email, gender, role, password string
		err := selDB.Scan(&id, &name, &age, &email, &gender, &role, &password)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.Age = age
		user.Email = email
		user.Password = password
		user.Gender = gender
		user.Role = role
	}
	tmpl.ExecuteTemplate(w, "Show", user)
	defer db.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()
	tool := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tool)
	log.Println("DELETE " + tool)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func ViewAddUser(w http.ResponseWriter, r *http.Request) {
	res := []models.AddUserRequest{}
	tmpl.ExecuteTemplate(w, "Add", res)
}

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := config.DBInit()

	selDB, err := db.Query("SELECT id, name, age, email, gender, role FROM users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	user := models.AddUserRequest{}
	res := []models.AddUserRequest{}

	for selDB.Next() {
		//var id, rating int
		var id, name, age, email, gender, role string
		err := selDB.Scan(&id, &name, &age, &email, &gender, &role)
		if err != nil {
			panic(err.Error())
		}
		//log.Println("Listing Row: Id " + id + " | name " + name + " | age " + age + " | email " + email + " | gender " + gender + " | role " + role)

		user.Id = id
		user.Name = name
		user.Age = age
		user.Email = email
		user.Gender = gender
		user.Role = role
		res = append(res, user)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

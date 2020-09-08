package boba

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var db *sql.DB

type Boba struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func Init(database *sql.DB) {
	db = database
	statement, err := db.Prepare("CREATE Table boba(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, price TEXT NOT NULL)")
	if err != nil {
		panic(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err.Error())
	}
}

func CreateBoba(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	statement, err := db.Prepare("INSERT INTO boba (name, price) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	price := keyVal["price"]
	if name != "" && price != "" {
		_, err = statement.Exec(name, price)
		if err != nil {
			panic(err.Error())
		}
		fmt.Fprintf(writer, "New boba was created.")
	} else {
		fmt.Fprintf(writer, "Failed to create new boba.")
	}
}

func GetBoba(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	result, err := db.Query("SELECT id, name, price FROM boba WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var boba Boba
		err := result.Scan(&boba.ID, &boba.Name, &boba.Price)
		if err != nil {
			panic(err.Error())
		}
		json.NewEncoder(writer).Encode(boba)
	}
}

func GetBobas(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT id, name, price FROM boba")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var bobas []Boba
	for result.Next() {
		var boba Boba
		err := result.Scan(&boba.ID, &boba.Name, &boba.Price)
		if err != nil {
			panic(err.Error())
		}
		bobas = append(bobas, boba)
	}
	json.NewEncoder(writer).Encode(bobas)
}

func UpdateBoba(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	statement, err := db.Prepare("UPDATE boba SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["name"]
	newPrice := keyVal["price"]
	if newName != "" && newPrice != "" {
		_, err = statement.Exec(newName, newPrice, params["id"])
		if err != nil {
			panic(err.Error())
		}
		fmt.Fprintf(writer, "Boba with ID = %s was updated", params["id"])
	} else {
		fmt.Fprintf(writer, "Boba with ID = %s was  not updated", params["id"])
	}
}

func DeleteBoba(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	statement, err := db.Prepare("DELETE FROM boba WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = statement.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(writer, "Boba with ID = %s was deleted", params["id"])
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type MTravel struct {
	TravelID         string `json:"TravelID"`
	TravelName       string `json:"TravelName"`
	CompanyName      string `json:"CompanyName"`
	LicenseNumber    string `json:"LicenseNumber"`
	LicenseNumberHaj string `json:"LicenseNumberHaj"`
	Address          string `json:"Address"`
	ProvinceID       string `json:"ProvinceID"`
	CityID           string `json:"CityID"`
	ContactPerson    string `json:"ContactPerson"`
	MobileNumber     string `json:"MobileNumber"`
	OfficeNumber     string `json:"OfficeNumber"`
	Email            string `json:"Email"`
	OfficeEmail      string `json:"OfficeEmail"`
	Website          string `json:"Website"`
	TravelStatus     string `json:"TravelStatus"`
	Rating           string `json:"Rating"`
	UsrName          string `json:"UsrName"`
	Psword           string `json:"Psword"`
	Logo             string `json:"Logo"`
}

// Get all orders

func getMTravel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m_travel []MTravel

	sql := `SELECT
				TravelID,
				IFNULL(TravelName,''),
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(LicenseNumber,'') LicenseNumber,
				IFNULL(LicenseNumberHaj,'') LicenseNumberHaj,
				IFNULL(Address,'') Address,
				IFNULL(ProvinceID,'') ProvinceID,
				IFNULL(CityID,'') CityID ,
				IFNULL(ContactPerson,'') ContactPerson,
				IFNULL(MobileNumber,'') MobileNumber,
				IFNULL(OfficeNumber,'') OfficeNumber,
				IFNULL(Email,'') Email,
				IFNULL(OfficeEmail,'') OfficeEmail,
				IFNULL(Website,'') Website ,
				IFNULL(TravelStatus,'') TravelStatus,
				IFNULL(Rating,'') Rating,
				IFNULL(UsrName,'') UsrName,
				IFNULL(Psword,'') Psword ,
				IFNULL(Logo,'') Logo
			FROM m_travel`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var mTravel MTravel
		err := result.Scan(&mTravel.TravelID, &mTravel.TravelName, &mTravel.CompanyName,
			&mTravel.LicenseNumber, &mTravel.LicenseNumberHaj, &mTravel.Address, &mTravel.ProvinceID,
			&mTravel.CityID, &mTravel.ContactPerson, &mTravel.MobileNumber, &mTravel.OfficeNumber,
			&mTravel.Email, &mTravel.OfficeEmail, &mTravel.Website, &mTravel.TravelStatus,
			&mTravel.Rating, &mTravel.UsrName, &mTravel.Psword, &mTravel.Logo)

		if err != nil {
			panic(err.Error())
		}
		m_travel = append(m_travel, mTravel)
	}

	json.NewEncoder(w).Encode(m_travel)
}

func createMTravel(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		TravelID := r.FormValue("travelID")
		TravelName := r.FormValue("travelName")
		CompanyName := r.FormValue("companyName")
		LicenseNumber := r.FormValue("licenseNumber")
		LicenseNumberHaj := r.FormValue("licenseNumberHaj")
		Address := r.FormValue("address")
		ProvinceID := r.FormValue("provinceID")
		CityID := r.FormValue("cityID")
		ContactPerson := r.FormValue("contactPerson")
		MobileNumber := r.FormValue("mobileNumber")
		OfficeNumber := r.FormValue("officeNumber")
		Email := r.FormValue("email")
		OfficeEmail := r.FormValue("officeEmail")
		Website := r.FormValue("website")
		TravelStatus := r.FormValue("travelStatus")
		Rating := r.FormValue("rating")
		UsrName := r.FormValue("usrName")
		Psword := r.FormValue("psword")
		Logo := r.FormValue("logo")

		stmt, err := db.Prepare("INSERT INTO m_travel (TravelID,TravelName,CompanyName,LicenseNumber,LicenseNumberHaj,Address,ProvinceID,CityID,ContactPerson,MobileNumber,OfficeNumber,Email,OfficeEmail,Website,TravelStatus,Rating,UsrName,Psword,Logo ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(TravelID, TravelName, CompanyName, LicenseNumber, LicenseNumberHaj, Address, ProvinceID, CityID, ContactPerson, MobileNumber, OfficeNumber, Email, OfficeEmail, Website, TravelStatus, Rating, UsrName, Psword, Logo)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {

			fmt.Fprintf(w, "Date Created")
			//http.Redirect(w, r, "/", 301)
		}

	}
}

func getMTravels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m_travel []MTravel
	params := mux.Vars(r)

	sql := `SELECT
				TravelID,
				IFNULL(TravelName,''),
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(LicenseNumber,'') LicenseNumber,
				IFNULL(LicenseNumberHaj,'') LicenseNumberHaj,
				IFNULL(Address,'') Address,
				IFNULL(ProvinceID,'') ProvinceID,
				IFNULL(CityID,'') CityID ,
				IFNULL(ContactPerson,'') ContactPerson,
				IFNULL(MobileNumber,'') MobileNumber,
				IFNULL(OfficeNumber,'') OfficeNumber,
				IFNULL(Email,'') Email,
				IFNULL(OfficeEmail,'') OfficeEmail,
				IFNULL(Website,'') Website ,
				IFNULL(TravelStatus,'') TravelStatus,
				IFNULL(Rating,'') Rating,
				IFNULL(UsrName,'') UsrName,
				IFNULL(Psword,'') Psword ,
				IFNULL(Logo,'') Logo
			FROM m_travel WHERE TravelID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mTravel MTravel

	for result.Next() {

		err := result.Scan(&mTravel.TravelID, &mTravel.TravelName, &mTravel.CompanyName,
			&mTravel.LicenseNumber, &mTravel.LicenseNumberHaj, &mTravel.Address, &mTravel.ProvinceID,
			&mTravel.CityID, &mTravel.ContactPerson, &mTravel.MobileNumber, &mTravel.OfficeNumber,
			&mTravel.Email, &mTravel.OfficeEmail, &mTravel.Website, &mTravel.TravelStatus,
			&mTravel.Rating, &mTravel.UsrName, &mTravel.Psword, &mTravel.Logo)

		if err != nil {
			panic(err.Error())
		}

		m_travel = append(m_travel, mTravel)
	}

	json.NewEncoder(w).Encode(m_travel)
}

func updateMTravel(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newTravelName := r.FormValue("TravelName")

		stmt, err := db.Prepare("UPDATE m_travel SET TravelName = ? WHERE TravelID = ?")

		_, err = stmt.Exec(newTravelName, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Travel with TravelID = %s was updated", params["id"])
	}
}

func deleteMTravel(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM m_travel WHERE TravelID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Travel with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var m_travel []MTravel

	TravelID := r.FormValue("travelID")
	TravelName := r.FormValue("travelName")

	sql := `SELECT
				TravelID,
				IFNULL(TravelName,''),
				IFNULL(CompanyName,'') CompanyName,
				IFNULL(LicenseNumber,'') LicenseNumber,
				IFNULL(LicenseNumberHaj,'') LicenseNumberHaj,
				IFNULL(Address,'') Address,
				IFNULL(ProvinceID,'') ProvinceID,
				IFNULL(CityID,'') CityID ,
				IFNULL(ContactPerson,'') ContactPerson,
				IFNULL(MobileNumber,'') MobileNumber,
				IFNULL(OfficeNumber,'') OfficeNumber,
				IFNULL(Email,'') Email,
				IFNULL(OfficeEmail,'') OfficeEmail,
				IFNULL(Website,'') Website ,
				IFNULL(TravelStatus,'') TravelStatus,
				IFNULL(Rating,'') Rating,
				IFNULL(UsrName,'') UsrName,
				IFNULL(Psword,'') Psword ,
				IFNULL(Logo,'') Logo
			FROM m_travel WHERE TravelID = ? AND TravelName = ?`

	result, err := db.Query(sql, TravelID, TravelName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var mTravel MTravel

	for result.Next() {

		err := result.Scan(&mTravel.TravelID, &mTravel.TravelName, &mTravel.CompanyName,
			&mTravel.LicenseNumber, &mTravel.LicenseNumberHaj, &mTravel.Address, &mTravel.ProvinceID,
			&mTravel.CityID, &mTravel.ContactPerson, &mTravel.MobileNumber, &mTravel.OfficeNumber,
			&mTravel.Email, &mTravel.OfficeEmail, &mTravel.Website, &mTravel.TravelStatus,
			&mTravel.Rating, &mTravel.UsrName, &mTravel.Psword, &mTravel.Logo)

		if err != nil {
			panic(err.Error())
		}

		m_travel = append(m_travel, mTravel)
	}

	json.NewEncoder(w).Encode(m_travel)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/m_travel", getMTravel).Methods("GET")
	r.HandleFunc("/m_travel/{id}", getMTravels).Methods("GET")
	r.HandleFunc("/m_travel", createMTravel).Methods("POST")
	r.HandleFunc("/m_travel/{id}", updateMTravel).Methods("PUT")
	r.HandleFunc("/m_travel/{id}", deleteMTravel).Methods("DELETE")

	//new
	r.HandleFunc("/getmtravel", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))
}

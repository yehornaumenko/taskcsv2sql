package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Vehicle struct {
	Id           int     `csv:"id"`
	URL          string  `csv:"url"`
	Region       string  `csv:"region"`
	RegionURL    string  `csv:"region_url"`
	Price        float32 `csv:"price"`
	Year         int     `csv:"year"`
	Manufacturer string  `csv:"manufacturer"`
	Model        string  `csv:"model"`
	Condition    string  `csv:"model"`
	Cylinders    string  `csv:"cylinders"`
	Fuel         string  `csv:"fuel"`
	Odometer     string  `csv:"odometer"`
	TitleStatus  string  `csv:"title_status"`
	Transmission string  `csv:"transmission"`
	VIN          string  `csv:"vin"`
	Drive        string  `csv:"drive"`
	Size         string  `csv:"size"`
	Type         string  `csv:"type"`
	PaintColor   string  `csv:"paint_color "`
	ImageUrl     string  `csv:"image_url"`
	Description  string  `csv:"description"`
	County       string  `csv:"county"`
	State        string  `csv:"state"`
	Lat          string  `csv:"lat"`
	Long         string  `csv:"long"`
	PostingDate  string  `csv:"posting_date"`
}

type VehicleModel struct {
	gorm.Model
	Id           int     `db:"id"`
	URL          string  `db:"url"`
	Region       string  `db:"region"`
	RegionURL    string  `db:"region_url"`
	Price        float32 `db:"price"`
	Year         int     `db:"year"`
	Manufacturer string  `db:"manufacturer"`
	NewModel     string  `db:"model"`
	Condition    string  `db:"model"`
	Cylinders    string  `db:"cylinders"`
	Fuel         string  `db:"fuel"`
	Odometer     string  `db:"odometer"`
	TitleStatus  string  `db:"title_status"`
	Transmission string  `db:"transmission"`
	VIN          string  `db:"vin"`
	Drive        string  `db:"drive"`
	Size         string  `db:"size"`
	Type         string  `db:"type"`
	PaintColor   string  `db:"paint_color "`
	ImageUrl     string  `db:"image_url"`
	Description  string  `db:"description"`
	County       string  `db:"county"`
	State        string  `db:"state"`
	Lat          string  `db:"lat"`
	Long         string  `db:"long"`
	PostingDate  string  `db:"posting_date"`
}

func main() {
	clientsFile, err := os.OpenFile("vehicles.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []Vehicle{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	MakeDB(clients)
}

func MakeDB(vehicles []Vehicle) {
	db, err := gorm.Open(sqlite.Open("vehicles.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.Table("vehicles").AutoMigrate(&Vehicle{})
	if err != nil {
		panic(err)
	}

	cre := vehicles[:5]
	tx := db.Create(&cre)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
}
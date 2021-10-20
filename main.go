package main

import (
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	configFileName *string

	postgresDSN        string
	vehiclesSQLiteName string
	csvFileName        string
)

const (
	vehiclesTableName = "vehicles"
)

type conf struct {
		PostgresDSN    string `yaml:"postgresDSN"`
		SQLiteFileName string `yaml:"sqliteFileName"`
		CSVFileName    string `yaml:"csvFileName"`
	}

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

func configFromFile(fileName string) *conf {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	config := &conf{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}

	return config
}

func main() {
	configFileName = parseCLIArguments()

	cfg := configFromFile(*configFileName)

	postgresDSN = cfg.PostgresDSN
	vehiclesSQLiteName = cfg.SQLiteFileName
	csvFileName = cfg.CSVFileName

	clientsFile, err := os.OpenFile(csvFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	var clients []Vehicle

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	makeSQLiteTable(clients)
}

func parseCLIArguments() *string {
	const configFlagName = "config"

	configFileName = flag.String("config", "", "config filename in the root of directory")
	flag.Parse()

	if configFileName == nil || *configFileName == "" {
		panic(errors.Errorf("Argument %q is empty", configFlagName))
	}

	return configFileName
}

func makeSQLiteTable(vehicles []Vehicle) {
	db, err := gorm.Open(sqlite.Open(vehiclesSQLiteName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.Table(vehiclesTableName).AutoMigrate(&Vehicle{})
	if err != nil {
		panic(err)
	}

	cre := vehicles[:5]
	tx := db.Create(&cre)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
}

func makePostgresTable(vehicles []Vehicle) {

	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Table(vehiclesTableName).AutoMigrate(&Vehicle{})
	if err != nil {
		panic(err)
	}

	cre := vehicles[:5]
	tx := db.Create(&cre)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

}

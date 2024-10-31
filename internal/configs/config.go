package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost                string
	DBPort                int
	DBUser                string
	DBName                string
	DBPassword            string
	ServerPort            string
	JwtSecretKey          string
	PgConnectionString    string
	FileKeys              string
	Bucket                string
	BaseUrlGoogleStorage  string
	QueryingPersonData    string
	FileConfigFirebase    string
	FileConfigGoogleDrive string
	FileTokenGooGleDrive string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Invalid DB_PORT: %v", err)
		return nil, err
	}
	config := &Config{
		DBHost:                os.Getenv("DB_HOST"),
		DBPort:                port,
		DBUser:                os.Getenv("DB_USER"),
		DBPassword:            os.Getenv("DB_PASSWORD"),
		DBName:                os.Getenv("DB_NAME"),
		ServerPort:            os.Getenv("PORT"),
		JwtSecretKey:          os.Getenv("JWT_SECRET_KEY"),
		PgConnectionString:    os.Getenv("POSTGRES_CONNECTION_STRING"),
		FileKeys:              os.Getenv("FILE_CONFIG_GOOGLE_STORAGE"),
		Bucket:                os.Getenv("BUCKET_NAME"),
		BaseUrlGoogleStorage:  os.Getenv("BASE_URL_GOOGLE_STORAGE"),
		FileConfigFirebase:    os.Getenv("FILE_CONFIG_FIREBASE"),
		FileConfigGoogleDrive: os.Getenv("FILE_CONFIG_GOOGLE_DRIVE"),
		FileTokenGooGleDrive: os.Getenv("FILE_TOKEN_GOOGLE_DRIVE"),
		QueryingPersonData: os.Getenv("QUERYING_PERSONDATA"),
	}

	// config := &Config{
	// 	DBHost:               "localhost", //os.Getenv("DB_HOST"),
	// 	DBPort:               5432,
	// 	DBUser:               "root",                                                              //os.Getenv("DB_USER"),
	// 	DBPassword:           "123",                                                               //os.Getenv("DB_PASSWORD"),
	// 	DBName:               "beautyfinder",                                                      //os.Getenv("DB_NAME"),
	// 	ServerPort:           "5000",                                                              //os.Getenv("PORT"),
	// 	JwtSecretKey:         "230E7C21-DE21-4506-BFBD-18F7319B3FC4",                              //os.Getenv("JWT_SECRET_KEY"),
	// 	PgConnectionString:   "postgresql://postgres.jhufdsjhajeslrjiavkp:%23wert%40y123%40@aws-0-us-east-1.pooler.supabase.com:6543/postgres", //os.Getenv("POSTGRES_CONNECTION_STRING"),
	// 	FileKeys:             "../land-8c098-d87756162e9f.json",                                   //os.Getenv("FILE_CONFIG_GOOGLE_STORAGE"),
	// 	Bucket:               "document-land",                                                     //os.Getenv("BUCKET_NAME"),
	// 	BaseUrlGoogleStorage: "https://storage.googleapis.com",                                    //os.Getenv("BASE_URL_GOOGLE_STORAGE"),
	// 	QueryingPersonData:   "https://www.sepe.gov.ao/ao/actions/bi.ajcall.php?bi",
	// }
	return config, nil
}

package configs

type Config struct{
	DBHost               string
	DBPort               int
	DBUser               string
	DBName               string
	DBPassword           string
	ServerPort           string
	JwtSecretKey         string
	PgConnectionString   string
	FileKeys             string
	Bucket               string
	BaseUrlGoogleStorage string
}

func LoadConfig() (*Config, error) {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Printf("Error loading .env file: %v", err)
	// }
	// port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	log.Printf("Invalid DB_PORT: %v", err)
	// 	return nil, err
	// }

	// config := &Config{
	// 	DBHost:               os.Getenv("DB_HOST"),
	// 	DBPort:               port,
	// 	DBUser:               os.Getenv("DB_USER"),
	// 	DBPassword:           os.Getenv("DB_PASSWORD"),
	// 	DBName:               os.Getenv("DB_NAME"),
	// 	ServerPort:           os.Getenv("PORT"),
	// 	JwtSecretKey:         os.Getenv("JWT_SECRET_KEY"),
	// 	PgConnectionString:   os.Getenv("POSTGRES_CONNECTION_STRING"),
	// 	FileKeys:             os.Getenv("FILE_CONFIG_GOOGLE_STORAGE"),
	// 	Bucket:               os.Getenv("BUCKET_NAME"),
	// 	BaseUrlGoogleStorage: os.Getenv("BASE_URL_GOOGLE_STORAGE"),
	// }

	config := &Config{
		DBHost:               "localhost", //os.Getenv("DB_HOST"),
		DBPort:               5432,
		DBUser:               "root", //os.Getenv("DB_USER"),
		DBPassword:           "123", //os.Getenv("DB_PASSWORD"),
		DBName:               "beautyfinder", //os.Getenv("DB_NAME"),
		ServerPort:           "5000", //os.Getenv("PORT"),
		JwtSecretKey:         "230E7C21-DE21-4506-BFBD-18F7319B3FC4", //os.Getenv("JWT_SECRET_KEY"),
		PgConnectionString:   "postgresql://root:123@localhost:5432/beautyfinder?sslmode=disable",   //os.Getenv("POSTGRES_CONNECTION_STRING"),
		FileKeys:             "land-8c098-d87756162e9f.json",//os.Getenv("FILE_CONFIG_GOOGLE_STORAGE"),
		Bucket:               "document-land", //os.Getenv("BUCKET_NAME"),
		BaseUrlGoogleStorage: "https://storage.googleapis.com", //os.Getenv("BASE_URL_GOOGLE_STORAGE"),
	}
	return config, nil
}


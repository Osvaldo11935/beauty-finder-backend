package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"src/internal/configs"
	service_interface "src/internal/services/interface_services"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDriveService struct {
	srv *drive.Service
}

func NewGoogleDriveService() service_interface.IFileManager {
	envConfig, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return nil
	}

	b, err := os.ReadFile(envConfig.FileConfigGoogleDrive)
	if err != nil {
		log.Fatalf("unable to read client secret file: %v", err)
		return nil
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("unable to parse client secret file to config: %v", err)
		return nil
	}

	client := GetClient(config, envConfig)

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("unable to retrieve Drive client: %v", err)
		return nil
	}

	return &GoogleDriveService{srv: srv}
}

func (d *GoogleDriveService) Upload(c *gin.Context, filePath *string) (*string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	fileName := file.Filename
	if fileName == "" {
		return nil, err
	}

	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	driveFile := &drive.File{Name: fileName}

	uploadedFile, err := d.srv.Files.Create(driveFile).Media(f).Do()

	if err != nil {
		return nil, err
	}

	return &uploadedFile.Id, nil
}

// func (d *GoogleDriveService) Download(fileId string) (*string, error) {
// 	resp, err := d.srv.Files.Get(fileId).Download()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to download file: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file data: %v", err)
// 	}
// 	encodedData := base64.StdEncoding.EncodeToString(data)
// 	contentType := resp.Header.Get("Content-Type")
// 	dataURI := fmt.Sprintf("data:%s;base64,%s", contentType, encodedData)

// 	return &dataURI, nil
// }

// func (d *GoogleDriveService) Download(fileId string) ([]byte, error) {
// 	resp, err := d.srv.Files.Get(fileId).Download()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	data, err := io.ReadAll(resp.Body)
// 	return data, err
// }
func (s *GoogleDriveService) GetFileUrl(filePath string) (*string, error) {
	return nil, nil
}
func (s *GoogleDriveService) Download(filePath string) ([]byte, error) {
	return nil, nil
}

func GetClient(config *oauth2.Config, envConfig *configs.Config) *http.Client {
	tokFile := envConfig.FileTokenGooGleDrive
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	// if _, err := fmt.Scan(&authCode); err != nil {
	// 	log.Fatalf("Unable to read authorization code: %v", err)
	// }

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"src/internal/configs"
	service_interface "src/internal/services/interface_services"

	"github.com/gin-gonic/gin"
)

type SupabaseStorageService struct {
	BaseURL string
	Bucket  string
}

func NewSupabaseStorageService() service_interface.IFileManager {
	loadEnv, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return nil
	}

	return &SupabaseStorageService{BaseURL: loadEnv.SupaBaseUrl, Bucket: loadEnv.SupaBaseBucket}

}

func (s *SupabaseStorageService) Upload(fctx *gin.Context, filePath *string) (*string, error) {

	fileHeader, err := fctx.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("erro ao obter arquivo do contexto: %w", err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.BaseURL, s.Bucket, filePath)

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileContent))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("erro ao fazer upload: Status %d - %s", resp.StatusCode, string(body))
	}

	log.Println("Upload realizado com sucesso!")

	return &fileHeader.Filename, nil
}

func (s *SupabaseStorageService) GetFileUrl(filePath string) (*string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.BaseURL, s.Bucket, filePath)
	return &url, nil
}
func (s *SupabaseStorageService) Download(filePath string) ([]byte, error) {
	url := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.BaseURL, s.Bucket, filePath)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("erro ao baixar arquivo: Status %d - %s", resp.StatusCode, string(body))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	log.Println("Download realizado com sucesso!")
	return data, nil
}

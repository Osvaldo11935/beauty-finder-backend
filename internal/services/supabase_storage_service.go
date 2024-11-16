package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"src/internal/configs"
	service_interface "src/internal/services/interface_services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseStorageService struct {
	BaseURL     string
	Bucket      string
	StorageURL  string
	AccessToken string
}

func NewSupabaseStorageService() service_interface.IFileManager {
	loadEnv, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return nil
	}

	return &SupabaseStorageService{
		BaseURL:     loadEnv.SupaBaseUrl,
		Bucket:      loadEnv.SupaBaseBucket,
		AccessToken: loadEnv.SupaBaseToken,
		StorageURL:  loadEnv.SupaBaseStorageUrl}

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

	storageClient := storage_go.NewClient(s.StorageURL, s.AccessToken, nil)

	fileReader := bytes.NewReader(fileContent)

	fileNameSplit := strings.Split(fileHeader.Filename, ".")

	if len(fileNameSplit) < 2 {
		return nil, fmt.Errorf("erro: nome do arquivo inválido, sem extensão")
	}

	ext := fileNameSplit[len(fileNameSplit)-1]

	fileName := fmt.Sprintf("%s.%s", uuid.New().String(), ext)

	path := fmt.Sprintf("%s/%s", *filePath, fileName)

	contentType := fileHeader.Header.Get("Content-Type")

	result, err := storageClient.UploadFile(s.Bucket, path, fileReader, storage_go.FileOptions{
		ContentType: &contentType})

	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	log.Println("Result ", result)
	log.Println("Upload realizado com sucesso!")

	return &fileName, nil
}

func (s *SupabaseStorageService) GetFileUrl(filePath string) (*string, error) {

	storageClient := storage_go.NewClient(s.StorageURL, s.AccessToken, nil)

	result := storageClient.GetPublicUrl(s.Bucket, filePath)

	return &result.SignedURL, nil
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

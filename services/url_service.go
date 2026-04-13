package services

import (
	"fmt"
	"time"

	"github.com/go-short/models"
	"github.com/go-short/repository"
	"github.com/go-short/utils"
)

func CreateURL(original string) (models.URL, error) {
	code := utils.GenerateShortCode(6)
	fmt.Println(code)

	url := models.URL{
		ShortCode: code,
		URL:       original,
		Clicks:    0,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	url, err := repository.Save(url)
	return url, err
}

func GetURL(code string) (models.URL, error) {
	return repository.Get(code)
}

func UpdateURL(code string, newURL string) error {
	return repository.Update(code, newURL)
}

func DeleteURL(code string) error {
	return repository.Delete(code)
}

func IncrementClicks(code string) {
	repository.IncrementClicks(code)
}

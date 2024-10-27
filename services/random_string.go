package services

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type RandomStringService interface {
	GetRandomString(length int) (string, error)
}

type RealRandomStringService struct {
	Client *resty.Client
}

func (r *RealRandomStringService) GetRandomString(length int) (string, error) {
	url := fmt.Sprintf("https://www.random.org/strings/?num=1&len=%d&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new", length)

	resp, err := r.Client.R().Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", errors.New("failed to fetch random string from external API")
	}

	return resp.String(), nil
}

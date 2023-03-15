package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Asliddin3/open-api-bot/api/repo"
	"github.com/Asliddin3/open-api-bot/storage"
)

const (
	publicAPIURL = "https://pixabay.com/api/"
	accessKey    = "34408309-5af3d1cfbbe0b2d0c4bc877a1"
)

var (
	//ErrGotEmptyResponse will return when public api return empty response
	ErrGotEmptyResponse = errors.New("got empty response from api")
)

type PixabayClient struct{}

//NewPixabayClient this func will register storage for saving user search
func NewPixabayClient(db *storage.StorageI) repo.PixabayRepo {
	return &PixabayClient{}
}

type Hit struct {
	ID      int64  `json:"id"`
	PageURL string `json:"pageURL"`
}

type Response struct {
	Total     int64 `json:"total"`
	TotalHits int64 `json:"totalHits"`
	Hits      []Hit `json:"hits"`
}

//GetVideo this fun will get video from pixabay api video by name
func (p *PixabayClient) GetVideo(name string) (string, error) {
	urlReq := fmt.Sprintf("%s/videos/?key=%s&q=%s", publicAPIURL, accessKey, name)
	req, err := http.NewRequest("GET", urlReq, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var apiResponse Response
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return "", fmt.Errorf("got error while decode video response %v", err)
	}
	hitsCount := apiResponse.TotalHits
	if hitsCount == 0 {
		return "", ErrGotEmptyResponse
	}
	url := apiResponse.Hits[getRand(int(hitsCount))].PageURL

	return url, nil
}

//GetPhoto this fun will get photo from pixabay api photo by name
func (p *PixabayClient) GetPhoto(name string) (string, error) {
	urlReq := fmt.Sprintf("%s/?key=%s&q=%s&image_type=photo", publicAPIURL, accessKey, name)
	req, err := http.NewRequest("GET", urlReq, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var apiResponse Response
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return "", fmt.Errorf("got error while decode photo response %v", err)
	}
	hitsCount := apiResponse.TotalHits
	if hitsCount == 0 {
		return "", ErrGotEmptyResponse
	}
	url := apiResponse.Hits[getRand(int(hitsCount))].PageURL
	return url, nil
}

func getRand(totalHits int) int {
	min := 0
	max := totalHits
	index := rand.Intn(max-min) + min
	return index
}

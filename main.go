package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	baseURL       = "https://api.nasa.gov"
	roversBaseUrl = baseURL + "/mars-photos/api/v1/rovers/curiosity/photos?earth_date=%s&api_key=%s"
)

var (
	apiKey = os.Getenv("NASA_API_KEY")
)

type RoverResp struct {
	Photos []Photo `json:"photos"`
}

type Photo struct {
	ImageSRC string `json:"img_src"`
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Date argument is required")
		os.Exit(1)
	}

	dateStr := args[0]
	if err := getRoverFotos(dateStr); err != nil {
		fmt.Printf("error: %v", err.Error())
		os.Exit(2)
	}
}

func getRoverFotos(date string) error {
	res, err := http.Get(fmt.Sprintf(roversBaseUrl, date, apiKey))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out := RoverResp{}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return err
	}

	fmt.Println("max photos found", len(out.Photos))
	for i, photo := range out.Photos {
		if i >= 10 {
			break
		}
		fmt.Println(photo.ImageSRC)
	}
	return nil
}

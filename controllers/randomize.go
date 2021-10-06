package controllers

import(
	"weather-App/models"
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
    "io/ioutil"
)

func Randomize(weather *models.Weather)  {
	ticker := time.NewTicker(5 * time.Second)
		quit := make(chan struct{})
		for {
			select {
			case <- ticker.C:
				rand.Seed(time.Now().UnixNano())
				weather.Water =  rand.Intn(100)
				weather.Wind = rand.Intn(100)
				fmt.Println(weather)
		
				weatherJson, _ := json.Marshal(weather)
				err := ioutil.WriteFile("output.json", weatherJson, 0644)
				if err != nil {
					fmt.Println("error", err.Error())
					return 
				}
				//  fmt.Printf("%+v", weather)
			case <- quit:
					ticker.Stop()
					return
			}

		}
}
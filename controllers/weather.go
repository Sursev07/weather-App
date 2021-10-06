package controllers

import (
	"weather-App/models"
	// "encoding/json"
    "fmt"
    // "io/ioutil"
	// "math/rand"
	// "time"
	"net/http"
	"html/template"
	"path"
	// "runtime"

	// "github.com/gin-gonic/gin"
)


func GetWeather(w http.ResponseWriter, r *http.Request, weather *models.Weather) {

	
	// go func ()  {
		
	// }()
   
	// err := json.Unmarshal(jsonBlob, &weather)
	
	fmt.Println(weather, ">>>")
	var filepath = path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	 
	var data = map[string]interface{}{
		"Water": weather.Water,
		"Wind":  weather.Wind,
		"Status": CheckStatus(weather.Wind, weather.Water),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
}

func  CheckStatus(wind int, water int) string  {
	var status = ""
	if water<5 && wind <6 {
		status = "Aman"
	} else if (water >= 6 && water <= 8) && (wind >= 7 && wind <=15 ) {
		status = "Siaga"
	} else if water > 8 &&  wind >15  {
		status = "Bahaya"
	} else {
		status = "Unpredictable"
	}
	fmt.Println(status, "statusss")
	return status
}

// func GetWeather(w http.ResponseWriter, r *http.Request) {
// 	// cCp := c.Copy()
//     ticker := time.NewTicker(5 * time.Second)
//     quit := make(chan struct{})
//     go func() {
//         for {
//            select {
//             case <- ticker.C:
// 				weather := models.Weather{}
// 	// var jsonBlob = []byte(`{Water:0, Waind:0}`)
// 				rand.Seed(time.Now().UnixNano())
// 				weather.Water =  rand.Intn(100)
// 				weather.Wind = rand.Intn(100)
// 				fmt.Println(weather)
// 				// err := json.Unmarshal(jsonBlob, &weather)
				

// 				weatherJson, _ := json.Marshal(weather)
// 				err := ioutil.WriteFile("output.json", weatherJson, 0644)
// 				if err != nil {
// 					fmt.Println("error", err.Error())
// 					return 
// 				}
// 				fmt.Printf("%+v", weather)
				
// 				var filepath = path.Join("views", "index.html")
// 				tmpl, err := template.ParseFiles(filepath)
// 				if err != nil {
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 					return
// 				}

// 				var data = map[string]interface{}{
// 					"Water": weather.Water,
// 					"Wind":  weather.Wind,
// 					"Status": CheckStatus(weather),
// 				}

// 				err = tmpl.Execute(w, data)
// 				if err != nil {
// 					fmt.Println("error", err.Error())
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 					return 
// 				}
				
// 				// result := gin.H {
// 				// 	"status": weather,
// 				// }
// 				//  OutputJSON(w, weather)
//             case <- quit:
//                 ticker.Stop()
//                 return
//             }
//         }
//      }()
// }
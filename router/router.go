package router

import(
	"weather-App/controllers"
	"weather-App/models"
	"fmt"
	"net/http"
	"runtime"
)
func StartApp()  {
	// router := gin.Default()
	fmt.Println("SEVENNN")
	runtime.GOMAXPROCS(2)
	var weather models.Weather
	go controllers.Randomize(&weather)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// fmt.Println(">>>ini", weather)
		controllers.GetWeather(w, r, &weather)
	})
	fmt.Println("Run on server 9000")
	http.ListenAndServe(":9000", nil)
}
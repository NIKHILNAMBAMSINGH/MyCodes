package bootstraps

import (
	"GolangmMicroservices/webapp/controllers"
	"net/http"
)

func BootApplication() {
	http.HandleFunc("/employee", controllers.GetEmployee)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

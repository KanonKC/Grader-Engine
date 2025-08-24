package routes

import (
	controllers_grader "ModelGrader-Grader/controllers/grader"
	services_grader "ModelGrader-Grader/services/grader"
	services_sandbox "ModelGrader-Grader/services/sandbox"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	sandboxSvc := services_sandbox.New(8)
	sandboxSvc.Init()
	graderSvc := services_grader.New(sandboxSvc)
	graderCtrl := controllers_grader.New(graderSvc)

	// Make post method
	mux.HandleFunc("/output", graderCtrl.GenerateOutput)
}

package controllers_grader

import (
	services_grader "ModelGrader-Grader/services/grader"
	"ModelGrader-Grader/types"
	"encoding/json"
	"net/http"
)

type GenerateOutputRequest struct {
	Code  string   `json:"code"`
	Lang  string   `json:"lang"`
	Input []string `json:"input"`
}

type GraderController interface {
	GenerateOutput(w http.ResponseWriter, r *http.Request)
}

type graderController struct {
	graderSvc services_grader.GraderService
}

func (gc *graderController) GenerateOutput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GenerateOutputRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Convert string lang to ProgrammingLanguage type
	var lang types.ProgrammingLanguage
	switch req.Lang {
	case "python":
		lang = types.Python
	case "c":
		lang = types.C
	case "cpp":
		lang = types.CPP
	default:
		lang = types.Python // default to Python
	}

	res, err := gc.graderSvc.GenerateOutput(req.Code, lang, req.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func New(graderSvc services_grader.GraderService) GraderController {
	return &graderController{
		graderSvc: graderSvc,
	}
}

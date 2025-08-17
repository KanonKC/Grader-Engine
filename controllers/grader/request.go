package controllers_grader

type GenerateOutputRequest struct {
	Code        string   `json:"code"`
	Lang        string   `json:"lang"`
	Input       []string `json:"input"`
	TimeLimitMs int      `json:"time_limit_ms"`
}

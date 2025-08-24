package services_grader

type ProgrammingLanguage string

const (
	Python ProgrammingLanguage = "python"
	CPP    ProgrammingLanguage = "cpp"
	C      ProgrammingLanguage = "c"
)

type GenerateOutputRequest struct {
	Code        string              `json:"code"`
	Lang        ProgrammingLanguage `json:"lang"`
	Input       []string            `json:"input"`
	TimeLimitMs int                 `json:"time_limit_ms"`
}

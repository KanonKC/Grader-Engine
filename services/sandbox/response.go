package services_sandbox

type WrittenFile struct {
	Filename string
	Content  string
}

type RuntimeOutput struct {
	IsError          bool   `json:"is_error"`
	IsTimeout        bool   `json:"is_timeout"`
	IsMemoryExceeded bool   `json:"is_memory_exceeded"`
	InputIndex       int    `json:"input_index"`
	InputContent     string `json:"input_content"`
	OutputContent    string `json:"output_content"`
	ExecutionTimeMs  int    `json:"execution_time_ms"`
	MemoryUsageKB    int    `json:"memory_usage_kb"`
	Error            string `json:"error"`
}

type RuntimeResult struct {
	IsError          bool            `json:"is_error"`
	IsTimeout        bool            `json:"is_timeout"`
	IsMemoryExceeded bool            `json:"is_memory_exceeded"`
	Output           []RuntimeOutput `json:"output"`
}

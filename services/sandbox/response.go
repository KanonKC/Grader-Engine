package services_sandbox

type WrittenFile struct {
	Filename string
	Content  string
}

type RuntimeOutput struct {
	IsError          bool
	IsTimeout        bool
	IsMemoryExceeded bool
	InputContent     string
	OutputContent    string
	ExecutionTimeMs  int
	MemoryUsageKB    int
}

type RuntimeResult struct {
	IsError          bool
	IsTimeout        bool
	IsMemoryExceeded bool
	Output           []RuntimeOutput
}

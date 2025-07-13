package types

type ProgrammingLanguage string

const (
	Python ProgrammingLanguage = "python"
	CPP    ProgrammingLanguage = "cpp"
	C      ProgrammingLanguage = "c"
)

type WrittenFile struct {
	Filename string
	Content  string
}

type StatusArray int

const (
	Available StatusArray = 0
	Busy      StatusArray = 1
	Error     StatusArray = 2
)

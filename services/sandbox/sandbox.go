package services_sandbox

import (
	"ModelGrader-Grader/types"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

type SandboxService interface {
	Init() error
	FindAvailableSandbox() (int, error)
	ReleaseSandbox(id int) error
	MakeBusy(id int) error
	WriteInput(id int, content string) error
	WriteCode(id int, lang types.ProgrammingLanguage, content string) error
	RunCode(id int, lang types.ProgrammingLanguage) (*RuntimeResult, error)
	RunCodePython(id int) (*RuntimeResult, error)
}

type sandboxService struct {
	size        int
	statusArray []types.StatusArray
}

func (s *sandboxService) Init() error {
	s.statusArray = make([]types.StatusArray, s.size)
	for i := range s.statusArray {
		s.statusArray[i] = types.Available
		err := os.MkdirAll(fmt.Sprintf("./tmp/sandbox/%d", i), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sandboxService) FindAvailableSandbox() (int, error) {
	for i, status := range s.statusArray {
		if status == types.Available {
			return i, nil
		}
	}
	return -1, nil
}

func (s *sandboxService) ReleaseSandbox(id int) error {
	if s.statusArray[id] != types.Busy {
		return errors.New("sandbox is not busy")
	}

	inputsDir := fmt.Sprintf("./tmp/sandbox/%d/inputs", id)
	if err := os.RemoveAll(inputsDir); err != nil {
		return errors.New("failed to delete inputs directory")
	}

	s.statusArray[id] = types.Available
	return nil
}

func (s *sandboxService) MakeBusy(id int) error {
	if s.statusArray[id] != types.Available {
		return errors.New("sandbox is not available")
	}
	s.statusArray[id] = types.Busy
	return nil
}

func (s *sandboxService) WriteInput(id int, content string) error {

	filename := uuid.New().String()
	filename = strings.ReplaceAll(filename, "-", "")

	inputsDir := fmt.Sprintf("./tmp/sandbox/%d/inputs", id)
	if err := os.MkdirAll(inputsDir, 0755); err != nil {
		return fmt.Errorf("failed to create inputs directory: %w", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", inputsDir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)

	return nil
}

func (s *sandboxService) WriteCode(id int, lang types.ProgrammingLanguage, content string) error {
	var fileType string
	switch lang {
	case types.Python:
		fileType = "py"
	case types.C:
		fileType = "c"
	case types.CPP:
		fileType = "cpp"
	}

	if fileType == "" {
		return errors.New("unsupported language")
	}

	filename := "main." + fileType
	file, err := os.Create(fmt.Sprintf("./tmp/sandbox/%d/%s", id, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(content)

	return nil
}

func (s *sandboxService) RunCode(id int, lang types.ProgrammingLanguage) (*RuntimeResult, error) {
	switch lang {
	case types.Python:
		return s.RunCodePython(id)
	case types.C, types.CPP:
		return nil, errors.New("C/C++ execution not implemented yet")
	default:
		return nil, errors.New("unsupported language")
	}
}

func (s *sandboxService) RunCodePython(id int) (*RuntimeResult, error) {
	// Check if inputs directory exists
	inputsDir := fmt.Sprintf("./tmp/sandbox/%d/inputs", id)
	if _, err := os.Stat(inputsDir); os.IsNotExist(err) {
		// If inputs directory doesn't exist, return empty result
		return &RuntimeResult{
			IsError:          false,
			IsTimeout:        false,
			IsMemoryExceeded: false,
			Output:           []RuntimeOutput{},
		}, nil
	}

	// Find all files in the sandbox directory
	fmt.Println("inputsDir", inputsDir)
	inputFiles, err := os.ReadDir(inputsDir)
	if err != nil {
		fmt.Println("Fail 4.1")
		return nil, err
	}

	var runtimeOutputs []RuntimeOutput
	for _, input := range inputFiles {
		// Execute the Python file in the sandbox directory
		cmd := exec.Command("python3", "main.py")
		cmd.Dir = fmt.Sprintf("./tmp/sandbox/%d", id)

		inputData, err := os.ReadFile(fmt.Sprintf("./tmp/sandbox/%d/inputs/"+input.Name(), id))
		if err != nil {
			fmt.Println("Fail 4.2")
			return nil, err
		}

		// Set input data as stdin
		cmd.Stdin = strings.NewReader(string(inputData))

		// Get stdout only (not stderr)
		stdout, err := cmd.Output()

		runtimeOutput := &RuntimeOutput{
			IsError:          err != nil,
			IsTimeout:        false,
			IsMemoryExceeded: false,
			InputContent:     string(inputData), // TODO: read from input file
			OutputContent:    string(stdout),
			ExecutionTimeMs:  0, // TODO: measure execution time
			MemoryUsageKB:    0, // TODO: measure memory usage
		}
		runtimeOutputs = append(runtimeOutputs, *runtimeOutput)
	}

	return &RuntimeResult{
		IsError:          err != nil,
		IsTimeout:        false,
		IsMemoryExceeded: false,
		Output:           runtimeOutputs,
	}, nil
}

func New(size int) SandboxService {
	return &sandboxService{size: size}
}

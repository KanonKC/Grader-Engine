package services_grader

import (
	services_sandbox "ModelGrader-Grader/services/sandbox"
	"ModelGrader-Grader/types"
	"fmt"
)

type GraderService interface {
	GenerateOutput(code string, lang types.ProgrammingLanguage, inputFile []string) (*services_sandbox.RuntimeResult, error)
}

type graderService struct {
	sandboxService services_sandbox.SandboxService
}

func (gs *graderService) GenerateOutput(code string, lang types.ProgrammingLanguage, inputFile []string) (*services_sandbox.RuntimeResult, error) {

	// Step 1: Find available sandbox
	sid := -1
	var err error
	for sid == -1 {
		sid, err = gs.sandboxService.FindAvailableSandbox()
		if err != nil {
			fmt.Println("Fail 1")
			return nil, err
		}
	}

	gs.sandboxService.MakeBusy(sid)

	// Step 2: Write source code to target sandbox
	err = gs.sandboxService.WriteCode(sid, lang, code)
	if err != nil {
		fmt.Println("Fail 2")
		return nil, err
	}

	// Step 3: Write input files to target sandbox
	for index, input := range inputFile {
		err = gs.sandboxService.WriteInput(sid, input, index)
		if err != nil {
			fmt.Println("Fail 3")
			return nil, err
		}
	}

	// Step 4: Run the code
	runtimeResult, err := gs.sandboxService.RunCode(sid, lang)
	if err != nil {
		fmt.Println("Fail 4")
		return nil, err
	}

	defer gs.sandboxService.ReleaseSandbox(sid)

	return runtimeResult, nil
}

func New(sandboxService services_sandbox.SandboxService) GraderService {
	return &graderService{sandboxService: sandboxService}
}

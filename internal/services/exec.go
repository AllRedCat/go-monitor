package services

import (
	"os/exec"

	"monitor/internal/models"
)

func RunScript(req models.ExecRequest) (models.ExecResponse, error) {
	cmd := exec.Command("sh", req.Path)
	output, err := cmd.CombinedOutput()

	resp := models.ExecResponse{
		Output: string(output),
	}

	if err != nil {
		resp.Error = err.Error()
	}
	
	return resp, nil
}

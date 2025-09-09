package models

type ExecRequest struct {
	Path string `json:"path"`
}

type ExecResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

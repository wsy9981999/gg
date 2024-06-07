package dto

type FrontendConfig struct {
	Name    string   `json:"name,omitempty"`
	Plugins []string `json:"plugins,omitempty"`
	Env     []string `json:"env,omitempty"`
}

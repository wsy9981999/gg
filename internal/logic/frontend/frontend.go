package frontend

import (
	"gf-generate/internal/service"
)

func init() {
	service.RegisterFrontend(&sFrontend{})
}

type sFrontend struct {
}

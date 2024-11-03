package v2

import (
	"github.com/dappster-io/DappsterOS/codegen"
	"github.com/dappster-io/DappsterOS/service"
)

type DappsterOS struct {
	fileUploadService *service.FileUploadService
}

func NewDappsterOS() codegen.ServerInterface {
	return &DappsterOS{
		fileUploadService: service.NewFileUploadService(),
	}
}

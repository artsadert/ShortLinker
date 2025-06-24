package mapper

import (
	"github.com/artsadert/ShortLinker/internal/application/command"
	"github.com/artsadert/ShortLinker/internal/interface/api/rest/dto/response"
)

func ToCreateLinkResponse(create_command *command.CreateLinkCommandResult) *response.CreateLinkResponse {
	return &response.CreateLinkResponse{Long: create_command.Result.Long, Short: create_command.Result.Short}
}

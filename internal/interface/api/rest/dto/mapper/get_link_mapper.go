package mapper

import (
	"github.com/artsadert/ShortLinker/internal/application/command"
	"github.com/artsadert/ShortLinker/internal/interface/api/rest/dto/response"
)

func ToGetLinkResponse(get_command *command.GetLinkCommmandResult) *response.GetLinkResponse {
	return &response.GetLinkResponse{Long: get_command.Result.Long, Short: get_command.Result.Short}

}

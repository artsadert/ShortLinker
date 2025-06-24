package request

import "github.com/artsadert/ShortLinker/internal/application/command"

type GetLinkRequest struct {
	Link string `json:"link"`
}

func ToGetLinkCommand(request *GetLinkRequest) *command.GetLinkCommand {
	return &command.GetLinkCommand{Short: request.Link}
}

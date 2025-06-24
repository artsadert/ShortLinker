package request

import "github.com/artsadert/ShortLinker/internal/application/command"

type CreateLinkRequest struct {
	Link string `json:"link"`
}

func ToCreateLinkCommand(request *CreateLinkRequest) *command.CreateLinkCommand {
	return &command.CreateLinkCommand{Long: request.Link}
}

package interfaces

import (
	"github.com/artsadert/ShortLinker/internal/application/command"
)

type LinkService interface {
	CreateNewLink(*command.CreateLinkCommand) (*command.CreateLinkCommandResult, error)
	GetLink(*command.GetLinkCommand) (*command.GetLinkCommmandResult, error)
}

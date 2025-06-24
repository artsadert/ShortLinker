package mapper

import (
	"github.com/artsadert/ShortLinker/internal/application/command"
	"github.com/artsadert/ShortLinker/internal/domain/entities"
)

func ToEntityLong(command *command.CreateLinkCommand) *entities.Link {
	return entities.NewLink(command.Long)
}

func ToEntityShort(command *command.GetLinkCommand) *entities.Link {
	return entities.NewLink(command.Short)
}

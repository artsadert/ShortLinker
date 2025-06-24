package service

import (
	"fmt"
	"github.com/artsadert/ShortLinker/internal/application/command"
	"github.com/artsadert/ShortLinker/internal/application/common"
	"github.com/artsadert/ShortLinker/internal/application/interfaces"
	"github.com/artsadert/ShortLinker/internal/application/mapper"
	"github.com/artsadert/ShortLinker/internal/domain/repositories"
)

type LinkService struct {
	linkRepository repositories.LinkRepository
}

func NewLinkService(link_repository repositories.LinkRepository) interfaces.LinkService {
	return &LinkService{linkRepository: link_repository}
}

func (l *LinkService) CreateNewLink(link_command *command.CreateLinkCommand) (*command.CreateLinkCommandResult, error) {
	// business logic required cause we need to move this to entities to make it shorter
	if link_command == nil {
		return nil, fmt.Errorf("command is required")
	}

	entity := mapper.ToEntityLong(link_command)

	short, err := l.linkRepository.Create(entity)
	if err != nil {
		return nil, err
	}

	res := command.CreateLinkCommandResult{Result: &common.LinkResult{Short: short, Long: link_command.Long}}

	return &res, nil
}

func (l *LinkService) GetLink(link_command *command.GetLinkCommand) (*command.GetLinkCommmandResult, error) {
	// thing to find Long link from short link, no business logic required
	if link_command == nil {
		return nil, fmt.Errorf("command is required")
	}

	entity := mapper.ToEntityShort(link_command)

	long, err := l.linkRepository.Get(entity)
	if err != nil {
		return nil, err
	}

	res := command.GetLinkCommmandResult{Result: &common.LinkResult{Short: link_command.Short, Long: long}}

	return &res, nil
}

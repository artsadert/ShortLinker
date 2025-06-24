package repositories

import (
	"github.com/artsadert/ShortLinker/internal/domain/entities"
)

type LinkRepository interface {
	Create(*entities.Link) (string, error)
	Get(*entities.Link) (string, error)
}

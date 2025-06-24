package command

import "github.com/artsadert/ShortLinker/internal/application/common"

type GetLinkCommand struct {
	Short string
}

type GetLinkCommmandResult struct {
	Result *common.LinkResult
}


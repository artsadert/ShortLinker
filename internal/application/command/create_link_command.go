package command

import "github.com/artsadert/ShortLinker/internal/application/common"

// command layer operates an unreal models
type CreateLinkCommand struct {
	Long string
}

type CreateLinkCommandResult struct {
	Result *common.LinkResult
}

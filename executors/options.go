package executors

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/utils"
)

type Options struct {
	Task          *models.Task
	Action        *models.Action
	Configuration *models.ActionConfiguration
	Variables     map[string]string
}

func (r *Options) MergeVariables() map[string]string {
	if r.Configuration == nil {
		return utils.Merge(r.Action.Variables, r.Task.Variables, r.Variables)
	}

	return utils.Merge(r.Action.Variables, r.Task.Variables, r.Configuration.Variables, r.Variables)
}

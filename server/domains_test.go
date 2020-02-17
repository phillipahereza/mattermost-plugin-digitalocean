package main

import "github.com/mattermost/mattermost-server/v5/model"

func (s *DoPluginTestSuite) TestListDomainsCommandFunc() {
	mockPlugin := Plugin{}
	commandArgs := &model.CommandArgs{Command: "/jira settings", UserId: "1"}
	s.Run("Test List Domains when there no available domains", func() {
		response, err :=mockPlugin.listDomainsCommandFunc(s.client, commandArgs)
	})
}

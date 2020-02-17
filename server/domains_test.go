package main

import (
	"context"
	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
)

func (s *DoPluginTestSuite) TestListDomainsCommandFunc() {
	mockPlugin := &Plugin{BotUserID: "1"}
	commandArgs := &model.CommandArgs{Command: "/jira settings", UserId: "1"}
	s.Run("Test List Domains when there no available domains", func() {
		s.client.EXPECT().ListDomains(context.TODO(), nil).Return([]godo.Domain{}, &godo.Response{}, nil).Times(1)
		response, err := mockPlugin.listDomainsCommandFunc(s.client, commandArgs)

		s.Require().Nil(err)
		s.Require().Equal("You don't have any domains configured. Use `/do create-domain  <domainName> <ipAddress[optional]>` to provision one", response.Text)
	})
}

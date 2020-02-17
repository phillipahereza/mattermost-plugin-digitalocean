package main

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/mock"
)

func (s *DoPluginTestSuite) TestListDomainsCommandFunc() {
	p := &Plugin{BotUserID: "1"}
	api := &plugintest.API{}

	commandArgs := &model.CommandArgs{Command: "/jira list-domains", UserId: "1"}

	tests := map[string]struct {
		message         string
		domains         []godo.Domain
		commandResponse *model.CommandResponse
	}{
		"Test List Domains when there no available domains": {
			message:         "You don't have any domains configured. Use `/do create-domain  <domainName> <ipAddress[optional]>` to provision one",
			domains:         []godo.Domain{},
			commandResponse: &model.CommandResponse{},
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			s.client.EXPECT().ListDomains(context.TODO(), nil).Return(tt.domains, &godo.Response{}, nil).Times(1)

			api.On("LogError", mock.AnythingOfTypeArgument("string")).Return(nil)
			api.On("SendEphemeralPost", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Post")).Run(func(args mock.Arguments) {
				post := args.Get(1).(*model.Post)
				s.Assert().Equal(tt.message, post.Message)
			}).Once().Return(&model.Post{})

			p.SetAPI(api)

			response, err := p.listDomainsCommandFunc(s.client, commandArgs)

			s.Require().Nil(err)
			s.Require().Equal(tt.commandResponse.Text, response.Text)
		})

	}
}

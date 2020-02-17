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

	commandArgs := &model.CommandArgs{Command: "/do list-domains", UserId: "1"}

	message := "You don't have any domains configured. Use `/do create-domain  <domainName> <ipAddress[optional]>` to provision one"

	s.Run("Test List Domains when there no available domains", func() {
		s.client.EXPECT().ListDomains(context.TODO(), nil).Return([]godo.Domain{}, &godo.Response{}, nil).Times(1)

		api.On("LogError", mock.AnythingOfTypeArgument("string")).Return(nil)
		api.On("SendEphemeralPost", mock.AnythingOfType("string"), mock.AnythingOfType("*model.Post")).Run(func(args mock.Arguments) {
			post := args.Get(1).(*model.Post)
			s.Assert().Equal(message, post.Message)
		}).Once().Return(&model.Post{})

		p.SetAPI(api)

		_, err := p.listDomainsCommandFunc(s.client, commandArgs)

		s.Require().Nil(err)
	})

}

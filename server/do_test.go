package main

import (
	"github.com/golang/mock/gomock"
	"github.com/phillipahereza/mattermost-plugin-digitalocean/server/mocks"
	"github.com/stretchr/testify/suite"
)

type DoPluginTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	client   *mocks.MockDigitalOceanService
}

func (s *DoPluginTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.client = mocks.NewMockDigitalOceanService(s.mockCtrl)
}

func (s *DoPluginTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}
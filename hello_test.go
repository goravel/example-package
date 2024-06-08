package hello

import (
	"testing"

	configmock "github.com/goravel/framework/mocks/config"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	suite.Suite
	hello      *Hello
	mockConfig *configmock.Config
}

func TestHelloTestSuite(t *testing.T) {
	suite.Run(t, new(HelloTestSuite))
}

func (s *HelloTestSuite) SetupTest() {
	s.mockConfig = &configmock.Config{}
	s.hello = NewHello(s.mockConfig)
}

func (s *HelloTestSuite) TestWorld() {
	s.mockConfig.On("GetString", "hello.name").Return("Package").Once()

	s.Equal("Welcome To Goravel Package", s.hello.World())
	s.mockConfig.AssertExpectations(s.T())
}

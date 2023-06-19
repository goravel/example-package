package hello

import (
	"testing"

	configmocks "github.com/goravel/framework/contracts/config/mocks"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	suite.Suite
	hello      *Hello
	mockConfig *configmocks.Config
}

func TestHelloTestSuite(t *testing.T) {
	suite.Run(t, new(HelloTestSuite))
}

func (s *HelloTestSuite) SetupTest() {
	s.mockConfig = &configmocks.Config{}
	s.hello = NewHello(s.mockConfig)
}

func (s *HelloTestSuite) TestWorld() {
	s.mockConfig.On("GetString", "hello.name").Return("Package").Once()

	s.Equal("Welcome To Goravel Package", s.hello.World())
	s.mockConfig.AssertExpectations(s.T())
}

package fireness

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestIPListSuite(t *testing.T) {
	suite.Run(t, &IPListSuite{})
}

type IPListSuite struct {
	suite.Suite
}

func (suite *IPListSuite) Test_Empty() {
	var list IPList
	assert.True(suite.T(), list.IsEmpty())
	list = append(list, "127.0.0.1")
	assert.False(suite.T(), list.IsEmpty())
}

func (suite *IPListSuite) Test_IsNil() {
	var list IPList
	assert.True(suite.T(), list.IsNil())
}

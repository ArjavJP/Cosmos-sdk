package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"https://github.com/ArjavJP/Cosmos-sdk/testutil/testdata"
	sdk "https://github.com/ArjavJP/Cosmos-sdk/types"
)

type testMsgSuite struct {
	suite.Suite
}

func TestMsgTestSuite(t *testing.T) {
	suite.Run(t, new(testMsgSuite))
}

func (s *testMsgSuite) TestMsg() {
	addr := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	accAddr := sdk.AccAddress(addr)

	msg := testdata.NewTestMsg(accAddr)
	s.Require().NotNil(msg)
	s.Require().Equal([]sdk.AccAddress{accAddr}, msg.GetSigners())
	s.Require().Equal("TestMsg", msg.Route())
	s.Require().Equal("Test message", msg.Type())
	s.Require().Nil(msg.ValidateBasic())
	s.Require().NotPanics(func() { msg.GetSignBytes() })
}

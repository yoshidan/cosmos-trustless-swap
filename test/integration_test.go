package test

import (
	app2 "swap/app"
	"swap/x/swap/keeper"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	fooDenom = "foo"
	barDenom = "bar"
)

type IntegrationTestSuite struct {
	suite.Suite
	app         *app2.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := Setup(suite.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	app.SwapKeeper.SetParams(ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.SwapKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx
	suite.queryClient = queryClient
}

func (suite *IntegrationTestSuite) TestSendSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send1_______________")
	senderBalance := sdk.NewCoins(newFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv1_______________")
	receiverBalance := sdk.NewCoins(newBarCoin(100))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	response, err := server.Send(ctx, &types.MsgSend{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          newFooCoin(10).String(),
		AmountToReceive: newBarCoin(5).String(),
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), response.Id)
}

func newFooCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(fooDenom, amt)
}

func newBarCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(barDenom, amt)
}

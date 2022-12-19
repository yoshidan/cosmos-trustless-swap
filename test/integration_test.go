package test

import (
	"fmt"
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

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          newFooCoin(10).String(),
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.Send(ctx, sendParam)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), response.Id)
	suite.Require().Equal(response.Id, app.SwapKeeper.GetMaxSwapID(ctx))
	swap, found := app.SwapKeeper.GetSwap(ctx, response.Id)
	suite.Require().True(found)
	suite.Require().Equal(types.SwapStatus_Active, swap.Status)
	suite.Require().Equal(sender.String(), swap.Sender)
	suite.Require().Equal(receiver.String(), swap.Receiver)
	suite.Require().Equal(sendParam.Amount, swap.Amount)
	suite.Require().Equal(sendParam.AmountToReceive, swap.AmountToReceive)
	balance, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balance.Balance.Amount.Uint64())
}

func (suite *IntegrationTestSuite) TestCancelSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send2_______________")
	senderBalance := sdk.NewCoins(newFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv2_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          newFooCoin(10).String(),
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.Send(ctx, sendParam)
	cancelParam := &types.MsgCancel{
		Creator: sender.String(),
		Id:      response.Id,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().NoError(err)
	swap, found := app.SwapKeeper.GetSwap(ctx, response.Id)
	suite.Require().True(found)
	suite.Require().Equal(types.SwapStatus_Cancelled, swap.Status)

	// Check the token return
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceSender, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: sender.String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(100), balanceSender.Balance.Amount.Uint64())

	// Already cancelled
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().Equal(fmt.Sprintf("actual = 2, expected = 0: %s", types.ErrInvalidSwapStatus), err.Error())

	_, err = server.Receive(ctx, &types.MsgReceive{
		Creator: receiver.String(),
		Id:      response.Id,
	})
	suite.Require().Equal(fmt.Sprintf("actual = 2, expected = 0: %s", types.ErrInvalidSwapStatus), err.Error())

}

func (suite *IntegrationTestSuite) TestReceiveSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send3_______________")
	senderBalance := sdk.NewCoins(newFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv3_______________")
	receiverBalance := sdk.NewCoins(newBarCoin(100))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          newFooCoin(10).String(),
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.Send(ctx, sendParam)
	receiveParam := &types.MsgReceive{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().NoError(err)
	swap, found := app.SwapKeeper.GetSwap(ctx, response.Id)
	suite.Require().True(found)
	suite.Require().Equal(types.SwapStatus_Closed, swap.Status)

	// Check token swapped
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceReceiver, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: receiver.String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balanceReceiver.Balance.Amount.Uint64())

	balanceSender, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: sender.String(),
		Denom:   barDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSender.Balance.Amount.Uint64())

	// Already received
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().Equal(fmt.Sprintf("actual = 1, expected = 0: %s", types.ErrInvalidSwapStatus), err.Error())

	_, err = server.Cancel(ctx, &types.MsgCancel{
		Creator: sender.String(),
		Id:      response.Id,
	})
	suite.Require().Equal(fmt.Sprintf("actual = 1, expected = 0: %s", types.ErrInvalidSwapStatus), err.Error())
}

func newFooCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(fooDenom, amt)
}

func newBarCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(barDenom, amt)
}

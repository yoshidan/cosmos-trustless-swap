package test

import (
	app2 "swap/app"
	"swap/test/internal"
	"swap/x/swap/keeper"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/types/errors"

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
	app := internal.Setup(suite.T(), false)
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

	// default class
	class := nft.Class{
		Id: "classId",
	}
	suite.Require().NoError(app.NFTKeeper.SaveClass(ctx, class))
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
	queryResponse, err := app.SwapKeeper.Show(ctx, &types.QueryShowRequest{Id: response.Id})
	swap := queryResponse.Swap
	suite.Require().NoError(err)
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

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: sender.String(),
		Id:      response.Id,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetSwap(ctx, response.Id)
	suite.Require().False(found)

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
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.Receive(ctx, &types.MsgReceive{
		Creator: receiver.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
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

	// Receive
	receiveParam := &types.MsgReceive{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetSwap(ctx, response.Id)
	suite.Require().False(found)

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
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.Cancel(ctx, &types.MsgCancel{
		Creator: sender.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *IntegrationTestSuite) TestCancelError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send4_______________")
	senderBalance := sdk.NewCoins(newFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv4_______________")
	receiverBalance := sdk.NewCoins(newBarCoin(1))
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

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)
}

func (suite *IntegrationTestSuite) TestReceiveError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send6_______________")
	senderBalance := sdk.NewCoins(newFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv6_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          newFooCoin(10).String(),
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.Send(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceive{
		Creator: sender.String(),
		Id:      response.Id,
	}
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)

	receiveParam2 := &types.MsgReceive{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.Receive(ctx, receiveParam2)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

func (suite *IntegrationTestSuite) TestSendNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send1_______________")
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft1",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, sender))

	receiver := sdk.AccAddress("recv1_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.SendNFT(ctx, sendParam)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), response.Id)
	suite.Require().Equal(response.Id, app.SwapKeeper.GetMaxNFTSwapID(ctx))
	queryResponse, err := app.SwapKeeper.ShowNFT(ctx, &types.QueryShowNFTRequest{Id: response.Id})
	swap := queryResponse.Swap
	suite.Require().NoError(err)
	suite.Require().Equal(sender.String(), swap.Sender)
	suite.Require().Equal(receiver.String(), swap.Receiver)
	suite.Require().Equal(sendParam.ClassId, swap.ClassId)
	suite.Require().Equal(sendParam.NftId, swap.NftId)
	suite.Require().Equal(sendParam.AmountToReceive, swap.AmountToReceive)
	ownerResponse, err := app.NFTKeeper.Owner(ctx, &nft.QueryOwnerRequest{
		ClassId: item.ClassId,
		Id:      item.Id,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(), ownerResponse.Owner)
}

func (suite *IntegrationTestSuite) TestCancelNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send2_______________")
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft2",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, sender))

	receiver := sdk.AccAddress("recv2_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.SendNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: sender.String(),
		Id:      response.Id,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetNFTSwap(ctx, response.Id)
	suite.Require().False(found)

	// Check the token return
	ownerResponse, err := app.NFTKeeper.Owner(ctx, &nft.QueryOwnerRequest{
		ClassId: item.ClassId,
		Id:      item.Id,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(sender.String(), ownerResponse.Owner)

	// Already cancelled
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.ReceiveNFT(ctx, &types.MsgReceiveNFT{
		Creator: receiver.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *IntegrationTestSuite) TestReceiveNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send3_______________")
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft3",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, sender))

	receiver := sdk.AccAddress("recv3_______________")
	receiverBalance := sdk.NewCoins(newBarCoin(100))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.SendNFT(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceiveNFT{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetNFTSwap(ctx, response.Id)
	suite.Require().False(found)

	// Check token swapped
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   fooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	ownerResponse, err := app.NFTKeeper.Owner(ctx, &nft.QueryOwnerRequest{
		ClassId: item.ClassId,
		Id:      item.Id,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(receiver.String(), ownerResponse.Owner)

	balanceSender, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: sender.String(),
		Denom:   barDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSender.Balance.Amount.Uint64())

	// Already received
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.CancelNFT(ctx, &types.MsgCancelNFT{
		Creator: sender.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *IntegrationTestSuite) TestCancelNFTError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send4_______________")
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft4",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, sender))

	receiver := sdk.AccAddress("recv4_______________")
	receiverBalance := sdk.NewCoins(newBarCoin(1))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.SendNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)
}

func (suite *IntegrationTestSuite) TestReceiveNFTError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send6_______________")
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft6",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, sender))

	receiver := sdk.AccAddress("recv6_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: newBarCoin(5).String(),
	}
	response, err := server.SendNFT(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceiveNFT{
		Creator: sender.String(),
		Id:      response.Id,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)

	receiveParam2 := &types.MsgReceiveNFT{
		Creator: receiver.String(),
		Id:      response.Id,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam2)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

func newFooCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(fooDenom, amt)
}

func newBarCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(barDenom, amt)
}

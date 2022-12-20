package test

import (
	"testing"
	"time"

	app2 "github.com/yoshidan/cosmos-trustless-swap/app"
	"github.com/yoshidan/cosmos-trustless-swap/test/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/keeper"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SwapTestSuite struct {
	suite.Suite
	app         *app2.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestSwapTestSuite(t *testing.T) {
	suite.Run(t, new(SwapTestSuite))
}

func (suite *SwapTestSuite) SetupTest() {
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

func (suite *SwapTestSuite) TestSendSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send1_______________")
	senderBalance := sdk.NewCoins(internal.NewFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv1_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          internal.NewFooCoin(10).String(),
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.Send(ctx, sendParam)
	suite.Require().NoError(err)
	queryResponse, err := app.SwapKeeper.Show(ctx, &types.QueryShowRequest{Id: 1, Sender: sender.String()})
	swap := queryResponse.Swap
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), swap.Id)
	suite.Require().Equal(sender.String(), swap.Creator)
	suite.Require().Equal(receiver.String(), swap.Receiver)
	suite.Require().Equal(sendParam.Amount, swap.Amount)
	suite.Require().Equal(sendParam.AmountToReceive, swap.AmountToReceive)
	balance, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balance.Balance.Amount.Uint64())
}

func (suite *SwapTestSuite) TestCancelSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send2_______________")
	senderBalance := sdk.NewCoins(internal.NewFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv2_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          internal.NewFooCoin(10).String(),
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.Send(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: sender.String(),
		Id:      1,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetSwap(ctx, sender.String(), 1)
	suite.Require().False(found)

	// Check the token return
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceSender, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: sender.String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(100), balanceSender.Balance.Amount.Uint64())

	// Already cancelled
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.Receive(ctx, &types.MsgReceive{
		Creator: receiver.String(),
		Id:      1,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestReceiveSuccess() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send3_______________")
	senderBalance := sdk.NewCoins(internal.NewFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv3_______________")
	receiverBalance := sdk.NewCoins(internal.NewBarCoin(100))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          internal.NewFooCoin(10).String(),
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.Send(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceive{
		Creator: receiver.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetSwap(ctx, sender.String(), 1)
	suite.Require().False(found)

	// Check token swapped
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceReceiver, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: receiver.String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balanceReceiver.Balance.Amount.Uint64())

	balanceSender, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: sender.String(),
		Denom:   internal.BarDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSender.Balance.Amount.Uint64())

	// Already received
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.Cancel(ctx, &types.MsgCancel{
		Creator: sender.String(),
		Id:      1,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestCancelError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send4_______________")
	senderBalance := sdk.NewCoins(internal.NewFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv4_______________")
	receiverBalance := sdk.NewCoins(internal.NewBarCoin(1))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          internal.NewFooCoin(10).String(),
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.Send(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: receiver.String(),
		Id:      1,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestReceiveError() {
	app, ctx := suite.app, suite.ctx

	sender := sdk.AccAddress("send6_______________")
	senderBalance := sdk.NewCoins(internal.NewFooCoin(100))
	senderAccount := app.AccountKeeper.NewAccountWithAddress(ctx, sender)
	app.AccountKeeper.SetAccount(ctx, senderAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, sender, senderBalance))

	receiver := sdk.AccAddress("recv6_______________")

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSend{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		Amount:          internal.NewFooCoin(10).String(),
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.Send(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceive{
		Creator: sender.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.Receive(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)

	receiveParam2 := &types.MsgReceive{
		Creator: receiver.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.Receive(ctx, receiveParam2)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

func (suite *SwapTestSuite) TestSendNFTSuccess() {
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
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.SendNFT(ctx, sendParam)
	suite.Require().NoError(err)
	queryResponse, err := app.SwapKeeper.ShowNFT(ctx, &types.QueryShowNFTRequest{Id: 1, Sender: sender.String()})
	swap := queryResponse.Swap
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), swap.Id)
	suite.Require().Equal(sender.String(), swap.Creator)
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

func (suite *SwapTestSuite) TestCancelNFTSuccess() {
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
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.SendNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: sender.String(),
		Id:      1,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetNFTSwap(ctx, sender.String(), 1)
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
		Id:      1,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestReceiveNFTSuccess() {
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
	receiverBalance := sdk.NewCoins(internal.NewBarCoin(100))
	receiverAccount := app.AccountKeeper.NewAccountWithAddress(ctx, receiver)
	app.AccountKeeper.SetAccount(ctx, receiverAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, receiver, receiverBalance))

	// Send
	server := keeper.NewMsgServerImpl(app.SwapKeeper)
	sendParam := &types.MsgSendNFT{
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.SendNFT(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceiveNFT{
		Creator: receiver.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SwapKeeper.GetNFTSwap(ctx, sender.String(), 1)
	suite.Require().False(found)

	// Check token swapped
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
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
		Denom:   internal.BarDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSender.Balance.Amount.Uint64())

	// Already received
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)

	_, err = server.CancelNFT(ctx, &types.MsgCancelNFT{
		Creator: sender.String(),
		Id:      1,
	})
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestCancelNFTError() {
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
	receiverBalance := sdk.NewCoins(internal.NewBarCoin(1))
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
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.SendNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: receiver.String(),
		Id:      1,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSwapNotFound, err)
}

func (suite *SwapTestSuite) TestReceiveNFTError() {
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
		Id:              1,
		Creator:         sender.String(),
		Receiver:        receiver.String(),
		ClassId:         item.ClassId,
		NftId:           item.Id,
		AmountToReceive: internal.NewBarCoin(5).String(),
	}
	_, err := server.SendNFT(ctx, sendParam)

	// Receive
	receiveParam := &types.MsgReceiveNFT{
		Creator: sender.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)

	receiveParam2 := &types.MsgReceiveNFT{
		Creator: receiver.String(),
		Sender:  sender.String(),
		Id:      1,
	}
	_, err = server.ReceiveNFT(ctx, receiveParam2)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

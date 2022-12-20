package test

import (
	"swap/x/swap/keeper"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"swap/x/swap/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

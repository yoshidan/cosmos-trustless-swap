package test

import (
	"testing"
	"time"

	app2 "github.com/yoshidan/cosmos-trustless-swap/app"
	"github.com/yoshidan/cosmos-trustless-swap/test/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/keeper"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SaleTestSuite struct {
	suite.Suite
	app         *app2.App
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestSaleTestSuite(t *testing.T) {
	suite.Run(t, new(SaleTestSuite))
}

func (suite *SaleTestSuite) SetupTest() {
	app := internal.Setup(suite.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	app.SaleKeeper.SetParams(ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.SaleKeeper)
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

func (suite *SaleTestSuite) TestSellSuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send1_______________")
	sellerBalance := sdk.NewCoins(internal.NewFooCoin(100))
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, seller, sellerBalance))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSell{
		Creator: seller.String(),
		Amount:  internal.NewFooCoin(10).String(),
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.Sell(ctx, sendParam)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), response.Id)
	suite.Require().Equal(response.Id, app.SaleKeeper.GetMaxSaleID(ctx))
	queryResponse, err := app.SaleKeeper.Show(ctx, &types.QueryShowRequest{Id: response.Id})
	swap := queryResponse.Sale
	suite.Require().NoError(err)
	suite.Require().Equal(seller.String(), swap.Seller)
	suite.Require().Equal(sendParam.Amount, swap.Amount)
	suite.Require().Equal(sendParam.Price, swap.Price)
	balance, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balance.Balance.Amount.Uint64())
}

func (suite *SaleTestSuite) TestCancelSuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send2_______________")
	sellerBalance := sdk.NewCoins(internal.NewFooCoin(100))
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, seller, sellerBalance))

	buyer := sdk.AccAddress("recv2_______________")

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSell{
		Creator: seller.String(),
		Amount:  internal.NewFooCoin(10).String(),
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.Sell(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: seller.String(),
		Id:      response.Id,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SaleKeeper.GetSale(ctx, response.Id)
	suite.Require().False(found)

	// Check the token return
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceSeller, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: seller.String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(100), balanceSeller.Balance.Amount.Uint64())

	// Already cancelled
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)

	_, err = server.Buy(ctx, &types.MsgBuy{
		Creator: buyer.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)
}

func (suite *SaleTestSuite) TestBuySuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send3_______________")
	sellerBalance := sdk.NewCoins(internal.NewFooCoin(100))
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, seller, sellerBalance))

	buyer := sdk.AccAddress("recv3_______________")
	buyerBalance := sdk.NewCoins(internal.NewBarCoin(100))
	buyerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, buyer)
	app.AccountKeeper.SetAccount(ctx, buyerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, buyer, buyerBalance))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSell{
		Creator: seller.String(),
		Amount:  internal.NewFooCoin(10).String(),
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.Sell(ctx, sendParam)

	// Buy
	receiveParam := &types.MsgBuy{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.Buy(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SaleKeeper.GetSale(ctx, response.Id)
	suite.Require().False(found)

	// Check token swapped
	balanceModule, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), balanceModule.Balance.Amount.Uint64())

	balanceBuyer, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: buyer.String(),
		Denom:   internal.FooDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(10), balanceBuyer.Balance.Amount.Uint64())

	balanceSeller, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: seller.String(),
		Denom:   internal.BarDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSeller.Balance.Amount.Uint64())

	// Already received
	_, err = server.Buy(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)

	_, err = server.Cancel(ctx, &types.MsgCancel{
		Creator: seller.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)
}

func (suite *SaleTestSuite) TestCancelError() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send4_______________")
	sellerBalance := sdk.NewCoins(internal.NewFooCoin(100))
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, seller, sellerBalance))

	buyer := sdk.AccAddress("recv4_______________")
	buyerBalance := sdk.NewCoins(internal.NewBarCoin(1))
	buyerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, buyer)
	app.AccountKeeper.SetAccount(ctx, buyerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, buyer, buyerBalance))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSell{
		Creator: seller.String(),
		Amount:  internal.NewFooCoin(10).String(),
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.Sell(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancel{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.Cancel(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)
}

func (suite *SaleTestSuite) TestBuyError() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send6_______________")
	sellerBalance := sdk.NewCoins(internal.NewFooCoin(100))
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, seller, sellerBalance))

	buyer := sdk.AccAddress("recv6_______________")

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSell{
		Creator: seller.String(),
		Amount:  internal.NewFooCoin(10).String(),
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.Sell(ctx, sendParam)

	// Buy
	receiveParam := &types.MsgBuy{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.Buy(ctx, receiveParam)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

func (suite *SaleTestSuite) TestSellNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send1_______________")
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft1",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, seller))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSellNFT{
		Creator: seller.String(),
		ClassId: item.ClassId,
		NftId:   item.Id,
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.SellNFT(ctx, sendParam)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), response.Id)
	suite.Require().Equal(response.Id, app.SaleKeeper.GetMaxNFTSaleID(ctx))
	queryResponse, err := app.SaleKeeper.ShowNFT(ctx, &types.QueryShowNFTRequest{Id: response.Id})
	swap := queryResponse.Sale
	suite.Require().NoError(err)
	suite.Require().Equal(seller.String(), swap.Seller)
	suite.Require().Equal(sendParam.ClassId, swap.ClassId)
	suite.Require().Equal(sendParam.NftId, swap.NftId)
	suite.Require().Equal(sendParam.Price, swap.Price)
	ownerResponse, err := app.NFTKeeper.Owner(ctx, &nft.QueryOwnerRequest{
		ClassId: item.ClassId,
		Id:      item.Id,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(app.AccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String(), ownerResponse.Owner)
}

func (suite *SaleTestSuite) TestCancelNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send2_______________")
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft2",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, seller))

	buyer := sdk.AccAddress("recv2_______________")

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSellNFT{
		Creator: seller.String(),
		ClassId: item.ClassId,
		NftId:   item.Id,
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.SellNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: seller.String(),
		Id:      response.Id,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().NoError(err)
	_, found := app.SaleKeeper.GetNFTSale(ctx, response.Id)
	suite.Require().False(found)

	// Check the token return
	ownerResponse, err := app.NFTKeeper.Owner(ctx, &nft.QueryOwnerRequest{
		ClassId: item.ClassId,
		Id:      item.Id,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(seller.String(), ownerResponse.Owner)

	// Already cancelled
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)

	_, err = server.BuyNFT(ctx, &types.MsgBuyNFT{
		Creator: buyer.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)
}

func (suite *SaleTestSuite) TestBuyNFTSuccess() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send3_______________")
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft3",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, seller))

	buyer := sdk.AccAddress("recv3_______________")
	buyerBalance := sdk.NewCoins(internal.NewBarCoin(100))
	buyerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, buyer)
	app.AccountKeeper.SetAccount(ctx, buyerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, buyer, buyerBalance))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSellNFT{
		Creator: seller.String(),
		ClassId: item.ClassId,
		NftId:   item.Id,
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.SellNFT(ctx, sendParam)

	// Buy
	receiveParam := &types.MsgBuyNFT{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.BuyNFT(ctx, receiveParam)
	suite.Require().NoError(err)
	_, found := app.SaleKeeper.GetNFTSale(ctx, response.Id)
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
	suite.Require().Equal(buyer.String(), ownerResponse.Owner)

	balanceSeller, err := app.BankKeeper.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: seller.String(),
		Denom:   internal.BarDenom,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(5), balanceSeller.Balance.Amount.Uint64())

	// Already received
	_, err = server.BuyNFT(ctx, receiveParam)
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)

	_, err = server.CancelNFT(ctx, &types.MsgCancelNFT{
		Creator: seller.String(),
		Id:      response.Id,
	})
	suite.Require().ErrorIs(types.ErrSaleNotFound, err)
}

func (suite *SaleTestSuite) TestCancelNFTError() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send4_______________")
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft4",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, seller))

	buyer := sdk.AccAddress("recv4_______________")
	buyerBalance := sdk.NewCoins(internal.NewBarCoin(1))
	buyerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, buyer)
	app.AccountKeeper.SetAccount(ctx, buyerAccount)
	suite.Require().NoError(testutil.FundAccount(app.BankKeeper, ctx, buyer, buyerBalance))

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSellNFT{
		Creator: seller.String(),
		ClassId: item.ClassId,
		NftId:   item.Id,
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.SellNFT(ctx, sendParam)

	// Cancel
	cancelParam := &types.MsgCancelNFT{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.CancelNFT(ctx, cancelParam)
	suite.Require().ErrorIs(types.ErrInsufficientPermission, err)
}

func (suite *SaleTestSuite) TestBuyNFTError() {
	app, ctx := suite.app, suite.ctx

	seller := sdk.AccAddress("send6_______________")
	sellerAccount := app.AccountKeeper.NewAccountWithAddress(ctx, seller)
	app.AccountKeeper.SetAccount(ctx, sellerAccount)
	item := nft.NFT{
		ClassId: "classId",
		Id:      "nft6",
	}
	suite.Require().NoError(app.NFTKeeper.Mint(ctx, item, seller))

	buyer := sdk.AccAddress("recv6_______________")

	// Sell
	server := keeper.NewMsgServerImpl(app.SaleKeeper)
	sendParam := &types.MsgSellNFT{
		Creator: seller.String(),
		ClassId: item.ClassId,
		NftId:   item.Id,
		Price:   internal.NewBarCoin(5).String(),
	}
	response, err := server.SellNFT(ctx, sendParam)

	// Buy
	receiveParam := &types.MsgBuyNFT{
		Creator: buyer.String(),
		Id:      response.Id,
	}
	_, err = server.BuyNFT(ctx, receiveParam)
	suite.Require().ErrorIs(errors.ErrInsufficientFunds, err)
}

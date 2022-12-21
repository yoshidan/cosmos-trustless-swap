# cosmos-trustless-swap

**cosmos-trustless-swap** is a module for swapping and selling tokens between two users.  
It is implemented as [Cosmos module](https://github.com/cosmos/cosmos-sdk).

## Features

* Swap Funsible Token
* Swap Non Funsible Token
* Sell Funsible Token
* Sell Non Funsible Token

### Command List

## Installation

```
go get github.com/yoshidan/cosmos-trustless-swap
```

Modify `app/app.go` included in the blockchain boilerplate output by [Ignite CLI](https://ignite.com/cli) as follows.  

* Import `nft` and `swap` and `sale` modules.
* `nft` is not required if you do not swap or sell NFTs.
```go
import (
	 ...
    "github.com/cosmos/cosmos-sdk/x/nft"
    nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
    nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"

    swapmodule "github.com/yoshidan/cosmos-trustless-swap/x/swap"
    swapmodulekeeper "github.com/yoshidan/cosmos-trustless-swap/x/swap/keeper"
    swapmoduletypes "github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

    salemodule "github.com/yoshidan/cosmos-trustless-swap/x/sale"
    salemodulekeeper "github.com/yoshidan/cosmos-trustless-swap/x/sale/keeper"
    salemoduletypes "github.com/yoshidan/cosmos-trustless-swap/x/sale/types"	
)
```

* Add these modules to `ModuleBasics` and `maccPerms`
```go
var (
    ModuleBasics = module.NewBasicManager(
        ...
        nftmodule.AppModuleBasic{},
        swapmodule.AppModuleBasic{},
        salemodule.AppModuleBasic{}, 
    )

    maccPerms = map[string][]string{
        ...
        nft.ModuleName:                 nil,
        swapmoduletypes.ModuleName:     nil,
        salemoduletypes.ModuleName:     nil,
    }
)
```

* Add the `keepers` to `App` struct 
```go
type App struct {
	...
	NFTKeeper  nftkeeper.Keeper
	SwapKeeper swapmodulekeeper.Keeper
	SaleKeeper salemodulekeeper.Keeper
}
```

* Initialize the modules in `func New`
```go
func New(
    ...
) *App {
    ...
    keys := sdk.NewKVStoreKeys(
        ...
        nft.StoreKey,
        swapmoduletypes.StoreKey,
        salemoduletypes.StoreKey,
    )
    ...
	
    app.NFTKeeper = nftkeeper.NewKeeper(
        keys[nft.StoreKey],
        appCodec,
        app.AccountKeeper,
        app.BankKeeper,
    )
    nftModules := nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry)

    app.SwapKeeper = *swapmodulekeeper.NewKeeper(
        appCodec,
        keys[swapmoduletypes.StoreKey],
        keys[swapmoduletypes.MemStoreKey],
        app.GetSubspace(swapmoduletypes.ModuleName),
        app.AccountKeeper,
        app.BankKeeper,
        app.NFTKeeper,
    )
    swapModule := swapmodule.NewAppModule(appCodec, app.SwapKeeper, app.AccountKeeper, app.BankKeeper)

    app.SaleKeeper = *salemodulekeeper.NewKeeper(
        appCodec,
        keys[salemoduletypes.StoreKey],
        keys[salemoduletypes.MemStoreKey],
        app.GetSubspace(salemoduletypes.ModuleName),
        app.BankKeeper,
        app.AccountKeeper,
        app.NFTKeeper,
    )
    saleModule := salemodule.NewAppModule(appCodec, app.SaleKeeper, app.AccountKeeper, app.BankKeeper)

    app.mm = module.NewManager(
        ...
        nftModules,
        swapModule,
        saleModule,
    )

    app.mm.SetOrderBeginBlockers(
        ...
        nft.ModuleName,
        swapmoduletypes.ModuleName,
        salemoduletypes.ModuleName,
    )

    app.mm.SetOrderEndBlockers(
        ...
        nft.ModuleName,
        swapmoduletypes.ModuleName,
        salemoduletypes.ModuleName,
    )
	
    app.mm.SetOrderInitGenesis(
        ...
        nft.ModuleName, 
        swapmoduletypes.ModuleName, 
        salemoduletypes.ModuleName, 
    )

    app.sm = module.NewSimulationManager(
        ...
        nftModules,
        swapModule,
        saleModule, 
    )
    ...
    return app
}

```
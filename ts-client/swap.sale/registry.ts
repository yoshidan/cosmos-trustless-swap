import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSell } from "./types/swap/sale/tx";
import { MsgBuy } from "./types/swap/sale/tx";
import { MsgCancelNFT } from "./types/swap/sale/tx";
import { MsgSellNFT } from "./types/swap/sale/tx";
import { MsgBuyNFT } from "./types/swap/sale/tx";
import { MsgCancel } from "./types/swap/sale/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/swap.sale.MsgSell", MsgSell],
    ["/swap.sale.MsgBuy", MsgBuy],
    ["/swap.sale.MsgCancelNFT", MsgCancelNFT],
    ["/swap.sale.MsgSellNFT", MsgSellNFT],
    ["/swap.sale.MsgBuyNFT", MsgBuyNFT],
    ["/swap.sale.MsgCancel", MsgCancel],
    
];

export { msgTypes }
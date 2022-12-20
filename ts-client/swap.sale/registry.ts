import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBuyNFT } from "./types/swap/sale/tx";
import { MsgBuy } from "./types/swap/sale/tx";
import { MsgSell } from "./types/swap/sale/tx";
import { MsgCancel } from "./types/swap/sale/tx";
import { MsgCancelNFT } from "./types/swap/sale/tx";
import { MsgSellNFT } from "./types/swap/sale/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/swap.sale.MsgBuyNFT", MsgBuyNFT],
    ["/swap.sale.MsgBuy", MsgBuy],
    ["/swap.sale.MsgSell", MsgSell],
    ["/swap.sale.MsgCancel", MsgCancel],
    ["/swap.sale.MsgCancelNFT", MsgCancelNFT],
    ["/swap.sale.MsgSellNFT", MsgSellNFT],
    
];

export { msgTypes }
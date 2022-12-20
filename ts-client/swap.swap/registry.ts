import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCancel } from "./types/swap/swap/tx";
import { MsgCancelNFT } from "./types/swap/swap/tx";
import { MsgReceive } from "./types/swap/swap/tx";
import { MsgReceiveNFT } from "./types/swap/swap/tx";
import { MsgSend } from "./types/swap/swap/tx";
import { MsgSendNFT } from "./types/swap/swap/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/swap.swap.MsgCancel", MsgCancel],
    ["/swap.swap.MsgCancelNFT", MsgCancelNFT],
    ["/swap.swap.MsgReceive", MsgReceive],
    ["/swap.swap.MsgReceiveNFT", MsgReceiveNFT],
    ["/swap.swap.MsgSend", MsgSend],
    ["/swap.swap.MsgSendNFT", MsgSendNFT],
    
];

export { msgTypes }
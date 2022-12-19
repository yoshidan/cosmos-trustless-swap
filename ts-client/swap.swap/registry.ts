import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSend } from "./types/swap/swap/tx";
import { MsgReceive } from "./types/swap/swap/tx";
import { MsgCancel } from "./types/swap/swap/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/swap.swap.MsgSend", MsgSend],
    ["/swap.swap.MsgReceive", MsgReceive],
    ["/swap.swap.MsgCancel", MsgCancel],
    
];

export { msgTypes }
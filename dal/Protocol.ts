import { encode } from '@msgpack/msgpack';

export interface Method {
    method: string;
    args: any;
}

export interface Request {
    id: number;
    db: string;
    commands: Method[];
}

export const METHODS = "In|Find|Select|Fields|Join|Group|Sort|Limit|Offset|Delete|Insert|Set|Update|OnConflict|DoUpdate|DoNothing".split("|");

export function encodeRequest(request: Request): Uint8Array {
    return encode(request);
}
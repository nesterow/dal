import type { Request } from "./Protocol";
import { METHODS } from "./Protocol";

type Primitive = string | number | boolean | null;

interface Filter extends Record<string, unknown> {
    $eq?: Primitive;
    $ne?: Primitive;
    $gt?: Primitive;
    $gte?: Primitive;
    $lt?: Primitive;
    $lte?: Primitive;
    $in?: Primitive[];
    $nin?: Primitive[];
    $like?: string;
    $nlike?: string;
    $glob?: string;
    $between?: [Primitive, Primitive];
    $nbetween?: [Primitive, Primitive];
}
interface FindFilter {
    [key: string]: Primitive | Filter | Filter[] | undefined;
}

type JoinCondition = "inner" | "left" | "cross" | "full outer";
type JoinFilter = {
  $for: string;
  $do: FindFilter;
  $as?: JoinCondition;
};

export type SortOptions = Record<string, 1 | -1 | "asc" | "desc">;


export default class Builder {
    private request: Request;
    constructor(database: string) {
        this.request = {
            id: 0,
            db: database,
            commands: [],
        };
    }
    private format(): void {
        this.request.commands = METHODS.map((method) => {
            const command = this.request.commands.find((command) => command.method === method);
            return command;
        }).filter(Boolean) as Request["commands"];
    }
    In(table: string): Builder {
        this.request.commands.push({ method: "In", args: [table] });
        return this;
    }
    Find(filter: FindFilter): Builder {
        this.request.commands.push({ method: "Find", args: [filter] });
        return this;
    }
    Select(fields: string[]): Builder {
        this.request.commands.push({ method: "Select", args: fields });
        return this;
    }
    Fields(fields: string[]): Builder {
        this.Select(fields);
        return this;
    }
    Join(...joins: JoinFilter[]): Builder {
        this.request.commands.push({ method: "Join", args: joins });
        return this;
    }
    Group(fields: string[]): Builder {
        this.request.commands.push({ method: "Group", args: fields });
        return this;
    }
    Sort(fields: SortOptions): Builder {
        this.request.commands.push({ method: "Sort", args: fields });
        return this;
    }
    Limit(limit: number): Builder {
        this.request.commands.push({ method: "Limit", args: [limit] });
        return this;
    }
    Offset(offset: number): Builder {
        this.request.commands.push({ method: "Offset", args: [offset] });
        return this;
    }
    Delete(): Builder {
        this.request.commands.push({ method: "Delete", args: [] });
        return this;
    }
    Insert(data: Record<string, unknown>): Builder {
        this.request.commands.push({ method: "Insert", args: [data] });
        return this;
    }
    Set(data: Record<string, unknown>): Builder {
        this.request.commands.push({ method: "Set", args: [data] });
        return this;
    }
    Update(data: Record<string, unknown>): Builder {
        this.Set(data);
        return this;
    }
    OnConflict(...fields: string[]): Builder {
        this.request.commands.push({ method: "OnConflict", args: fields });
        return this;
    }
    DoUpdate(...fields: string[]): Builder {
        this.request.commands.push({ method: "DoUpdate", args: fields });
        return this;
    }
    DoNothing(): Builder {
        this.request.commands.push({ method: "DoNothing", args: [] });
        return this;
    }
    
}
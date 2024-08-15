import type { Request } from "./Protocol";
import { METHODS, encodeRequest, decodeRowsIterator } from "./Protocol";

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

type SortOptions = Record<string, 1 | -1 | "asc" | "desc">;

type Options = {
    database: string;
    url: string;
};


export default class Builder <
I extends abstract new (...args: any) => any,
>{
    private request: Request;
    private url: string;
    private dtoTemplate: new (...args: any) => any = Object;
    private methodCalls: Map<string, unknown[]> = new Map(); // one call per method
    private headerRow: unknown[] | null = null;
    constructor(opts: Options) {
        this.request = {
            id: 0,
            db: opts.database,
            commands: [],
        };
        this.url = opts.url;
    }
    private formatRequest(): void {
        this.request.commands = []
        METHODS.forEach((method) => {
            const args = this.methodCalls.get(method);
            if (!args) {
                return;
            }
            this.request.commands.push({ method, args });
        })
    }
    private formatRow(data: unknown[]){
        if (!this.dtoTemplate) {
            return data;
        }
        const instance = new this.dtoTemplate(data);
        for (const idx in this.headerRow!) {
            const header = this.headerRow[idx] as string;
            if (header in instance) {
                instance[header] = data[idx];
            }
        }
        return instance;
    }
    In(table: string): Builder<I> {
        this.methodCalls.set("In", [table]);
        return this;
    }
    Find(filter: FindFilter): Builder<I> {
        this.methodCalls.set("Find", [filter]);
        return this;
    }
    Select(fields: string[]): Builder<I> {
        this.methodCalls.set("Select", fields);
        return this;
    }
    Fields(fields: string[]): Builder<I> {
        this.Select(fields);
        return this;
    }
    Join(...joins: JoinFilter[]): Builder<I> {
        this.methodCalls.set("Join", joins);
        return this;
    }
    Group(fields: string[]): Builder<I> {
        this.methodCalls.set("Group", fields);
        return this;
    }
    Sort(fields: SortOptions): Builder<I> {
        this.methodCalls.set("Sort", [fields]);
        return this;
    }
    Limit(limit: number): Builder<I> {
        this.methodCalls.set("Limit", [limit]);
        return this;
    }
    Offset(offset: number): Builder<I> {
        this.methodCalls.set("Offset", [offset]);
        return this;
    }
    Delete(): Builder<I> {
        this.methodCalls.set("Delete", []);
        return this;
    }
    Insert(data: Record<string, unknown>): Builder<I> {
        this.methodCalls.set("Insert", [data]);
        return this;
    }
    Set(data: Record<string, unknown>): Builder<I> {
        this.methodCalls.set("Set", [data]);
        return this;
    }
    Update(data: Record<string, unknown>): Builder<I> {
        this.Set(data);
        return this;
    }
    OnConflict(...fields: string[]): Builder<I> {
        this.methodCalls.set("OnConflict", fields);
        return this;
    }
    DoUpdate(...fields: string[]): Builder<I> {
        this.methodCalls.delete("DoNothing");
        this.methodCalls.set("DoUpdate", fields);
        return this;
    }
    DoNothing(): Builder<I> {
        this.methodCalls.delete("DoUpdate");
        this.methodCalls.set("DoNothing", []);
        return this;
    }
    async *Rows<T = InstanceType<I>>(): AsyncGenerator<T> {
        this.formatRequest();
        const response = await fetch(this.url, {
            method: "POST",
            body: new Blob([encodeRequest(this.request)]),
            headers: {
                "Content-Type": "application/x-msgpack",
            },
        });
        if (response.status !== 200) {
            throw new Error(await response.text());
        }

        const iterator = decodeRowsIterator(response.body!);
        for await (const row of iterator) {
            if (this.headerRow === null) {
                this.headerRow = row.r;
                await iterator.next();
                continue;
            }
            yield this.formatRow(row.r);
        }
    }
    As<T extends new (...args: any) => any>(template: T): Builder<T> {
        this.dtoTemplate = template;
        return this;
    }
    async Query<T = InstanceType<I>>(): Promise<T[]> {
        const rows = this.Rows();
        const result = [];
        for await (const row of rows) {
            result.push(row);
        }
        return result
    }
    
}
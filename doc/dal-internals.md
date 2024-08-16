# DAL Internal Architecture

- The Client is written in TypeScript.
- The DAL server written in Golang. 

## Components

(Top to bottom)

- NodeJS Client
- Protocol
- Builder
- DB Adapter

## NodeJs Client

Client consists of a query builder and protocol decoder/encoder.

- Query Builder is a light builder which constructs the query object for the server.
- Protocol is a decoder/encoder that utilizes messagepack.

```bash
------------------
|- dal
    |- Builder.ts
    |- Protocol.ts
    |_ index.ts
|...
```

## Protocol


Protocol utilizes messagepack for encoding and decoding the messages.

There following types of encoded data:
- Row stream
- Query (request)
- Response (exec result)

Locations:

```bash
------------------
|- dal
    |- Protocol.ts
|_...
|- pkg
    |- proto
        |...
|...
```

### Row Stream

- The server sends streaming (chunked) data to the client, every chunk is a row.
- Every row starts with a 3-byte header `{0x81, 0xa1, 0x72}` 
- The first row is the header row, which contains the column names.

Parsing the row stream (pseudo code):

```python
header = [0x81, 0xa1, 0x72]
input: byte[] = ...
buffer: byte[] = []
output: byte[][] = []
while i < input.length:
    if input[i] != 0x81:
        buffer << input[i]
        i += 1
    else if input[i:3] == header:
        output << header + buffer
        buffer = []
        i += 3
output << header + buffer
```

MessagePack schema for the row stream:
```go
type Row struct {
	Data []interface{} `msg:"r"`
}
// { r: [] }
```

### Query

- The client utilizes a "light builder" which prepares a list of callbacks for the SQL query builder.
- The Query consits of the following fields:
  - Id: uint32 (optional)
  - Db: string (required, database name)
  - Commands: []BuilderMethod (required, list of Builder arguments)



```go
type BuilderMethod struct {
	Method string        `msg:"method"`
	Args   []interface{} `msg:"args"`
}

type Request struct {
	Id       uint32          `msg:"id"`
	Db       string          `msg:"db"`
	Commands []BuilderMethod `msg:"commands"`
}
```

### Response
The response is inteneded for operation results that don't return rows.

```go
type Response struct {
	Id           uint32 `msg:"i"`
	RowsAffected int64  `msg:"ra"`
	LastInsertId int64  `msg:"li"`
}
```

## Builder

The builder is a set of methods that are used to construct the SQL query.

- The sql query is constructed by the server.
- The client utilizes a "light builder" which prepares a list of callbacks for the server builer.

Locations:

```bash
------------------
|- dal
    |- Builder.ts
|_...
|- pkg
    |- builder
        |...
    |- filters
        |...
|...
```

### Builder Methods
In|Find|Select|Fields|Join|Group|Sort|Limit|Offset|Delete|Insert|Set|Update|OnConflict|DoUpdate|DoNothing
[TS Docs]()
[Golang Docs]()

## DB Adapter

- Adapter provides the interface for the database driver.
- Adapter package also provides utilitities for specific SQL Dialects.

Locations:

```bash
------------------
|- pkg
    |- Adapter
        |...
|...
```
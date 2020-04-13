# Gosura

Go client for the Hasura API.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Gosura](#gosura)
    - [Implemented query types](#implemented-query-types)
        - [Schema/Metadata API](#schemametadata-api)
        - [Data API](#data-api)
        - [PG Dump API](#pg-dump-api)
        - [Config API](#config-api)
    - [Installation](#installation)
    - [Usage](#usage)
        - [Single query](#single-query)
        - [Permissions DSL](#permissions-dsl)
        - [Bulk queries](#bulk-queries)
        - [Data API](#data-api-1)
    - [Changelog](#changelog)
    - [Contributing](#contributing)

<!-- markdown-toc end -->

## Implemented query types

### Schema/Metadata API

- [x] Bulk
- [x] Run SQL
- [x] Tables/Views
  - [x] track\_table
  - [x] set\_table\_is\_enum
  - [x] track\_table v2
  - [x] set\_table\_custom\_fields
  - [x] untrack\_table
- [x] Custom Functions
  - [x] track\_function
  - [x] untrack\_function
- [x] Relationships
  - [x] create\_object\_relationship
  - [x] create\_array\_relationship
  - [x] drop\_relationship
  - [x] set\_relationship\_comment
- [x] Permissions
  - [x] create\_insert\_permission
  - [x] drop\_insert\_permission
  - [x] create\_select\_permission
  - [x] drop\_select\_permission
  - [x] create\_update\_permission
  - [x] drop\_update\_permission
  - [x] create\_delete\_permission
  - [x] drop\_delete\_permission
  - [x] set\_permission\_comment
- [x] Computed Fields
  - [x] add\_computed\_field
  - [x] drop\_computed\_field
- [x] Event Triggers
  - [x] create\_event\_trigger
  - [x] delete\_event\_trigger
  - [x] invoke\_event\_trigger
- [x] Remote Schemas
  - [x] add\_remote\_schema
  - [x] remove\_remote\_schema
  - [x] reload\_remote\_schema
- [x] Query Collections
  - [x] create\_query\_collection
  - [x] drop\_query\_collection
  - [x] add\_query\_to\_collection
    - [x] add\_collection\_to\_allowlist
  - [x] drop\_collection\_from\_allowlist
- [x] Manage Metadata
  - [x] export\_metadata
  - [x] replace\_metadata
  - [x] reload\_metadata
  - [x] clear\_metadata
  - [x] get\_inconsistent\_metadata
  - [x] drop\_inconsistent\_metadata

### Data API

- [x] select
- [ ] insert
- [ ] update
- [ ] delete

### PG Dump API

- [ ] pg_dump

### Config API

- [x] config

## Installation

```sh
go get github.com/wgnet/gosura/gosura
```

## Usage

### Single query

```go
client := NewHasuraClient().
    URL("https://hasura.example.com").
    Endpoint("/location_behind_proxy/v1/query").
    SetAdminSecret("super_secret").
    SkipTLSVerify(true)

query := NewRunSqlQuery()
if err := query.SetArgs(RunSqlArgs{
    SQL: "select count(*) from public.posts",
}); err != nil {
    panic(err)
}

data, err := client.Do(query)
if err != nil {
    log.Fatal(err)
}

if data != nil {
    log.Printf("Result type: %s\nResult: %+v",
        data.(RunSqlResponse).ResultType,
        data.(RunSqlResponse).Result,
    )
}
```

### Permissions DSL

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/wgnet/gosura/gosura"
)

func main() {
	insert := gosura.NewCreateInsertPermissionQuery()

	perm := gosura.NewInsertPermission()

	perm.Columns = gosura.NewPGColumn().AddColumn("*")
	perm.Set = map[string]interface{}{
		"AuthorID": "X-HASURA-USER-ID",
	}

	andExp := gosura.NewBoolExp()
	exp1 := gosura.NewBoolExp()
	exp1.AddKV("UserID", 1)
	exp2 := gosura.NewBoolExp()
	exp2.AddKV("GroupID", 100)
	andExp.AddExp(gosura.AND_EXP_TYPE, exp1, exp2)

	perm.Check = andExp
	perm.Check.AddKV("AuthorID", "X-HASURA-USER-ID")

	insert.SetArgs(gosura.CreateInsertPermissionArgs{
		Table:      "test",
		Role:       "my_role",
		Comment:    "This is a test insert permission",
		Permission: perm,
	})

	data, err := json.MarshalIndent(insert, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

Result JSON:
```go
{
  "args": {
    "table": "test",
    "role": "my_role",
    "permission": {
      "check": {
        "$and": [
          {
            "UserID": 1
          },
          {
            "GroupID": 100
          }
        ],
        "AuthorID": "X-HASURA-USER-ID"
      },
      "set": {
        "AuthorID": "X-HASURA-USER-ID"
      },
      "columns": "*"
    },
    "comment": "This is a test insert permission"
  },
  "type": "create_insert_permission"
}
```

### Bulk queries

Ok, let's go to rewriting previous example

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/wgnet/gosura/gosura"
)

func insertQuery() gosura.Query {
	insert := gosura.NewCreateInsertPermissionQuery()

	perm := gosura.NewInsertPermission()

	perm.Columns = gosura.NewPGColumn().AddColumn("*")
	perm.Set = map[string]interface{}{
		"AuthorID": "X-HASURA-USER-ID",
	}

	andExp := gosura.NewBoolExp()
	exp1 := gosura.NewBoolExp()
	exp1.AddKV("UserID", 1)
	exp2 := gosura.NewBoolExp()
	exp2.AddKV("GroupID", 100)
	andExp.AddExp(gosura.AND_EXP_TYPE, exp1, exp2)

	perm.Check = andExp
	perm.Check.AddKV("AuthorID", "X-HASURA-USER-ID")

	insert.SetArgs(gosura.CreateInsertPermissionArgs{
		Table:      "test",
		Role:       "my_role",
		Comment:    "This is a test insert permission",
		Permission: perm,
	})
	return insert
}

func runSQLQuery() gosura.Query {
	q := gosura.NewRunSqlQuery()
	q.SetArgs(gosura.RunSqlArgs{
		SQL: "select count(*) from public.test",
	})
	return q
}

func main() {
	bulk := gosura.NewBulkQuery()
	bulk.SetArgs(insertQuery())
	bulk.SetArgs(runSQLQuery())

	data, err := json.MarshalIndent(bulk, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

And result JSON is

```json
{
  "args": [
    {
      "args": {
        "table": "test",
        "role": "my_role",
        "permission": {
          "check": {
            "$and": [
              {
                "UserID": 1
              },
              {
                "GroupID": 100
              }
            ],
            "AuthorID": "X-HASURA-USER-ID"
          },
          "set": {
            "AuthorID": "X-HASURA-USER-ID"
          },
          "columns": "*"
        },
        "comment": "This is a test insert permission"
      },
      "type": "create_insert_permission"
    },
    {
      "args": {
        "sql": "select count(*) from public.test",
        "cascade": false,
        "check_metadata_consistency": false
      },
      "version": 1,
      "type": "run_sql"
    }
  ],
  "type": "bulk"
}
```

Run this query with the client.Do()

```go
data, err := client.Do(bulk)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", data)
```

### Data API

Get the Hasura permissions just for example:

```go
query := gosura.NewSelectQuery()

tableColumn := gosura.NewSelectColumn("table_name")
permsColumn := gosura.NewSelectColumn("permissions")
permsColumn.AddColumn("*", nil)

args := gosura.SelectArgs{
	Table: gosura.TableArgs{
		Name:   "hdb_table",
		Schema: "hdb_catalog",
	},
	Columns: []*gosura.SelectColumn{tableColumn, permsColumn},
}
query.SetArgs(args)
```

And a query JSON is
```json
{
  "type": "select",
  "args": {
    "table": {
      "schema": "hdb_catalog",
      "name": "hdb_table"
    },
    "columns": [
      "hdb_table",
      {
        "columns": [
          "*"
        ],
        "name": "permissions"
      }
    ]
  }
}
```

Run this query with `client.Do(query)`

## Changelog

See [CHANGELOG](CHANGELOG.md).

## Contributing

See [CONTRIB](CONTRIB.md).

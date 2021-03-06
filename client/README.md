# MiniDB Client

![GitHub](https://img.shields.io/github/license/ritsource/mini-db.svg)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ritsource/mini-db.svg)
![Travis (.com)](https://img.shields.io/travis/com/ritsource/mini-db.svg)

**MiniDB-Client** contains methods that programmatically interacts with the [**MiniDB-Server**](https://github.com/ritsource/mini-db/#mini-db).

# Quick start

### Installation

First, install MiniDB using and MiniDB Golang-client

```shell
go get github.com/ritsource/mini-db
go get github.com/ritsource/mini-db/client
```

### Start Server

Start a TCP-Server using MiniDB-CLI, details about starting the **MiniDB-Server** and other Server-related options [here](https://github.com/ritsource/mini-db#minidb-server)

```shell
mini-db server --backup # --backup persists the data in the filesystem
```

### Interacting with Server

Here's a pretty straightforward **code-example** that interacts with the Server using **MiniDB-Client**

```go
package main

import (
    "fmt"
    "github.com/ritsource/mini-db/client"
)

func main() {
    // Create a client instance (mdb)
    mdb := client.New("tcp", "localhost:8000") // By default the MiniDB-Server listens on Port-8000

    // Communicating to the Server
    resp0, err := mdb.Set("myname", "Ritwik Saha", "str") // "myname" => key, "Ritwik Saha" => value, "str" => data-type
    resp1, err := mdb.Get("myname")
    resp2, err := mdb.Delete("myname")
    resp3, err := mdb.Get("myname")

    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Printf("resp0: %+v\n", resp0) // resp0["status"] == 200
    fmt.Printf("resp1: %+v\n", resp1) // resp1["data"] == "Ritwik Saha"
    fmt.Printf("resp2: %+v\n", resp2) // resp2["status"] == 200
    fmt.Printf("resp3: %+v\n", resp3) // resp3["error"] != nil && resp3["status"] == 400
}

```

# Documentation

Read the API-Docs [https://godoc.org/github.com/ritsource/mini-db/client](https://godoc.org/github.com/ritsource/mini-db/client)

# Data Type Declaration

**MiniDB** stores **key-value pairs**, and supports **3 types** of values - **String**, **Integer**, **Binary**. While inserting data through **Client** you can provide type in the **typ** argument in **Set** method. Example,

```go
resp0, err := mdb.Set("myname", "Ritwik Saha", "str") // Stores data as a String
resp1, err := mdb.Set("myname", "100", "int") // Stores data as Integer
resp2, err := mdb.Set("myname", "Ritwik Saha to be encoded", "bin") // Encodes the provided string value into binary
```

> NOTE: If provided "" in the type argement (no value), it will consider type to be a **String**

# Happy Hacking ...
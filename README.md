# shortuuid

*Note: This library was ported from the original github.com/lithammer/shortuuid/v3 repo*
*This library eschews the Google UUID library and uses the gofrs version*

A Go library that generates concise, unambiguous, URL-safe UUIDs. Based on and
compatible with the Python library
[`shortuuid`](https://github.com/stochastic-technologies/shortuuid).

Often, one needs to use non-sequential IDs in places where users will see them,
but the IDs must be as concise and easy to use as possible. shortuuid solves
this problem by generating UUIDs using
[google/uuid](https://github.com/gofrs/uuid) and then translating them to
base57 using lowercase and uppercase letters and digits, and removing
similar-looking characters such as l, 1, I, O and 0.

## Usage

```go
package main

import (
    "fmt"

    "github.com/17twenty/shortuuid"
)

func main() {
    u := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe
}
```

To use UUID v5 (instead of the default v4), use `NewWithNamespace(name string)`
instead of `New()`.

```go
shortuuid.NewWithNamespace("http://example.com")
```

It's possible to use a custom alphabet as well, though it has to be 57
characters long.

```go
alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxy="
shortuuid.NewWithAlphabet(alphabet) // u=BFWRLr5dXbeWf==iasZi
```

Bring your own encoder! For example, base58 is popular among bitcoin.

```go
package main

import (
    "fmt"
    "github.com/btcsuite/btcutil/base58"
    "github.com/17twenty/shortuuid"
    "github.com/gofrs/uuid"
)

type base58Encoder struct {}

func (enc base58Encoder) Encode(u uuid.UUID) string {
    return base58.Encode(u.Bytes())
}

func (enc base58Encoder) Decode(s string) (uuid.UUID, error) {
    return uuid.FromBytes(base58.Decode(s))
}

func main() {
    enc := base58Encoder{}
    fmt.Println(shortuuid.NewWithEncoder(enc)) // 6R7VqaQHbzC1xwA5UueGe6
}
```

## License

MIT as the original

# btcmarketsgo
Golang btcmarkets api client

## How to use
**NOTE: Full example available in /example**
### Auto setup
Example in example folder.
Either manual entry of API and private key, or store in a api.secret file within the example directory with the following format;

```
Public API key here
Private key here
```

There should only be two lines in the file, the API key and the private key.

 **Please note your keys should be kept private, use a file at your own risk**

### Setup
#### Import
```go
import "github.com/RyanCarrier/btcmarketsgo"
```
#### Client
```go
client, err := btcmarketsgo.NewDefaultClient(public, private)
```
### Tick
```go
//get tick
tick,err:=client.Tick()
//Print whole tick struct
fmt.Printf("%+v",tick)
//Print just the best ask
fmt.Println(tick.BestAsk)
```

### Buying/Selling

Price and volume when buying or selling are both \*10^-8, as specified in the BTCMarkets API;

`$12.34 = 1,234,000,000; 12.34BTC=1,234,000,000`


## Donate
Currency|Address
---|---
BTC | 1nUFrmrDCCmN6QNzXtfk3qN4w5jW2z9Sq
ZEC | t1c1YqjVd81HY3ALYE9oRWeR3pQAf6YwbhC
ETH | 0x9627ae1ab10a7172e7e25c167cb4f36d37ffdf08




------------
Software is provided as is, I'm not responsible for anything stupid you do or anything that goes wrong.

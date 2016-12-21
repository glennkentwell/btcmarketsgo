# btcmarkets-go
Golang btcmarkets api client

## How to use

### Example/setup
Example in example folder.
Either manual entry of API and private key, or store in a api.secret file within the example directory with the following format;

```
Public API key here
Private key here
```

There should only be two lines in the file, the API key and the private key.

 **Please note your keys should be kept private, use a file at your own risk**



### Monitoring
A ticker wrapper has been added for ease of action on price changes. A simple stdout printer is included in the example, a more complex example would be to send email on large price changes.

### Buying/Selling

Price and volume when buying or selling are both \*10^-8, as specified in the BTCMarkets API;

`$12.34 = 1,234,000,000; 12.34BTC=1,234,000,000`






------------
Software is provided as-is, I'm not responsible for anything stupid you do or anything that goes wrong.

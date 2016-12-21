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

The keys.go file can be used as a standalone to aid with setup of keys. Where as main is more of an example to help.

### Monitoring
A ticker has been setup for monitoring the price of BTC (or other instrument specified)

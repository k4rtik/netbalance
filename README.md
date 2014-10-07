# netbalance

Simple CLI client written in Go to check internet balance for You Broadband.

Works on Mac OS X, Linux and Windows.

## Usage
Assuming you have a go environment setup as described at [How to Write Go Code][code]:

1. `$ go get github.com/k4rtik/netbalance`
1. `$ go install github.com/k4rtik/netbalance`
1. Move the config file `.netbalance.gcfg` to your home directory and add your login credentials.
1. `$ netbalance`

## Sample Output
```
$ netbalance 
MBs: 45351.0, Days: 26.0
```

## Other providers
- For DSL providers like Airtel, ACT, see [Netlimit][netlimit].
- For other cable providers needing web based authentication, it should be easy to modify `netbalance`, please feel free to do so.

[code]: http://golang.org/doc/code.html
[netlimit]: https://github.com/arg0s/netlimit

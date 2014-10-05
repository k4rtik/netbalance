# netbalance

Simple cli client written in Go to check internet balance for You Broadband.

## Usage
Assuming you have a go environment setup as described at [How to Write Go Code][code]:
1. `$ go get github.com/k4rtik/netbalance`
1. Create config file `.netbalance.gcfg` like the following in your home directory:
```
[credentials]
login=<username>
password=<password>
```
1. `$ go install github.com/k4rtik/netbalance`
1. `$ netbalance`

[code]: http://golang.org/doc/code.html

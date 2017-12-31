<p align="center">
  <img alt="ctw Logo" src="https://raw.githubusercontent.com/eftakhairul/crypto-terminal-watch/master/assets/pic1.png" height="240" />
  <h1 align="center">Crypto Terminal Watch -- ctw</h1>
  <p align="center">A terminal based app for watching cryptocurrency rate in different currency</p>
  <p align="center">
    <a href="https://travis-ci.org/eftakhairul/crypto-terminal-watch"><img alt="Travis" src="https://travis-ci.org/eftakhairul/crypto-terminal-watch.svg?branch=master"></a>    
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>        
    <a href="https://goreportcard.com/report/github.com/eftakhairul/crypto-terminal-watch"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/eftakhairul/crypto-terminal-watch?style=flat-square"></a>
    <a href="http://godoc.org/github.com/eftakhairul/crypto-terminal-watch"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
    <a href="https://saythanks.io/to/eftakhairul"><img alt="SayThanks.io" src="https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square"></a>
    <a href="https://github.com/eftakhairul"><img alt="Written By: Eftakhairul Islam" src="https://img.shields.io/badge/powered%20by-Eftakhairul%20Islam-green.svg?style=flat-square"></a>
  </p>
</p>

## Example
```console
ctw BTC
```
It will returns exchnage rate only for Bitcoin (BTC) in USD. It almost supports all crypto currencies.


```console
ctw -c CAD ETH
```
It will returns exchnage rate only for Ethereum (ETH) in CAD.  You can get rate in some other currencies such as GBP, EUR, AUD, JPY etc.

```console
ctw 
```
It will returns top 10 crypto currency in USD

```console
ctw -c EUR
```
It will returns top 10 crypto currency in EUR. Default top 10 currencies. you can set limit. Look at next example.
 

 ```console
ctw -c GBP -l 20
```
It will returns top 20 crypto currency in GBP


## Manual Installation 
* MacOS 
 ```console
 brew tap eftakhairul/crypto-terminal-watch https://github.com/eftakhairul/crypto-terminal-watch
 brew install ctw
```

* Linux [Download here](https://github.com/eftakhairul/crypto-terminal-watch/raw/master/release/ctw_linux_amd64)
 ```console
 cp ctw_linux_amd64 /opt
 ln -s /opt/ctw_linux_amd64 /usr/local/bin/
```

* Windows [Download here](https://github.com/eftakhairul/crypto-terminal-watch/raw/master/release/ctw_windows_amd64.exe)
Run ctw_windows_amd64.exe in commadline

##Version
v0.0.1 

## Development setup
```console
mkdir -p $GOPATH/src/github.com/eftakhairul
git clone https://github.com/eftakhairul/crypto-terminal-watch
make setup_development
```

### Test
```console
make test
```


#### Authors 
* [Eftakhairul Islam]

####  Donate to this project
[Donate] - Donate through paypal


[Eftakhairul Islam]:https://eftakhairul.com
[Donate]:https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=eftakhairul%40gmail%2ecom&lc=CA&item_name=Eftakhairul%20world&item_number=web_product&currency_code=CAD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted
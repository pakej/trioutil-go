# Description #

TrioUtil is a golang package for use with Trio-Mobile SMS Provider. It's usage is based on Trio-Mobile's API.

# Installation #

```
$ go get https://github.com/zaimramlan/go-trioutil
```

# Usage #

Import TrioUtil
```
import "https://github.com/zaimramlan/go-trioutil"
```

Create a new trioutil reference
```
trio_token := "my-trio-mobile-token"
t := TrioUtil.New(trio_token)
```

Send an SMS
```
response, err := t.SendSms(phone, body)

if err != nil {
  // handle the error
  return
}

// do something with response
```

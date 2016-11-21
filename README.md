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

# Making Changes #
1. Create your feature branch (`git checkout -b new-feature`)  
2. Commit your changes (`git commit -am 'Some cool reflection'`)  
3. Push to the branch (`git push origin new-feature`)  
4. Create new Pull Request

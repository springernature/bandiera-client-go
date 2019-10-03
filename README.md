# bandiera-client-go

This is a client for talking to the [Bandiera][bandiera] feature flagging service from a Golang application.

This client is compatible with the v2 Bandiera API.

[![Build status][shield-build]][info-build]
[![MIT licensed][shield-license]][info-license]

## Installation

Installation is done using go module or `go get -u github.com/springernature/bandiera-client-golang`.

## Usage

```go
package example

import "github.com/springernature/bandiera-client-go"

func main() {
	b := bandiera.NewBandieraClient("http://bandiera-demo.herokuapp.com")

	if b.IsEnabled("pubserv", "show-new-search", nil) {
		// Do something nice
	}
}
```

`client.IsEnabled` also takes an optional params hash, this is for use with some of the more advanced features 
in Bandiera - user group and percentage based flags. It is in this params hash you pass in your `user_group` and `user_id`, i.e.:


```go
b.IsEnabled("pubserv", "show-new-search", bandiera.Params{"user_id": "1", "user_group": "Administrator"})
```

## Development

1. Fork this repo
2. Run `go mod download`
3. Run tests `go tests -v .`

## License

[&copy; 2019 Springer Nature](LICENSE.txt).

Bandiera Client Go is licensed under the [MIT License][mit]. 

[bandiera]: https://github.com/springernature/bandiera
[mit]: http://opensource.org/licenses/mit-license.php
[info-license]: LICENSE
[shield-license]: https://img.shields.io/badge/license-MIT-blue.svg
[info-build]: https://travis-ci.org/springernature/bandiera-client-go
[shield-build]: https://img.shields.io/travis/springernature/bandiera-client-go/master.svg






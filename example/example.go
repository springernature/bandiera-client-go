package example

import "github.com/springernature/bandiera-client-go"

func main() {
	b := bandiera.NewBandieraClient("http://bandiera-demo.herokuapp.com")

	if b.IsEnabled("pubserv", "show-new-search", nil) {
		// Do something nice
	}

	if b.IsEnabled("pubserv", "show-new-search", bandiera.Params{"user_id": "1", "user_group": "Administrator"}) {
		// Do some more nice things
	}
}

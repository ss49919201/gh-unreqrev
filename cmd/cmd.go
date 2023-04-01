package cmd

import (
	"fmt"

	"github.com/cli/go-gh"
)

func Run() {
	fmt.Println("hi world, this is the gh-unreqrev extension!")
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)
}

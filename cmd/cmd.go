package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cli/go-gh"
)

func Run() {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	ctx := context.Background()
	b, err := json.Marshal(map[string]any{
		"reviewers": []string{"ss49919201"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := map[string]any{}
	if err := client.DoWithContext(
		ctx,
		http.MethodDelete,
		// TODO: カレントディレクトリのGitから取得
		"repos/ss49919201/gh-unreqrev/pulls/1/requested_reviewers",
		bytes.NewReader(b),
		&resp,
	); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp["requested_reviewers"])
}

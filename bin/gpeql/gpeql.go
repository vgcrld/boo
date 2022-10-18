package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://my.int.galileosuite.com/atsgroup/graphql"

	payload := strings.NewReader("{\"query\":\"query {\\n  \\n  items(\\n    selector: [\\n      {\\n        types: [\\\"host\\\"]\\n      },{\\n        tags: [\\\"VMWAREHOST@LAYER\\\"]\\n      }\\n    ]\\n  ) {\\n    id\\n    label\\n    tags\\n  }\\n  \\n}\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")
	req.Header.Add("X-Galileo-Token", "eyJzZWMiOiJhNDlmNzc1ZTM0NDFmZTNlZmZhODMxMzJlOTQwYWU3OSIsInR5cCI6InVzZXIiLCJrZXkiOiJiNjU2ZTY5ZmUwNmZkYjVjZmFiY2RmNmZkMGEwNjhhNSJ9")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))

}

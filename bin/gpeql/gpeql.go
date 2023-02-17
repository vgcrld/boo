package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://my.int.galileosuite.com/atsgroup/graphql"

	payload := strings.NewReader("{\"query\":\"query {  items( selector: [ { types: [\\\"host\\\"] },{ tags: [\\\"VMWAREHOST@LAYER\\\"] } ] ) { id label tags } }\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")
	req.Header.Add("X-Galileo-Token", "eyJzZWMiOiIyODE0MDc1YjI1OThhMjAyMGFiNTM3ZjBjODNjNTc1MSIsInR5cCI6InVzZXIiLCJrZXkiOiI3MTE0ZmM0NTQ4M2M1Mjg3NDQ5YTY2NmM5YjE0NGYzYiJ9")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))

}

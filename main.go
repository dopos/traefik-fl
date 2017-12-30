/*
docker exec -ti dcape_webhook_1 curl http://traefik:8080/api/providers/docker/frontends \
  | jq '.[] | { "r": .routes[].rule | sub("Host:";""), "b": .backend | split("-";"" )[1] }'
*/

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
)

type route struct {
	Rule string `json:"rule"`
}

type front struct {
	EntryPoints    []string         `json:"entryPoints"`
	Backend        string           `json:"backend"`
	Routes         map[string]route `json:"routes"`
	PassHostHeader bool             `json:"passHostHeader"`
	Priority       int              `json:"priority"`
	BasicAuth      []string         `json:"basicAuth"`
}

type fronts map[string]front

type result map[string][]string

func main() {

	host := flag.String("host", "http://traefik:8080", "traefik host")
	url := flag.String("url", "/api/providers/docker/frontends", "traefik frontends url")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := http.Get(*host + *url)
		defer res.Body.Close()
		var body fronts
		json.NewDecoder(res.Body).Decode(&body)
		ret := result{}
		for _, v0 := range body {
			var v []string
			for _, v1 := range v0.Routes {
				v = append(v, strings.TrimPrefix(v1.Rule, "Host:"))
			}
			k := strings.TrimPrefix(v0.Backend, "backend-")
			ret[k] = v
		}
		json.NewEncoder(w).Encode(ret)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

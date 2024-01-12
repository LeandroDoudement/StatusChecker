package main

import (
	"fmt"
	"net/http"
)

func monitora(sites []string) {
	for i, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site", i, ":", site, "foi carregado com sucesso! Status Code:", resp.StatusCode)
			registraLog(site, true, resp.StatusCode)

		} else {
			fmt.Println("Site", i, ":", site, "está com problemas. Status Code:", resp.StatusCode)
			registraLog(site, false, resp.StatusCode)
		}
	}
}

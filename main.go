package main

import (
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/xluohome/phonedata"
)

//go:embed phone.dat
var phonedb []byte

func main() {

	ioutil.WriteFile("phone.dat", phonedb, 0644)

	http.HandleFunc("/phonedata", func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorln(err)
			return
		}
		defer r.Body.Close()

		g := gjson.ParseBytes(b)
		p, err := phonedata.Find(g.Get("phone").String())
		if err != nil {
			log.Errorln(err)
			return
		}

		x := map[string]interface{}{}
		x["data"] = p
		a, err := json.Marshal(x)
		if err != nil {
			log.Errorln(err)
			return
		}

		w.Write(a)
	})

	http.ListenAndServe("0.0.0.0:8082", nil)
}

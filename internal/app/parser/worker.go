package parser

import (
	"crypto/tls"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go/log"
)

// StartWorker ...
func (d *Parser) StartWorker(minDuration, maxDuration int) {
	rand.Seed(time.Now().UnixNano())
	client := http.Client{
		Timeout:   15 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	for {
		dur := random(minDuration, maxDuration)
		time.Sleep(time.Duration(dur) * time.Minute)

		resp, err := client.Get(d.Target.URL)
		if err != nil {
			log.Error(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			continue
		}

		if checkContains(string(body), d.Target.Text) {
			msgErr := d.Bot.SendMessage(d.Target.URL)
			if msgErr != nil {
				log.Error(msgErr)
			}
		}
	}
}

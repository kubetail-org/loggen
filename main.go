// Copyright 2024 Andres Morey
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lthibault/jitterbug/v2"
)

var ips = []string{}

var userAgents = []string{
	"curl/8.4.0",

	// googlebot
	"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",

	// yandex
	"Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)",

	// chrome on windows 10
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",

	// edge on windows 10
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",

	// safari on osx
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",

	// iphone
	"Mozilla/5.0 (iPhone13,2; U; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/15E148 Safari/602.1",

	// firefox on linux
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0",
}

var methods = map[string]int{
	"GET":  70,
	"POST": 30,
}

var endpoints = map[string]map[string]int{
	"GET": {
		"/":                30,
		"/favicon.ico":     10,
		"/settings":        15,
		"/about":           15,
		"/api/v1/settings": 20,
		"/api/v1/users":    10,
	},
	"POST": {
		"/api/v1/settings": 60,
		"/api/v1/users":    40,
	},
}

var httpCodes = map[string]int{
	"200": 80,
	"301": 5,
	"403": 10,
	"404": 5,
}

func pickRand(l []string) string {
	return (l[rand.Intn(len(l))])
}

func pickWeighted(m map[string]int) string {
	randWeight := rand.Intn(100)

	for key, weight := range m {
		randWeight -= weight
		if randWeight <= 0 {
			return key
		}
	}

	return ""
}

// IPsFromCIDR generates a slice of IPs from a given CIDR block
func IPsFromCIDR(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}

	// Remove network and broadcast addresses
	if len(ips) > 2 {
		return ips[1 : len(ips)-1]
	}
	return ips
}

// incIP increments an IP address
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func randLogLine(ansi bool) string {
	dateStr := time.Now().UTC().Format("02/01/2006:15:04:05")
	ip := pickRand(ips)
	method := pickWeighted(methods)
	endpoint := pickWeighted(endpoints[method])
	httpCode := pickWeighted(httpCodes)
	uaString := pickRand(userAgents)

	if ansi {
		ip = "\x1b[38;5;0045m" + ip + "\x1b[0m"
		method = "\x1b[33m" + method + "\x1b[0m"
		endpoint = "\x1b[32m" + endpoint + "\x1b[0m"
		httpCode = "\x1b[31m" + httpCode + "\x1b[0m"
		uaString = "\x1b[36m" + uaString + "\x1b[0m"
	}

	return fmt.Sprintf("%s - - [%s] \"%s %s HTTP/1.1\" %s 162 \"-\" \"%s\"",
		ip,
		dateStr,
		method,
		endpoint,
		httpCode,
		uaString,
	)
}

func main() {
	// cli options
	ansiPtr := flag.Bool("ansi", false, "Use ANSI color-encoding in examples")
	flag.Parse()

	interval := 1 * time.Second
	jitter := 1 * time.Second

	// initialize global ip list
	ips = append(ips, IPsFromCIDR("123.25.44.0/28")...)
	ips = append(ips, IPsFromCIDR("25.1.4.128/28")...)
	ips = append(ips, IPsFromCIDR("13.250.24.0/28")...)
	ips = append(ips, IPsFromCIDR("78.21.94.128/28")...)

	// initialize context and listen for termination signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// init timer with jitter
	ticker := jitterbug.New(interval, &jitterbug.Norm{Stdev: jitter})

	// main loop
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case <-ticker.C:
			fmt.Println(randLogLine(*ansiPtr))
		}
	}

	stop() // stop receiving signals as soon as possible
}

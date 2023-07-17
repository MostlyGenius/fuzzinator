package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	browsers  = []string{"Mozilla", "Chrome", "Safari", "Opera", "Edge", "Internet Explorer"}
	osList    = []string{"Windows NT 10.0; Win64; x64", "Macintosh; Intel Mac OS X 10_15_7", "X11; Linux x86_64", "iPhone; CPU iPhone OS 13_2_3 like Mac OS X", "Windows NT 6.1; Win64; x64"}
	languages = []string{"en-US", "en-GB", "es-ES", "fr-FR", "de-DE", "ru-RU", "ja-JP", "zh-CN"}
	versions  = map[string][]string{
		"Mozilla":           {"5.0", "4.0", "3.5"},
		"Chrome":            {"91.0.4472.124", "92.0.4515.159", "93.0.4577.63"},
		"Safari":            {"537.36", "537.31", "534.30"},
		"Opera":             {"60.0.3255.170", "58.0.3135.127", "56.0.3051.36"},
		"Edge":              {"91.0.4472.124", "92.0.902.67", "93.0.961.38"},
		"Internet Explorer": {"11.0", "10.0", "9.0"},
	}
)

func generateUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	browser := browsers[rand.Intn(len(browsers))]
	os := osList[rand.Intn(len(osList))]
	version := versions[browser][rand.Intn(len(versions[browser]))]

	userAgent := fmt.Sprintf("%s/%s (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s/%s Safari/537.36", browser, version, os, browser, version)
	return userAgent
}

func main() {
	userAgents := map[string]bool{}

	for len(userAgents) < 1000 {
		userAgent := generateUserAgent()
		userAgents[userAgent] = true
	}

	for userAgent := range userAgents {
		fmt.Println(userAgent)
	}
}

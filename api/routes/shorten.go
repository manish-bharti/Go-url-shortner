package routes

import "time"

type request struct {
	URL					string					`json:"url"`
	CustomShort			string					`json:"short"`
	Expiry				time.Duration			`json:"expiry"`
}

type response struct {
	URL				string 						`json:"url"`
	CustomShort		string						`json:"short"`
	Expiry			time.Duration				`json:"expiry"`		
	XRateRemaining	int							`json:"rate_limit"`
	XrateLimitRest	time.Duration				`json:"rate_limit_reset"`
}


func ShortenUrl(c *fiber.Ctx) error{

	body := new(request)

	if err := c.BodyParser(body); err != nil{

	}
}


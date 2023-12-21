package main

import (
	"flag"
	"fmt"
	"strings"

	tr "github.com/snakesel/libretranslate"
)

func main() {
	lang := flag.String("l", "auto", "Source language. E.g. 'en' or 'fr'")
	target := flag.String("t", CONFIG.Lang, "Target language. E.g. 'en' or 'fr'")
	isConfig := flag.Bool("config", false, `Set default config then exit. 
Options are lang(default target language), url(libretranslate server url), key(api key)
Usage: [option]=[value] E.g. lang=en key=your_api_key
        `)
	flag.Parse()
	args := flag.Args()

	if *isConfig {
		SetConfig(args)
		return
	}

	query := strings.Join(args, " ")
	if len(query) == 0 {
		fmt.Println("No text provided")
		return
	}

	res, err := Translate(query, *lang, *target)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	res.Render()
}

type Response struct {
	Text       string
	Lang       string
	Confidence float64
}

func (r *Response) Render() {
	fmt.Printf("\tLang: %s (%d%%)\n", r.Lang, int(r.Confidence))
	fmt.Printf("\tText: %s\n", r.Text)
}

func Translate(query, sourceLang, targetLang string) (*Response, error) {
	fmt.Printf("Translate(%s -> %s): %s\n", sourceLang, targetLang, query)
	translate := tr.New(tr.Config{Url: CONFIG.Url, Key: CONFIG.Key})
	res, err := translate.Translate(query, sourceLang, targetLang)
	if err != nil {
		return nil, err
	}

	data := Response{
		Text:       res["translatedText"].(string),
		Lang:       sourceLang,
		Confidence: 100,
	}

	if val, ok := res["detectedLanguage"]; ok {
		val := val.(map[string]interface{})
		data.Confidence = val["confidence"].(float64)
		data.Lang = val["language"].(string)
	}

	return &data, nil
}

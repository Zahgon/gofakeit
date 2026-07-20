package gofakeit

import (
	"text/template"
)

type TemplateOptions struct {
	Funcs template.FuncMap `fake:"-"`
	Data  any              `json:"data" xml:"data" fake:"-"`
}

func Template(template string, co *TemplateOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *Faker) Template(template string, co *TemplateOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type MarkdownOptions struct {
}

const templateMarkdown = `
{{$repo := Gamertag}}
{{$language := RandomString (SliceString "go" "python" "javascript")}}
{{$username := Gamertag}}
{{$weightedSlice := SliceAny "github.com" "gitlab.com" "bitbucket.org"}}
{{$weightedWeights := SliceF32 5 1 1}}
{{$domain := Weighted $weightedSlice $weightedWeights}}
{{$action := RandomString (SliceString "process" "run" "execute" "perform" "handle")}}
{{$usage := RandomString (SliceString "whimsical story" "quirky message" "playful alert" "funny request" "lighthearted command")}}
{{$result := RandomString (SliceString "success" "error" "unknown" "completed" "failed" "finished" "in progress" "terminated")}}

# {{$repo}}

*Author: {{FirstName}} {{LastName}}*

{{Paragraph}}

{{Paragraph}}

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Installation
{{if eq $language "go"}}'''go
go get {{$domain}}/{{$username}}/{{$repo}}
'''{{else if eq $language "python"}}'''bash
pip install {{$repo}}
'''{{else if eq $language "javascript"}}'''js
npm install {{$repo}}
'''{{end}}

## Usage
{{if eq $language "go"}}'''go
result := {{$repo}}.{{$action}}("{{ToLower $usage}}")
fmt.Println("{{ToLower $repo}} result:", "{{ToLower $result}}")
'''{{else if eq $language "python"}}'''python
result = {{ToLower $repo}}.{{$action}}("{{ToLower $usage}}")
print("{{ToLower $repo}} result:", "{{ToLower $result}}")
'''{{else if eq $language "javascript"}}'''javascript
const result = {{ToLower $repo}}.{{$action}}("{{ToLower $usage}}");
console.log("{{ToLower $repo}} result:", "{{ToLower $result}}");
'''{{end}}

## License
{{RandomString (SliceString "MIT" "Apache 2.0" "GPL-3.0" "BSD-3-Clause" "ISC")}}
`

func Markdown(co *MarkdownOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *Faker) Markdown(co *MarkdownOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type EmailOptions struct {
}

const templateEmail = `
Subject: {{RandomString (SliceString "Greetings" "Hello" "Hi")}} from {{FirstName}}!

Dear {{LastName}},

{{RandomString (SliceString "Greetings!" "Hello there!" "Hi, how are you?")}} {{RandomString (SliceString "How's everything going?" "I hope your day is going well." "Sending positive vibes your way.")}}

{{RandomString (SliceString "I trust this email finds you well." "I hope you're doing great." "Hoping this message reaches you in good spirits.")}} {{RandomString (SliceString  "Wishing you a fantastic day!" "May your week be filled with joy." "Sending good vibes your way.")}}

{{Paragraph}}

{{Paragraph}}

{{Paragraph}}

{{RandomString (SliceString "I would appreciate your thoughts on" "I'm eager to hear your feedback on" "I'm curious to know what you think about")}} it. If you have a moment, please feel free to check out the project on {{RandomString (SliceString "GitHub" "GitLab" "Bitbucket")}}

{{RandomString (SliceString "Your insights would be invaluable." "I'm eager to hear what you think." "Feel free to share your opinions with me.")}} {{RandomString (SliceString "Looking forward to your feedback!" "Your perspective is highly valued." "Your thoughts matter to me.")}}

{{RandomString (SliceString "Thank you for your consideration!" "I appreciate your attention to this matter." "Your support means a lot to me.")}} {{RandomString (SliceString "Wishing you a wonderful day!" "Thanks in advance for your time." "Your feedback is greatly appreciated.")}}

{{RandomString (SliceString "Warm regards" "Best wishes" "Kind regards" "Sincerely" "With gratitude")}}
{{FirstName}} {{LastName}}
{{Email}}
{{PhoneFormatted}}
`

func EmailText(co *EmailOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *Faker) EmailText(co *EmailOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var templateExclusion = []string{
	"RandomMapKey",
	"SQL",
	"Template",
}

func templateFuncMap(f *Faker, fm *template.FuncMap) *template.FuncMap {
	_ = "STUB: not implemented"
	return nil
}

func templateFunc(temp string, funcs *template.FuncMap, data any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addTemplateLookup() { _ = "STUB: not implemented"; return }

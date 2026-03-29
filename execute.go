package main

import (
	"bytes"
	"text/template"
)

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

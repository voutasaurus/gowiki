package main

import (
		"html/template"
		"regexp"
		"strings"
		// "html"
		//"errors"
)

func undiscuss(title string) string {
	return strings.TrimSuffix(title, "_discuss")
}

func titleDisplay(title string) string {
	return strings.Replace(title, "_", " ", -1)
}

func convertToLink(semlink []byte) []byte {
	m := len(semlink)
	pageName := string(semlink[1:m-1])
	address := "/view/" + strings.Replace(pageName, " ", "_", -1)
	link := "<a href=\"" + address + "\">" + pageName + "</a>"
	return []byte(link)
}

func convertToSectionHead(sectionhead []byte) []byte {
	m := len(sectionhead)
	sectionName := string(sectionhead[2:m-2])
	sectionID := strings.Replace(sectionName, " ", "_", -1)
	heading := "<h3 id=\"" + sectionID + "\">" + sectionName + "</h3>"
	return []byte(heading)
	// add section to TOC
}

func parseBody(body []byte) []byte {

	// HTML parsing
	safe := template.HTMLEscapeString(string(body))
	safe = strings.Replace(safe, "\n", "<br>", -1)
	
	// Replacements
	raw := []byte(safe)
	
	// hyperlinks
	semlink := regexp.MustCompile(`\[[0-9A-Za-z_ ]+\]`) // note the space
	raw = semlink.ReplaceAllFunc(raw, convertToLink)
	
	// section headings
	sectionhead := regexp.MustCompile(`\=\=[0-9A-Za-z_ ]+\=\=`)
	raw = sectionhead.ReplaceAllFunc(raw, convertToSectionHead)
	
	// Finished processing
	processed := raw
	
	return processed
}

func show(body []byte) template.HTML {
	return template.HTML(string(body))
}

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Hotel is a hotel in search results
var Hotel = MediaType("application/json", func() {
	Attributes(func() {
		Attribute("id", String)
		Attribute("name", String)
		Required("id", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

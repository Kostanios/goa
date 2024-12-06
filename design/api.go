package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("hotels", func() {
	Title("Hotels API")
	Description("The best API ever!")
	Scheme("http")
	Host("localhost:8080")
})

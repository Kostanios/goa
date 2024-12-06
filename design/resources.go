package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("hotels", func() {
	BasePath("/hotel")
	Action("search", func() {
		Description("Search hotels")
		Routing(GET("/search"))
		Response(OK, CollectionOf(Hotel))
	})
})

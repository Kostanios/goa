package main

import (
	"encoding/xml"
	"goa/app"
	"goa/file"

	"github.com/goadesign/goa"
)

// HotelsController implements the hotels resource.
type HotelsController struct {
	*goa.Controller
}

// NewHotelsController creates a hotels controller.
func NewHotelsController(service *goa.Service) *HotelsController {
	return &HotelsController{Controller: service.NewController("HotelsController")}
}

// Search runs the search action.
func (c *HotelsController) Search(ctx *app.SearchHotelsContext) error {
	// HotelsController_Search: start_implement

	// Read XML test file with hotel details
	fileReader, err := file.GetXMLFileReader()
	if err != nil {
		return ctx.InternalServerError(err)
	}
	// Decode XML
	searchResponse := &file.SearchResponse{}
	decoder := xml.NewDecoder(fileReader)
	if err := decoder.Decode(searchResponse); err != nil {
		return ctx.InternalServerError(err)
	}
	// Convert to the right response type
	hotels := make(app.HotelCollection, len(searchResponse.ServiceHotels))
	for i, h := range searchResponse.ServiceHotels {
		hotels[i] = &app.Hotel{
			ID:   h.HotelInfo.Code,
			Name: h.HotelInfo.Name,
		}
	}
	// Reply to client
	return ctx.OK(hotels)
	// HotelsController_Search: end_implement
}

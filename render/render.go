// Code generated by "pkgimport -p render -i github.com/unrolled/render -o render.go"; DO NOT EDIT.
// Install by "go get -u -v github.com/wzshiming/gotype/cmd/pkgimport";
//go:generate pkgimport -p render -i github.com/unrolled/render -o render.go

package render

import (
	origin "github.com/unrolled/render"
)

// type
type (

	// XML built-in renderer.
	XML = origin.XML

	// Text built-in renderer.
	Text = origin.Text

	// Render is a service that provides functions for easily writing JSON, XML,
	// binary data, and HTML templates out to a HTTP Response.
	Render = origin.Render

	// Options is a struct for specifying configuration options for the render.Render object.
	Options = origin.Options

	// JSONP built-in renderer.
	JSONP = origin.JSONP

	// JSON built-in renderer.
	JSON = origin.JSON

	// Head defines the basic ContentType and Status fields.
	Head = origin.Head

	// HTMLOptions is a struct for overriding some rendering Options for specific HTML call.
	HTMLOptions = origin.HTMLOptions

	// HTML built-in renderer.
	HTML = origin.HTML

	// Engine is the generic interface for all responses.
	Engine = origin.Engine

	// Delims represents a set of Left and Right delimiters for HTML template rendering.
	Delims = origin.Delims

	// Data built-in renderer.
	Data = origin.Data

	// BufferPool implements a pool of bytes.Buffers in the form of a bounded channel.
	// Pulled from the github.com/oxtoacart/bpool package (Apache licensed).
	//BufferPool = origin.BufferPool
)

// Declaration
var (

	// NewBufferPool creates a new BufferPool bounded to the given size.
	//NewBufferPool = origin.NewBufferPool

	// New constructs a new Render instance with the supplied options.
	New = origin.New

	// ContentXML header value for XML data.
	ContentXML = origin.ContentXML

	// ContentXHTML header value for XHTML data.
	ContentXHTML = origin.ContentXHTML

	// ContentType header constant.
	ContentType = origin.ContentType

	// ContentText header value for Text data.
	ContentText = origin.ContentText

	// ContentLength header constant.
	ContentLength = origin.ContentLength

	// ContentJSONP header value for JSONP data.
	ContentJSONP = origin.ContentJSONP

	// ContentJSON header value for JSON data.
	ContentJSON = origin.ContentJSON

	// ContentHTML header value for HTML data.
	ContentHTML = origin.ContentHTML

	// ContentBinary header value for binary data.
	ContentBinary = origin.ContentBinary
)

package index

import (
	"encoding/xml"
)

type Configuration struct {
	DSpaceHost      string
	Collections     []string
	SolrUrl         string
	SolrCore        string
	FileFormat      string
	IndexType       string
	EscapeUtf8      bool
	XmlFileLocation string
	HttpPort        string
	IpWhitelist     []string
	LogDir		    string
}

type SolrField struct {
	XMLName xml.Name `xml:"field"`
	Name    string   `xml:"name,attr"`
	Field   string   `xml:",cdata"`
}

type SolrPostXml struct {
	XMLName xml.Name `xml:"doc"`
	Fields  []SolrField
}

type Alto struct {
	XMLName  xml.Name  `xml:"alto"`
	Xmlns    string    `xml:"xmlns,attr,omitempty"`
	Description Description `xml:"Description"`
	Styles Styles `xml:"Styles"`
	Layout Layout `xml:"Layout"`
}
type Description struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	MeasurementUnit string `xml:"MeasurementUnit"`
	SourceImageInformation SourceImageInformation `xml:"sourceImageInformation"`
}

type SourceImageInformation struct {
	FileName string `xml:"fileName"'`
}

type Styles struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	ParagraphStyle []ParagraphStyle `xml":ParagraphStyle"`
}

type ParagraphStyle struct {
	Id string `xml:"ID,attr"`
	ALIGN string `xml:"ALIGN,attr"`
	LEFT string `xml:"LEFT,attr"`
	RIGHT string `xml:"RIGHT,attr"`
	FIRSTLINE string `xml:"FIRSTLINE,attr"`
	LINESPACE string `xml:"LINESPACE,attr"`
}

type Layout struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	Page Page	`xml:"Page"`
}
// Page Alto page element
type Page struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	Id            string `xml:"ID,attr"`
	PhysicalImgNr string `xml:"PHYSICAL_IMG_NR,attr"`
	Height        string `xml:"HEIGHT,attr"`
	Width         string `xml:"WIDTH,attr"`
	PrintSpace    PrintSpace `xml:"PrintSpace"`
}

// PrintSpace Alto composed block element (not always present, contains TextBlock elements)
type PrintSpace struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	TextBlock     []TextBlock     `xml:"TextBlock"`
	ComposedBlock []ComposedBlock `xml:"ComposedBlock"`
}

// ComposedBlock Alto composed block element (not always present, contains TextBlock elements)
type ComposedBlock struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	TextBlock []TextBlock `xml:"TextBlock"`
}

// TextBlock Alto text block element
type TextBlock struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	TextLine []TextLine `xml:"TextLine"`
}

// TextLine Alto line element
type TextLine struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	String []String `xml:"String"`
}

// String Alto element containing word information
type String struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	CONTENT string `xml:"CONTENT,attr"`
	HEIGHT  string `xml:"HEIGHT,attr"`
	WIDTH   string `xml:"WIDTH,attr"`
	VPOS    string `xml:"VPOS,attr"`
	HPOS    string `xml:"HPOS,attr"`
}

// OcrEl MiniOcr base element
type OcrEl struct {
	XMLName xml.Name `xml:"ocr"`
	Pages   []P      `xml:"p"`
}

// P MiniOcr page element
type P struct {
	XMLName    xml.Name `xml:"p"`
	Id         string   `xml:"xml:id,attr"`
	Dimensions string   `xml:"wh,attr"`
	Blocks     []B      `xml:"b"`
}

// B MiniOcr block element
type B struct {
	XMLName xml.Name `xml:"b"`
	Lines   []L      `xml:"l"`
}

// L MiniOcr line element
type L struct {
	XMLName xml.Name `xml:"l"`
	Words   []W      `xml:"w"`
}

// W MiniOcr word element
type W struct {
	XMLName       xml.Name `xml:"w"`
	CoorinateAttr string   `xml:"x,attr"`
	Content       string   `xml:",chardata"`
}

// IIIF structs.

type Manifest struct {
	Context     string     `json:"@context"`
	Type        string     `json:"@type"`
	Id          string     `json:"@id"`
	Label       string     `json:"label"`
	Metadata    []Metadata `json:"metadata"`
	Service     Service    `json:"service,omitempty"`
	SeeAlso     SeeAlso    `json:"SeeAlso,omitempty"`
	Sequences   []Sequence `json:"sequences"`
	Thumbnail   Thumbnail  `json:"thumbnail,omitempty"`
	ViewingHint string     `json:"viewingHint,omitempty"`
	Related     Related    `json:"related,omitempty"`
}
type Metadata struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Service struct {
	Context  string `json:"@context"`
	Id       string `json:"@id"`
	Profile  string `json:"profile"`
	Protocol string `json:"protocol"`
}

type SeeAlso struct {
	Id    string `json:"@id"`
	Type  string `json:"@type"`
	Label string `json:"label"`
}

type Sequence struct {
	Id       string   `json:"@id"`
	Type     string   `json:"@type"`
	Canvases []Canvas `json:"canvases"`
}

type Canvas struct {
	Id        string    `json:"@id"`
	Type      string    `json:"@type"`
	Label     string    `json:"label"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Images    []Image   `json:"images"`
	Width     int       `json:"width,int"`
	Height    int       `json:"height,int"`
}

type Thumbnail struct {
	Id      string  `json:"@id"`
	Type    string  `json:"@type"`
	Label   string  `json:"label"`
	Service Service `json:"service"`
	Format  string  `json:"format"`
}

type Image struct {
	Type       string   `json:"@type"`
	Motivation string   `json:"motivation"`
	Resource   Resource `json:"resource"`
	On         string   `json:"on"`
}

type Resource struct {
	Id      string  `json:"@id"`
	Type    string  `json:"@type"`
	Service Service `json:"service"`
	Format  string  `json:"format"`
}

type Related struct {
	Id     string `json:"@id"`
	Label  string `json:"label"`
	Format string `json:"format"`
}

type ResourceAnnotationList struct {
	Context   string               `json:"@context"`
	Id        string               `json:"@id"`
	Type      string               `json:"@type"`
	Resources []ResourceAnnotation `json:"resources"`
}
type ResourceAnnotation struct {
	Id         string                     `json:"@id"`
	Type       string                     `json:"@type"`
	Motivation string                     `json:"motivation"`
	Resource   ResourceAnnotationResource `json:"resource"`
}
type ResourceAnnotationResource struct {
	Id     string `json:"@id"`
	Label  string `json:"label"`
	Format string `json:"format"`
}

type SolrResponse struct {
	ResponseHeader struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
	} `json:"responseHeader"`
	Response struct {
		NumFound      int           `json:"numFound"`
		Start         int           `json:"start"`
		NumFoundExact bool          `json:"numFoundExact"`
		Docs          []Docs 		`json:"docs"`
	} `json:"response"`
}

type Docs struct {
	ManifestUrl []string `json:"manifest_url,omitempty"`
	OcrText string `json:"ocr_text,omitempty"`
}

type SolrCreatePost struct {
	Id          string `json:"id"`
	ManifestUrl string `json:"manifest_url"`
	OcrText     string `json:"ocr_text"`
}

type SolrDeletePost struct {
	Delete Delete  `json:"delete"`
}

type Delete struct {
	Query string `json:"query"`
}

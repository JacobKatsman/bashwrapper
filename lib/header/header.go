package header

// filename type string for take parameter  function 
type FileOfName string

// file operation for source and destination
type FileStructSource struct {}
type FileStructDestination struct {}

// interface file operation 
type FilStatService interface {
	GetSizeFile(filePath FileScheme) int
}

// file Item scheme
type FileScheme struct {
	Source string
	Destination string
}

func (c FileOfName) String() string {
        return string(c)
}

type SizeTable struct { 
    Extend string 
    Count  int
	Size   float64 
}

var T1 = []SizeTable{
    SizeTable{
 		Extend: "jpg", 
		Count:  0,
	    Size:   0,
    },
    SizeTable{
 		Extend: "pdf", 
		Count:  0,
	    Size:   0,
    },
    SizeTable{
 		Extend: "none", 
		Count:  0,
	    Size:   0,
    },		
}

var Jpeg_ext = []string {".jpg", ".jpeg"}
var Pdf_ext  = []string {".pdf", ".PDF"}

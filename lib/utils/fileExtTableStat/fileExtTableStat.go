package fileExtTableStat

import (
	"os"
	"fmt"
	"io/ioutil"
	"text/tabwriter"
	"path/filepath"
	"golang.org/x/exp/slices"
    . "main/lib/header"
	. "main/lib/utils/getFileSize"
)

// *
// * Calculation in the table that counting files and summurizes filesize by extension
// * Tabular representation of file sizes and numbers
// *

type SizeFormatSlice []float64
type ExtendedFileScheme  FileScheme

type FS1 interface { 
	CalcSize(fullfilePath string, Size []SizeTable) []SizeTable
}

type FileSizeTable struct{}

func MatchFilesType (fullfilePath string, file os.FileInfo, Size []SizeTable) []SizeTable{
            var  t FS1 = FileSizeTable{}
	 		if file.IsDir() {
 				var h = t.CalcSize((fullfilePath + "/" + file.Name()), Size)
			    for i := 0; i <= (len(h) - 1) ; i++ {
					Size[i].Size  = h[i].Size
					Size[i].Count = h[i].Count
                }
  			} else {
  	            extFile := filepath.Ext(file.Name())
  		        if (slices.Index(Jpeg_ext, extFile)) != -1  {
 					Size[0].Size += GetFileSize(fullfilePath + "/" + file.Name())
					Size[0].Count +=1
 		        } else if slices.Index(Pdf_ext, extFile) != -1  {
 					Size[1].Size += GetFileSize(fullfilePath + "/" + file.Name())
					Size[1].Count +=1
 		        } else {
 					Size[2].Size += GetFileSize(fullfilePath + "/" + file.Name())
					Size[2].Count +=1
 		        }
				 	
 		    }
	return Size
}

func (m FileSizeTable) CalcSize(fullfilePath string, Size []SizeTable) []SizeTable{
  var  LocalSize []SizeTable
  LocalSize = Size
  files, _ := ioutil.ReadDir(fullfilePath)
	for _, file := range files {
		LocalSize = MatchFilesType (fullfilePath,file, Size)
	 }
	return LocalSize
}

func CSize(fullfilePath string) string {
	var m FileSizeTable
    var t []SizeTable
	var s string

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
	
	t = m.CalcSize(fullfilePath,T1)
    w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', tabwriter.Debug)
    for i := range t {
        p := t[i]
	    fmt.Fprintf(w, "%d\t%s\t%d\t%.2fMB\t\n", i, p.Extend, p.Count, p.Size)
    }
	w.Flush()
	return s
}

package fileChangeCopyFiles

import (
	"os"
	"log"
	"io/ioutil"
	"path/filepath"
	"golang.org/x/exp/slices"
    . "main/lib/header"
	copy "main/lib/utils/module/fileOperationModule"
	jpgOperationModule  "main/lib/utils/module/jpgOperationModule"
	pdfOperationModule  "main/lib/utils/module/pdfOperationModule"
	timeOperationModule "main/lib/utils/module/timeOperationModule"
)

// *
// * Calculation in the table that counting files and summurizes filesize by extension
// * tabular representation of file sizes and numbers
// *

type SizeFormatSlice []float64
type ExtendedFileScheme  FileScheme

type FS1 interface { 
	CopyFiles(n FileScheme) bool
}

type FileSizeTable struct{}


func MatchFilesTypeDir (n FileScheme, file os.FileInfo) bool{
	var s FileScheme
	var m FileSizeTable
    err := os.Mkdir(n.Destination + "/" + file.Name(), (os.ModePerm));
		   if  err != nil { log.Fatal(err)}
	       s = FileScheme{
				   Source:      n.Source + "/" + file.Name(),
				   Destination: n.Destination + "/" + file.Name(),
	           }
		   m.CopyFiles(s)
	return true
}

func MatchFilesTypeFile (n FileScheme, file os.FileInfo) bool{
	extFile := filepath.Ext(file.Name())
	if (slices.Index(Jpeg_ext, extFile)) != -1  {
			copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
		    jpgOperationModule.SimpleJpg(n,file.Name())
	        } else if slices.Index(Pdf_ext, extFile) != -1  {
		    pdfOperationModule.SimplePdf(n,file.Name())
		    } else {
			copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
		   }
	return true
}


func (m FileSizeTable) CopyFiles(n FileScheme) bool{
  files, _ := ioutil.ReadDir(n.Source)
	for _, file := range files {
		if file.IsDir() {
			MatchFilesTypeDir  (n, file)
 			} else {			
			MatchFilesTypeFile (n, file)
		}
		strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
		timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
	 }
	return true
}

func CFiles(n FileScheme) string {
	var m FileSizeTable
	var s string
	m.CopyFiles(n)
    s = "Copy has compleat! \n"
	return s
}

package fileChangeCopyFiles

import (
	"os"
	"log"
	"io/ioutil"
	"path/filepath"
	"golang.org/x/exp/slices"
    . "main/lib/header"
	"sync"
	"fmt"
	fileStat "main/lib/utils/fileStat"
	copy "main/lib/utils/module/fileOperationModule"
	jpgOperationModule  "main/lib/utils/module/jpgOperationModule"
	pdfOperationModule  "main/lib/utils/module/pdfOperationModule"
	timeOperationModule "main/lib/utils/module/timeOperationModule"
)

// *
// * Calculation for table for counting files and summurizes filesize by the extension
// * (tabular representation of file sizes and numbers)
// *

type SizeFormatSlice []float64
type ExtendedFileScheme  FileScheme

type FS1 interface { 
	CopyFiles(n FileScheme) bool
}

type FileSizeTable struct{}


func MatchFilesTypeDir (id int, n FileScheme, file os.FileInfo) int{
	var s FileScheme
	var m FileSizeTable

	if (id%(block_size) == block_size) {
		wg.Wait()
    }

	fmt.Printf("source: %s id = %d  FOLDER COPY START \t\n", n.Source + "/" + file.Name(), id)
    err := os.Mkdir(n.Destination + "/" + file.Name(), (os.ModePerm));
		   if  err != nil { log.Fatal(err)}
	       s = FileScheme{
				   Source:      n.Source + "/" + file.Name(),
				   Destination: n.Destination + "/" + file.Name(),
		   }

	fmt.Printf("destination: %s id = %d FOLDER COPY DONE \t\n", n.Destination + "/" + file.Name(), id)
	i := m.CopyFiles(id, s)

	return i
}

func MatchFilesTypeFile (id  int, n FileScheme, file os.FileInfo) int{
	if (id%(block_size) == block_size) {
		wg.Wait()
    }
	
	extFile := filepath.Ext(file.Name())
	if (slices.Index(Jpeg_ext, extFile)) != -1  {
	    wg.Add(1)
		go func() {
			    fmt.Printf("source: %s id = %d  JPG START \t\n", n.Source + "/" + file.Name(), id)
		 		defer wg.Done()
		        copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
				jpgOperationModule.SimpleJpg(n,file.Name())
				strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
			    timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
			    fmt.Printf("destination: %s id = %d  JPG DONE \t\n", n.Source + "/" + file.Name(), id)
		}()
	        } else if slices.Index(Pdf_ext, extFile) != -1  {
		wg.Add(1)
		go func() {
			    fmt.Printf("source: %s id = %d  PDF START \t\n", n.Source + "/" + file.Name(), id)
				defer wg.Done()
				pdfOperationModule.SimplePdf(n,file.Name())
				strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
			    timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
			    fmt.Printf("destination: %s id = %d  PDF DONE \t\n", n.Source + "/" + file.Name(), id)
		}()
	} else {
		// overwise copy
		copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
	}
	return id
}


func (m FileSizeTable) CopyFiles(id int ,n FileScheme) int{
	i := id  
	files, _ := ioutil.ReadDir(n.Source)
	for _, file := range files {
		if file.IsDir() {
			i = MatchFilesTypeDir  ((i + 1), n, file)
		} else {
			copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
			i = MatchFilesTypeFile ((i + 1), n, file)
	     }
	 }
	return i
}

//  How much files into the folder?
//  посчитать файлы во всех рекурсивных директориях
func  getFileTotalCount(n FileScheme) int {
    m := fileStat.CSize(n.Source)
	return int(m[1])
}

//  Set Constant Name and global variables value
var wg sync.WaitGroup
var m FileSizeTable

var count_max_record int
var block_size = 10
var start_value = 0

func CFiles(n FileScheme) string {
	count_max_record = getFileTotalCount(n)
	if (count_max_record < block_size) {block_size = count_max_record} 
    fmt.Printf( "Total  files = %d \n", count_max_record)
	m.CopyFiles(start_value, n)
	wg.Wait()
	return "Processing does compleat! \n"
}

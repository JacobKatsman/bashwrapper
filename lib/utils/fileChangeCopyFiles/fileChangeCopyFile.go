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


func MatchFilesTypeDir (id int, n FileScheme, file os.FileInfo) bool{
	var s FileScheme
	var m FileSizeTable

    fmt.Printf("Folder processing  %d (id mod 5 = %d) START\n", id , id%(block_size + 1))
	
    err := os.Mkdir(n.Destination + "/" + file.Name(), (os.ModePerm));
		   if  err != nil { log.Fatal(err)}
	       s = FileScheme{
				   Source:      n.Source + "/" + file.Name(),
				   Destination: n.Destination + "/" + file.Name(),
		   }
	
	m.CopyFiles(id, s)

    fmt.Printf("Folder Processing %d DONE\n", (id))
	
	return true
}

func MatchFilesTypeFile (id  int, n FileScheme, file os.FileInfo) bool{
	extFile := filepath.Ext(file.Name())
	if (slices.Index(Jpeg_ext, extFile)) != -1  {
	    wg.Add(1)
		go func() {
		 		defer wg.Done()
		        copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
	            fmt.Printf("File JPG processing %d (id mod 5 = %d) %s START\n", id , id%(block_size + 1), file.Name())
				jpgOperationModule.SimpleJpg(n,file.Name())
				strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
				timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
	            fmt.Printf("File JPG Processing %d DONE\n", (id))
		}()

	        } else if slices.Index(Pdf_ext, extFile) != -1  {
		wg.Add(1)
		go func() {
					defer wg.Done()
	            fmt.Printf("File PDF processing  %d (id mod 5 = %d) %s START\n", id, id%(block_size + 1), file.Name())
				pdfOperationModule.SimplePdf(n,file.Name())
				strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
		        timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
	            fmt.Printf("File PDF Processing %d DONE\n", (id))
		}()
	} else {
		// ?????
	}
	

	return true
}


func (m FileSizeTable) CopyFiles(id int ,n FileScheme) bool{
	i := id + 1
	if (id%(block_size + 1) == block_size) {
		wg.Wait()
    }
	files, _ := ioutil.ReadDir(n.Source)
	for _, file := range files {
		if file.IsDir() {
			   MatchFilesTypeDir  (i, n, file)
		} else {
			copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
			   MatchFilesTypeFile (i, n, file)
	     }
	 }
	return true
}

func  getFileTotalCount(n FileScheme) int {
    m := fileStat.CSize(n.Source)
	return int(m[1])
}

//  Set Constant Name and global variables value
var wg sync.WaitGroup
var m FileSizeTable

//  how much files into the folder?
//  посчитать файлы во всех рекурсивных директориях
var count_max_record int
var block_size = 4
var start_value = 0

func CFiles(n FileScheme) string {
	count_max_record = getFileTotalCount(n)
	if (count_max_record < block_size) {block_size = count_max_record} 

    fmt.Printf( "Total  files = %d \n", count_max_record)
	
	m.CopyFiles(start_value, n)
	wg.Wait()
	return "Processing does compleat! \n"
}

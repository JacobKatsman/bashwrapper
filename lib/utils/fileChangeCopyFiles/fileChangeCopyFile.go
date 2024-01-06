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

	//------
    fmt.Printf("Folder processing  %d (id mod 5 = %d) START\n", id , id%(block_size + 1))
	//------
	
    err := os.Mkdir(n.Destination + "/" + file.Name(), (os.ModePerm));
		   if  err != nil { log.Fatal(err)}
	       s = FileScheme{
				   Source:      n.Source + "/" + file.Name(),
				   Destination: n.Destination + "/" + file.Name(),
		   }
	
	m.CopyFiles(id, s)

	//------
    fmt.Printf("Folder Processing %d DONE\n", (id))
	//------
	
	return true
}

func MatchFilesTypeFile (id  int, n FileScheme, file os.FileInfo) bool{

	//------
	fmt.Printf("File processing  %d (id mod 5 = %d) START\n", id , id%(block_size + 1))
	//------
	
	extFile := filepath.Ext(file.Name())
	if (slices.Index(Jpeg_ext, extFile)) != -1  {
		    // copy utils 
		    copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
		    // jpg utils
		    jpgOperationModule.SimpleJpg(n,file.Name())
	        } else if slices.Index(Pdf_ext, extFile) != -1  {
		    // pdf utils 
		    pdfOperationModule.SimplePdf(n,file.Name())
		    } else {
			copy.CopyFile (n.Source + "/" + file.Name(),n.Destination + "/" + file.Name())
	}
	
    //------
	fmt.Printf("File Processing %d DONE\n", (id))
	//------
	return true
}


func (m FileSizeTable) CopyFiles(id int ,n FileScheme) bool{
  files, _ := ioutil.ReadDir(n.Source)
	for _, file := range files {
		if file.IsDir() {
			MatchFilesTypeDir  (id, n, file)
		} else {
			MatchFilesTypeFile (id, n, file)
		}
		strmTime := timeOperationModule.GetCustomTime(filepath.Join(n.Source,file.Name()));
		timeOperationModule.ChangeCustomTime(filepath.Join(n.Destination,file.Name()), strmTime)
	 }
	return true
}

// Funcion does define repeat gorutine at the layer processig 
func processRepeatingBlock (id int, n FileScheme){
	wg.Add(1)
    i := id //???
	go func() {
		defer wg.Done()
            m.CopyFiles(i, n)
	}()
}

// Doing processing recursive call at every block by four element
// "residue of a modulo m"  https://t5k.org/glossary/page.php?sort=Residue
func blockProcessing(start int, n FileScheme) {
	a := start + 1
	if (start <= (count_max_record + 1)) {
		if (start%(block_size + 1) == block_size) {
			wg.Wait()
		} else {
		    processRepeatingBlock(a, n)
		}
		blockProcessing(a, n)
	}	
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
	//var m FileSizeTable
	//m.CopyFiles(n)
	count_max_record = getFileTotalCount(n)
	if (count_max_record < block_size) {block_size = count_max_record} 

	//------
    fmt.Printf( "Total  files = %d \n", count_max_record)
	//------
	
	blockProcessing(start_value, n)
	wg.Wait()
	return "Processing does compleat! \n"
}

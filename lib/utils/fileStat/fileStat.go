package fileStat

import (
	"io/ioutil"
	fileStatHeader "main/lib/header"
	. "main/lib/utils/getFileSize"
)

// *
// * get overall files size in the folder 
// *

type SizeFormatSlice []float64
type ExtendedFileScheme  fileStatHeader.FileScheme


type FS interface { 
    CalcSize(fullfilePath string, Size float64) float64
}

type FileSize struct{}

type FileStatMock struct{}

func (m FileSize) CalcSize(fullfilePath string, Size float64) float64{
	var  t FS = FileSize{}
    files, _ := ioutil.ReadDir(fullfilePath)
	        for _, file := range files { 
				if file.IsDir() {
					Size =  t.CalcSize((fullfilePath + "/" + file.Name()), Size)
				} else {
                    Size +=  GetFileSize((fullfilePath + "/" + file.Name())) 
			    }
			}
	return Size
}

func CSize(fullfilePath string) float64 {
	var mschemeUtils  ExtendedFileScheme
	var t FS = FileSize{}
	mschemeUtils = ExtendedFileScheme {
		    Source: fullfilePath,
		}
	return (t.CalcSize(mschemeUtils.Source, 0))
}

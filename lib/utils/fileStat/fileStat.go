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
    CalcSize(fullfilePath string, Size float64, CountFiles int) [2]float64
}

type FileSize struct{}

type FileStatMock struct{}

// Add CountFiles and return [2]array float for Size and Files Count
func (m FileSize) CalcSize(fullfilePath string, Size float64, CountFiles int ) [2]float64{
	var flArr [2]float64
	var  t FS = FileSize{}
	i := CountFiles
    files, _ := ioutil.ReadDir(fullfilePath)
	for _, file := range files {
		        i = i + 1
				if file.IsDir() {
					ArrSize :=  t.CalcSize((fullfilePath + "/" + file.Name()), Size, i)
					Size += ArrSize[0]
				} else {
                    Size += GetFileSize((fullfilePath + "/" + file.Name()))
 			    }
			}
	flArr[0]= Size
	flArr[1]= float64(i) 
	return flArr
}

func CSize(fullfilePath string) [2]float64 {
	var mschemeUtils  ExtendedFileScheme
	var t FS = FileSize{}
	mschemeUtils = ExtendedFileScheme {
		    Source: fullfilePath,
		}
	return t.CalcSize(mschemeUtils.Source, 0, 0)
}

package getFileSize

import (
	"os"
)

type SizeFormat float64
func GetFileSize (FullPathWithFile string ) float64 {
	file, err := os.Open(FullPathWithFile); if err != nil {panic(err)}
	currentPosition, _ := file.Seek(0, 2)
	defer file.Close()
	return (float64(currentPosition)/1024/1024)
}





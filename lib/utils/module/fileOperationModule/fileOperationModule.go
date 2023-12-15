package fileOperationModule

import (
     "io/ioutil"
	 "fmt"
	 "os"
	 "log"
)
// TODO Поправить 
func CopyFile (src string, dst string) {
    output, err := ioutil.ReadFile(src)
    if err != nil {
         fmt.Println("copy from " + fmt.Sprint(err) + " to " + string(output))
	}
	fileStat, err := os.Stat(src)
 	if err != nil {
		log.Fatal(err)
	}
    ioutil.WriteFile(dst, output, fileStat.Mode())
}

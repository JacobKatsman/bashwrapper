package jpgOperationModule

import (
	"os/exec"
	"fmt"
	"path/filepath"
	"regexp"
	. "main/lib/header"
)

func regexpDest (fileName string) string {
	 var re = regexp.MustCompile(`\(`)
	 fileName = re.ReplaceAllString(fileName, `\(`)
     var re1 = regexp.MustCompile(`\)`)
     fileName = re1.ReplaceAllString(fileName, `\)`)
return fileName
}


func SimpleJpg(n FileScheme, fileName string) bool{
	 fileName = regexpDest(fileName)
	 cpCmd := exec.Command("bash", "-c", "jpegoptim -p --strip-all --all-progressive -ptm85 " + filepath.Join(n.Destination,fileName))
	 output, err := cpCmd.CombinedOutput()
	 if err != nil {
         fmt.Println(" error " + fmt.Sprint(err) + ": " + string(output))
		  return false
     }
	return true
}

package pdfOperationModule

import (
	"os/exec"
	"fmt"
	"strings"
	"path/filepath"
	. "main/lib/header"
	timem "main/lib/utils/module/timeOperationModule"
)

func SimplePdf(n FileScheme, fileName string) bool{

	//fmt.Println("\nDestination  :=" + filepath.Join(n.Destination,fileName))
	//fmt.Println("source       :=" + filepath.Join(n.Source,fileName))
	//fmt.Println("fileName     :=" + fileName)  
	
	strmTime := timem.GetCustomTime(filepath.Join(n.Source,fileName));
	extraCmds := []string {
	    "ghostscript",
	 	"-sDEVICE=pdfwrite",
	 	"-dCompatibilityLevel=1.3",
	 	"-dPDFSETTINGS=/ebook",
        "-dNOPAUSE",
        "-dQUIET",
        "-dBATCH",
	    "-r300x300",
	    "-sOutputFile=" + filepath.Join(n.Destination,fileName)+" ",
		" "+filepath.Join(n.Source,fileName)+ " ",
    }

	justString := strings.Join(extraCmds," ")
	cmCmd := exec.Command("/bin/bash","-c", justString) 
	output, err := cmCmd.CombinedOutput()
 	 if err != nil {
         fmt.Println(" error " + fmt.Sprint(err) + ": " + string(output))
         return false
     }
	timem.ChangeCustomTime(filepath.Join(n.Destination,fileName), strmTime)
	//fmt.Println(string(output))
	return true
}

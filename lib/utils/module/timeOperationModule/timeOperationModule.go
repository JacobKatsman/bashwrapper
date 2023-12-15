package timeOperationModule

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func GetCustomTime(fileName string) time.Time {
 		fi, err := os.Stat(fileName)
         if err != nil {
             fmt.Println(err)
			 fmt.Println("Operation Error getCustomTime"); 
         }
         mTime := fi.Sys().(*syscall.Stat_t).Mtim
 		return time.Unix(mTime.Unix())
}	

func ChangeCustomTime(fileName string, strmTime time.Time)  {
 	    err := os.Chtimes(fileName, strmTime, strmTime)
 	    if err != nil {
 			fmt.Println(err)
 			fmt.Println("Operation Error changeCustomTime");
	    }
}

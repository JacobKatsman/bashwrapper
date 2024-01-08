package main

import (
	"fmt"
	"time"
	"os"
	"log"
	"io"
	"flag"
	"path/filepath"
	fileStat "main/lib/utils/fileStat"
	fileExtTableStat "main/lib/utils/fileExtTableStat"
	fileChangeCopyFiles "main/lib/utils/fileChangeCopyFiles"
	. "main/lib/header"
)

// **
// *  Copyright (c) [2023] Alexey Chikilevskiy (aka "JZKatsman") email: call89269081096@gmail.com
// *  Licensed under the Apache License, Version 2.0 (the «License»);
// *
// *  The above copyright notice and this permission notice shall be included in all
// *  copies or substantial portions of the Software.
//  
// *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// *  SOFTWARE.
// **

  // TODO struct 
  func timeChecker(n FileScheme) {
    defer duration(track("foo"))


	m := fileStat.CSize(n.Source)
	
    source := fmt.Sprintf("Source Overall file size = %.2fMB | Total FILES = %d (all) \n", m[0], int(m[1]))
	io.WriteString(os.Stdout, fmt.Sprintf("-------------------------------------------\n"))
	io.WriteString(os.Stdout, fileExtTableStat.CSize(n.Source))

	s := fileChangeCopyFiles.CFiles(n)


	v := fileStat.CSize(n.Destination)
    destination :=  fmt.Sprintf("Destination Overall size = %.2fMB | Total FILES = %d (all) \n",v[0], int(v[1]))
	
	io.WriteString(os.Stdout, fmt.Sprintf("--------------------------------------------\n"))
	io.WriteString(os.Stdout, fileExtTableStat.CSize(n.Destination))

	io.WriteString(os.Stdout, s)
	io.WriteString(os.Stdout, source)
	io.WriteString(os.Stdout, destination)

	
    }

    func track(msg string) (string, time.Time) {
      return msg, time.Now()
    }

    func duration(msg string, start time.Time) {
	Duration := time.Since(start) 
	log.Printf(" %v\n", Duration)
}

func RemoveContents(dir string) error {
    files, err := filepath.Glob(filepath.Join(dir, "*"))
    if err != nil {
        return err
    }
    for _, file := range files {
        err = os.RemoveAll(file)
        if err != nil {
            return err
        }
    }
    return nil
}


func main() {
	var source string
	var destination string

	flag.StringVar(&source, "source", "input_dir", "Directory")
	flag.StringVar(&destination, "destination", "output_dir", "Directory")
	flag.Parse()

	fmt.Println("Input folder path  :=" + source)
	fmt.Println("Output folder path :=" + destination)

	var err = RemoveContents(destination)
	if err != nil {
        return
    }
 	
    var s FileScheme = FileScheme {
		  Source:     source,
		  Destination:  destination,
	}
	timeChecker(s)
}

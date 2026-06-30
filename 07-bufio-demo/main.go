// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {

// 	nf, _ := os.Create("file.txt")
// 	defer nf.Close()
// 	bw := bufio.NewWriter(nf)
// 	fmt.Println(bw.Available())
// 	nf.WriteString("Reda Mhmd\n")
// 	bw.WriteString("Eid Hassan\n")
// 	bw.Flush()
// 	fmt.Println("done!!")
// 	nf.Seek(0, io.SeekStart)
// 	br := bufio.NewReaderSize(nf, 1024*1024)

// 	bs := bufio.NewScanner(br)
// 	nf.WriteString("12")
// 	bs.Split(bufio.ScanWords)
// 	// nf.Seek(0, io.SeekStart)
// 	for bs.Scan() {
// 		fmt.Println(bs.Text())
// 	}
// 	if err := bs.Err(); err != nil {
// 		fmt.Fprintln(os.Stdout, "reading standard input:", err)
// 	}

// }

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	newfile2, _ := os.Open("temp.txt")
	
	c, err := io.ReadAll(newfile2)
	fmt.Println("count: ", len(c), "err: ", err)
	ss, _ := os.Stat("temp.txt")
	fmt.Println(ss.Size())
}

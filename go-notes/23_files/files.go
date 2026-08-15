package main

import (
	"bufio"
	"fmt"
	"os"
)



func main(){
	// // f,err := os.Open("example.txt")

	// // if err != nil {
	// // 	panic(err)
	// // }

	// // fileInfo, err := f.Stat()

	// // fmt.Println(fileInfo.Name())
	// // fmt.Println(fileInfo.Size())
	// // fmt.Println(fileInfo.IsDir())
	// // fmt.Println(fileInfo.Mode())

	// // read file
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()


	// In basic text files, 1 byte contains exactly 1 character (like a letter, a number, or a space).The word "hello" is exactly 5 bytes long because it has 5 letters.Your example "hello golang" is exactly 12 bytes long (5 letters for hello + 1 space + 6 letters for golang).
	// // buffer is an array of bytes, a temporary storage in memory
	// buf := make([]byte, 12)
	// d, err := f.Read(buf)  // d is not the total size of the file. d is the number of items Go actually managed to put into your 12-byte bucket during that one specific read operation.Think of d as "Data successfully grabbed right now."Let’s trace your 50-byte file example across its lifecycles to see exactly what d and buf look like at each step:
	// Run 1:Your bucket (buf) has room for 12. Go reaches into the file and grabs 12 characters.d becomes 12.Your bucket is completely full of real data.
	// Run 2, 3, and 4:Go grabs 12 characters each time.d stays 12.
	// Run 5 (The Final Run):There are only 2 bytes left inside the file.Go looks at your 12-byte bucket. It places those final 2 characters into the first two slots of the bucket.d becomes 2. (Because it only read 2 new things).

	// if err != nil{
	// 	panic(err)
	// }
	//					12 always fixed
	//					|
	// for i := 0; i < len(buf); i++{   // in place of len(buf), we should put d
	// 	println("data", d, string(buf[i]))
	// }


	// --------

	// // readfile loads all the content in one go
	// data, err := os.ReadFile("example.txt")

	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// // read folders

	// dir, err := os.Open("../")
	// if err != nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo{
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// create a file
	f, err := os.Create("example2.txt")

	if err != nil{
		panic(err)
	}

	defer f.Close()

	f.WriteString("hi go")
	f.WriteString("Hello Language")

	// another way to insert data in file using bytes
	// bytes := []byte("hello golang")
	// f.Write(bytes)

	// -------------
	// reading from a file and streaming into different file, read and write to another file (streaming fashion)
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil{
			//  if error is not of end of file then panic other wise brake
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
		
	}
	writer.Flush() // writer.Flush() means "forcefully empty the temporary storage buffer and write all the data to its final destination right now."Every time your Go code communicates directly with your computer's hard drive or a network connection, it is physically slow. If you write a loop that sends text to a file character-by-character, and you talk to the hard drive 10,000 separate times, your application will crawl to a standstill.

	// The Solution: The "Holding Tank" (Buffer)To fix this speed problem, Go uses a buffered writer. When you write data, Go does not immediately send it to the file. Instead, it drops the text into a temporary holding tank in your computer's RAM (the buffer).RAM is incredibly fast, so your program finishes executing almost instantly.Go waits until that holding tank is 100% full before it automatically opens the gate and dumps the whole chunk onto the hard drive in one fast motion.

	// Where Flush() Comes In (The Missing Link)What happens if your program finishes running, but your holding tank is only half full?Because the tank isn't completely full yet, Go's automatic dump trigger never fires. The data will sit stranded inside your computer's RAM. If your program exits right there, that data is lost forever and will never appear in your file.By calling writer.Flush(), you manually override the system. You are screaming: "I don't care if the buffer isn't full yet; dump whatever is inside it into the file immediately!"

	// The data will continue to sit inside that RAM buffer until one of two things happens:
	// 1. The buffer gets 100% full: Go's bufio.Writer has a default internal capacity of 4,093 bytes. The moment you try to write the 4,094th byte, the buffer realizes it is full, automatically opens the gate, and dumps all 4,093 bytes onto the hard drive in one fast swoop.
	// 2. You call writer.Flush(): At the very end of your program, your file might only have a few bytes left, meaning the buffer is only partially filled. By calling writer.Flush(), you force Go to dump whatever is currently sitting in the RAM buffer straight into the destination file immediately.

	// fmt.Println("File written successfully")

	// ----------

	// delete file
	er := os.Remove("example.txt")

	if er != nil {
		panic(er)
	}

	fmt.Println(" File Deleted successfully")

}

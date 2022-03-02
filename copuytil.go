package copyutil

import (
	"os"
	"io"
	"github.com/cheggaaa/pb/v3"
)

func Copy(from string, to string, limit int, offset int) error {
	fileFrom, err := os.Open(from)
	fileFrom.Seek(offset, io.SeekStart)
	fileTo, err := os.Open(to)
	bar := pb.StartNew(count)
	_, err = io.CopyN(fileTo, fileFrom, limit)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
	if err != nil{
		return err
	}
	return nil
}
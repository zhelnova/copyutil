package copyutil

import (
	"io/ioutil"
	"os"
	"io"
	"github.com/cheggaaa/pb/v3"
	"time"
)

func Copy(from string, to string, limit int, offset int) error {
	fileFrom, err := os.Open(from)
	fileFrom.Seek(int64(offset), io.SeekStart)
	fileTo, err := os.Create(to)
	bFrom, err := ioutil.ReadAll(fileFrom)
	bar := pb.StartNew(limit)
	if len(bFrom) < limit {
		bar = pb.StartNew(len(bFrom))
	}
	defer func() {
		bar.Finish()
		fileFrom.Close()
		fileTo.Close()
	}()
	for i := 0; i < len(bFrom); i++ {
		if i == limit {
			break
		}
		bar.Increment()
		time.Sleep(time.Millisecond)
		_, err := fileTo.Write([]byte{bFrom[i]})
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

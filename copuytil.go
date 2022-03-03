package copyutil

import (
	"ioutil"
	"os"
	"io"
	"github.com/cheggaaa/pb/v3"
)

func Copy(from string, to string, limit int, offset int) error {
	fileFrom, err := os.Open(from)
	fileFrom.Seek(int64(offset), io.SeekStart)
	fileTo, err := os.Create(to)
	bFrom, err := ioutil.ReadAll(fileFrom)
	if len(bFrom) > limit {
		bar := pb.StartNew(bFrom)
	} else {
		bar := pb.StartNew(limit)
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
		_, err := fileTo.Write(bFrom[i])
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jpillora/GoSungrow/Only"
)

// FileRead Retrieves data from a local file.
func FileRead(fn string, ref interface{}) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		err = json.NewDecoder(f).Decode(&ref)
	}

	// for range Only.Once {
	//	fn := ep.GetFilename()
	//	if err != nil {
	//		break
	//	}
	//
	//	ret, err = os.FileRead(fn)
	//	if err != nil {
	//		break
	//	}
	// }

	return err
}

// FileWrite Saves data to a file path.
func FileWrite(fn string, ref interface{}, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()
		err = json.NewEncoder(f).Encode(ref)

		// fn := ep.GetFilename()
		// if err != nil {
		//	break
		// }
		//
		// err = os.FileWrite(fn, data, perm)
		// if err != nil {
		//	break
		// }
	}

	return err
}

// PlainFileRead Retrieves data from a local file.
func PlainFileRead(fn string) ([]byte, error) {
	var data []byte
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		data, err = ioutil.ReadAll(f)
	}

	return data, err
}

// PlainFileWrite Saves data to a file path.
func PlainFileWrite(fn string, data []byte, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			break
		}
		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		_, err = f.Write(data)
	}

	return err
}

// FileRemove Removes a file path.
func FileRemove(fn string) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			err = nil
			break
		}
		if err != nil {
			break
		}
		if f.IsDir() {
			err = errors.New("file is a directory")
			break
		}

		err = os.Remove(fn)
	}

	return err
}

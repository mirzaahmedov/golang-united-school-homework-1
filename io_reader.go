package homework

import (
	"io"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/
func SeekTillHalfOfString(strReader *strings.Reader) string {
	str, err := io.ReadAll(strReader)
        if err != nil {
          return ""
        }

        if len(str) % 2 == 0 {
          strReader.Seek(int64(len(str)/2), io.SeekStart)
        } else {
          strReader.Seek(int64(len(str)/2 + 1), io.SeekStart)
        }

	return ""
}

/*
ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks
*/
func ReaderSplit(strReader *strings.Reader, n int) []string {
	str, err := io.ReadAll(strReader)
	if err != nil {
		return []string{}
	}

	res := make([]string, len(str)/n+1)

	for i := 0; i < len(str); i += n {
		res = append(res, string(str[i:i+n]))
	}

	return res
}

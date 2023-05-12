/*
 * @Author: fyfishie
 * @Date: 2023-05-08:23
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-09:09
 * @@email: fyfishie@outlook.com
 * @Description: :)
 */
package store

import (
	"bufio"
	"os"
)

// raw loader load data as raw format (bytes,string .etc)
type RawLoader struct {
	rfi      *os.File
	rdr      *bufio.Reader
	dataPath string
}

func NewRawLoader(dataPath string) *RawLoader {
	return &RawLoader{dataPath: dataPath}
}
func (l *RawLoader) Open() (*RawLoader, error) {
	rfi, err := os.OpenFile(l.dataPath, os.O_RDONLY, 0000)
	if err != nil {
		return l, err
	}
	l.rfi = rfi
	l.rdr = bufio.NewReader(rfi)
	return l, nil
}

// func (l *RawLoader) NextStrings(number int) []string {
// 	scaner := bufio.NewScanner(l.rfi)
// 	scaner.Split(bufio.ScanLines)
// 	// res := []string{}
// 	for i := 0; i < number; i++ {
// 		scaner.Text()
// 		// line,err:=l.rdr.ReadString('\n')
// 		// if err != nil {
// 		// 	return nil, err
// 		// }
// 	}
// }

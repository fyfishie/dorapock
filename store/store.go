/*
 * @Author: fyfishie
 * @Date: 2023-05-08:10
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-08:11
 * @@email: fyfishie@outlook.com
 * @Description: :)
 */
package store

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Loader[T any] struct {
	dataPath string
	rdr      *bufio.Reader
}

func NewLoader[T any](dataPath string) *Loader[T] {
	return &Loader[T]{dataPath: dataPath}
}
func (l *Loader[T]) Open() error {
	rfi, err := os.OpenFile(l.dataPath, os.O_RDONLY, 0000)
	if err != nil {
		return err
	}
	defer rfi.Close()
	l.rdr = bufio.NewReader(rfi)
	return nil
}
func (l *Loader[T]) Next(number int) ([]*T, error) {
	next := []*T{}
	for i := 0; i < number; i++ {
		line, _, err := l.rdr.ReadLine()
		if err != nil {
			return next, err
		}
		var item T
		err = json.Unmarshal(line, &item)
		if err != nil {
			return next, err
		}
		next = append(next, &item)
	}
	return next, nil
}
func LoadAny[T any](dataPath string) ([]*T, error) {
	rfi, err := os.OpenFile(dataPath, os.O_RDONLY, 0000)
	if err != nil {
		return nil, err
	}
	defer rfi.Close()
	rdr := bufio.NewReader(rfi)
	res := []*T{}
	for {
		line, _, err := rdr.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		var item T
		err = json.Unmarshal(line, &item)
		if err != nil {
			return nil, err
		}
		res = append(res, &item)
	}
	return res, nil
}

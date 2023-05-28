/*
 * @Author: fyfishie
 * @Date: 2023-05-08:10
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-28:17
 * @@email: fyfishie@outlook.com
 * @Description: :)
 */
package store

import (
	"bufio"
	"encoding/json"
	"os"
)

type Loader[T any] struct {
	dataPath string
	rfi      *os.File
	scaner   *bufio.Scanner
}

func NewLoader[T any](dataPath string) *Loader[T] {
	return &Loader[T]{dataPath: dataPath}
}

func (l *Loader[T]) Open() error {
	rfi, err := os.Open(l.dataPath)
	if err != nil {
		panic(err.Error())
	}
	scaner := bufio.NewScanner(rfi)
	l.rfi = rfi
	l.scaner = scaner
	return nil
}

func (l *Loader[T]) NextMany(number int) ([]T, error) {
	next := []T{}
	for i := 0; i < number; i++ {
		if !l.scaner.Scan() {
			break
		}
		line := l.scaner.Bytes()
		var item T
		err := json.Unmarshal(line, &item)
		if err != nil {
			return next, err
		}
		next = append(next, item)
	}
	return next, nil
}

func (l *Loader[T]) Close() {
	l.rfi.Close()
}

func LoadAny[T any](dataPath string) ([]T, error) {
	rfi, err := os.Open(dataPath)
	if err != nil {
		panic(err.Error())
	}
	defer rfi.Close()
	scaner := bufio.NewScanner(rfi)
	res := []T{}
	for scaner.Scan() {
		line := scaner.Bytes()
		var item T
		err = json.Unmarshal(line, &item)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

// you'd must confirm that there exists next line by Loader.HasNext()
func (l *Loader[T]) Next() T {
	var item T
	line := l.scaner.Bytes()
	err := json.Unmarshal(line, &item)
	if err != nil {
		return item
	}
	return item
}

func (l *Loader[T]) HasNext() bool {
	return l.scaner.Scan()
}

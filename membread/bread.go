/*
 * @Author: fyfishie
 * @Date: 2023-05-12:16
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-12:16
 * @@email: fyfishie@outlook.com
 * @Description: :)
 */
package membread

import (
	"bufio"
	"fmt"
	"os"
)

type MemBread struct {
	memory     string
	targetPath string
	User       string
}

func NewMemBread(targetPath, user string) *MemBread {
	return &MemBread{targetPath: targetPath, User: user}
}
func (b *MemBread) Copy() error {
	rfi, err := os.OpenFile(b.targetPath, os.O_RDONLY, 0000)
	if err != nil {
		return err
	}
	defer rfi.Close()
	scaner := bufio.NewScanner(rfi)
	for scaner.Scan() {
		b.memory = b.memory + scaner.Text()
	}
	return nil
}
func (b *MemBread) Aowu() error {
	fmt.Println("init paste module...")
	fmt.Println("init succeed!")
	fmt.Println("paste memory into her cute mind...")
	fmt.Println("aowu~aowu~aowu~")
	fmt.Println("may you have nice score!")
	return nil
}

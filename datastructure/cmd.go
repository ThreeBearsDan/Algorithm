/**
 * Created by GoLand.
 * User: cruise
 * Date: 2021/6/5
 * Time: 8:21 下午
 */
package main

import (
	"datastructure/rbtree"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	random := flag.Bool("rand", false, "-rand")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)

	var j int
	for i := 0; i < 31; i++ {
		if *random {
			j = rand.Intn(100)
		} else {
			j = i
		}
		rbtree.Root = rbtree.Put(rbtree.Root, strconv.Itoa(j), strconv.Itoa(j))
		rbtree.DrawTree(rbtree.Root)
		<-ch
		fmt.Printf("\x1bc")
	}
	rbtree.DrawTree(rbtree.Root)
}
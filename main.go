package main

import (
	"fmt"
	"hash"
	"math/rand"
"time"
	"github.com/spaolacci/murmur3"
)

var mHashser hash.Hash32
var seed uint32

func init() {
  rand.Seed(time.Now().UnixNano())
	mHashser = murmur3.New32WithSeed(rand.Uint32())
}

func murmurHash(key string, size uint32) int {
	mHashser.Write([]byte(key))
	result := mHashser.Sum32() % size
	mHashser.Reset()
	return (int)(result)

}

type Bloomfilter struct {
	filter []uint8
	size   uint32
}

func newBloomFilter(size uint32) Bloomfilter {
	return Bloomfilter{
		make([]uint8, size/8),
		size,
	}
}

func (b *Bloomfilter) exist(key string) bool {
	idx := murmurHash(key, b.size)
	bytePosition := idx / 8
	bitPosition := idx % 8
	return (b.filter[bytePosition] & (1 << bitPosition)) > 0
}

func (b *Bloomfilter) add(key string) {
	idx := murmurHash(key, b.size)
	fmt.Println(idx)
	bytePosition := idx / 8
	bitPosition := idx % 8
	b.filter[bytePosition] = b.filter[bytePosition] | (1 << bitPosition)
}

func main() {
	bloomFilter := newBloomFilter(16)
	dataSet := []string{"a", "b", "c", "d", "e","f","g","h"}
	for _, key := range dataSet {
		bloomFilter.add(key)
	}
	dataSet = append(dataSet, "x")
	fmt.Println(bloomFilter.filter)
	for _, key := range dataSet {
		fmt.Println(bloomFilter.exist(key))
	}
}

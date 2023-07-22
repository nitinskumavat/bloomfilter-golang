package main

import (
	"fmt"
	"hash"

	"github.com/spaolacci/murmur3"
)

var mHashser hash.Hash32
var seed uint32

func init() {
	mHashser = murmur3.New32WithSeed(seed)
}

func murmurHash(key string, size uint32) int {
	mHashser.Write([]byte(key))
	result := mHashser.Sum32() % size
	mHashser.Reset()
	return (int)(result)

}

type Bloomfilter struct {
	filter []bool
	size   uint32
}

func newBloomFilter(size uint32) Bloomfilter {
	return Bloomfilter{
		make([]bool, size),
		size,
	}
}

func (b *Bloomfilter) exist(key string) bool {
	idx := murmurHash(key, b.size)
	return b.filter[idx]
}

func (b *Bloomfilter) add(key string) {
	idx := murmurHash(key, b.size)
	b.filter[idx] = true
}

func main() {
	seed = 10
	bloomFilter := newBloomFilter(16)
	dataSet := []string{"a", "b","c","d"}
	for _, key := range dataSet {
		bloomFilter.add(key)
	}
	dataSet = append(dataSet, "x")
	fmt.Println(bloomFilter.filter)
	for _, key := range dataSet {
		fmt.Println(bloomFilter.exist(key))
	}
}

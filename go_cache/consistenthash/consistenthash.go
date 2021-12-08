package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 实现一致性哈希算法

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

type Map struct {
	hash Hash
	// 虚拟节点倍数
	replicas int
	// 哈希环
	keys []int  // sorted
	// 虚拟节点与真实节点的隐射，键是虚拟节点哈希值，值是真实节点名称
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash: fn,
		hashMap: make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 虚拟节点名称是strconv.Itoa(i) + key
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	// 计算key对应的哈希值
	hash := int(m.hash([]byte(key)))
	// 二分查找对应虚拟节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 如果idx==len(m.keys),说明应该选择m.keys[0]
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

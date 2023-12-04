package lru

import (
	"container/list"
)

// Cache LRU缓存。并不是并发安全的
type Cache struct {
	maxBytes int64                    //是允许使用的最大内存
	nbytes   int64                    //当前已经使用的最大内存
	ll       *list.List               //双向链表list.List
	cache    map[string]*list.Element //键是字符串，值是双向链表中对应节点的指针
}

// 双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// Value 所占用的字节大小
type Value interface {
	Len() int
}

// New 实例化
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	//用于设置在缓存中的某个条目被删除时所执行的回调函数。
	//第一个参数是一个字符串类型键（key），表示被删除的缓存条目的键。
	//第二个参数是一个 Value 类型，表示被删除的缓存条目的值。
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

// Get 查找键的值
func (c *Cache) Get(key string) (value Value, ok bool) {

	if ele, ok := c.cache[key]; ok {
		//移动到队首
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) //如果断言成功，kv将保存entry类型的指针值；如果断言失败，会触发运行时错误。
		return kv.value, true
	}
	return
}

// RemoveOldest 移除最近最少访问的值
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
	}
}

// Add 新增/修改
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		//将节点加入队列首部 ele指向双向链表的指针
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len 缓存的数量
func (c *Cache) Len() int {
	return c.ll.Len()
}

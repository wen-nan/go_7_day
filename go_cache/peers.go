package go_cache

// PeerPicker PickPeer方法用于根据传入的key，选择相应的节点PeerGetter
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter Get方法从对应的group获取缓存值
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}

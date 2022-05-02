package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type CenterEtcd struct {
	Cli       *clientv3.Client
	LeaseID   clientv3.LeaseID
	AliveChan <-chan *clientv3.LeaseKeepAliveResponse
	Node      []Node
}

// Node key:/ProjectName/ChildNode val:Target
type Node struct {
	ProjectName string
	ChildNode   string
	Target      string
}

// NewCenterEtcd 对外调用，获取一个初始化好的的CenterEtcd，它可以调用大多数方法
// 但要求select其之下的aliveChan以便与etcd服务器建立心跳连接
// 一个CenterEtcd共同使用一个心跳
func NewCenterEtcd(endpoint string) (c *CenterEtcd) {
	c = &CenterEtcd{
		Node: make([]Node, 24),
	}
	client, err := newClient(endpoint)
	if err != nil {
		return
	}
	c.Cli = client
	err = c.newLease(5)
	if err != nil {
		return
	}
	alive, err := c.Cli.KeepAlive(context.TODO(), c.LeaseID)
	if err != nil {
		return
	}
	c.AliveChan = alive
	return
}

// NewClient 与etcd服务器连接，返回一个client
func newClient(endpoint string) (client *clientv3.Client, err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		//log.Println(err)
		return
	}
	return
}

// NewLease 设置一个租约给CenterEtcd
func (s *CenterEtcd) newLease(ttl int64) error {
	lease, err := s.Cli.Grant(context.Background(), ttl)
	if err != nil {
		log.Println(err)
		return err
	}
	s.LeaseID = lease.ID
	return nil
}

// AddNode 添加一个KV ,应该有一个主项目名和节点名,对外可调用,如/user/add
func (s *CenterEtcd) AddNode(node *Node) error {
	s.Node = append(s.Node, *node)
	key := "/" + node.ProjectName + "/" + node.ChildNode
	_, err := s.Cli.Put(context.TODO(), key, node.Target, clientv3.WithLease(s.LeaseID))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// NewNode 创建一个节点,不添加进center
func NewNode(projectName, childName, target string) (n *Node) {
	n = &Node{
		ProjectName: projectName,
		ChildNode:   childName,
		Target:      target,
	}
	return
}

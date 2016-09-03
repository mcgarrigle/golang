package node

import "fmt"
import "golang/ip"

// -----------------------------------------------------------

type Node struct {
    Hostname string
    Eth0 ip.IPV4
    Eth1 ip.IPV4
}

func New(hostname string) Node {
  node := Node{}
  node.Hostname = hostname
  node.Eth0     = ip.New("127.0.0.1")
  return node
}

func (n Node) String() string {
    return fmt.Sprintf("%s: eth0 %s; eth1 %s", n.Hostname, n.Eth0, n.Eth1)
}

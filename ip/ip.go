package ip

import (
    "fmt"
    "strings"
    "strconv"
)

// -------------------------------

type IPV4 [4]byte

func New(s string) IPV4 {
    var list	[]string = strings.Split(s, ".")
    var octets	[]byte   = make([]byte, 0, 4)
    var v	string
    var n	int

    for _, v = range list {
        n, _ = strconv.Atoi(v)
	octets = append(octets, byte(n))
    }
    return IPV4{octets[0], octets[1], octets[2], octets[3]} 
}

func (ip IPV4) String() string {
    return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
}


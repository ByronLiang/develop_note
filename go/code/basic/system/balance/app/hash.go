package app

import (
    "fmt"
    "hash/crc32"
)

type HashBalance struct {

}

func NewHashBalance() *HashBalance {
    return &HashBalance{}
}

func (h HashBalance) DoBalance(i interface{}, key string) (*Instance, error)  {
    resources := i.(Resources)
    crcTable := crc32.MakeTable(crc32.IEEE)
    hashVal := crc32.Checksum([]byte(key), crcTable)
    //crc32.ChecksumIEEE([]byte(key))
    fmt.Println("hash val: ", int(hashVal), hashVal)
    index := int(hashVal) % len(resources)
    return resources[index], nil
}

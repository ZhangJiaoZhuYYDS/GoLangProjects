// @Author zhangjiaozhu 2023/11/25 9:39:00
package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}

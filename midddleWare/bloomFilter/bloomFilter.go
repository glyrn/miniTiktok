package bloomFilter

import (
	"github.com/willf/bloom"
	"math"
)

var BloomFilter *bloom.BloomFilter

func InitBloom() {
	// 预期存储元素数量
	expectedElements := 1000

	// 期望的误判率
	falsePositiveRate := 0.01

	// 计算位数组大小
	bitArraySize := int(-float64(expectedElements) * math.Log(falsePositiveRate) / math.Pow(math.Log(2), 2))

	// 计算哈希函数数量
	hashFunctionsCount := int(float64(bitArraySize) / float64(expectedElements) * math.Log(2))

	// 创建布隆过滤器
	BloomFilter = bloom.New(uint(bitArraySize), uint(hashFunctionsCount))
}

/*
	插入数据到数据库的操作

	1. 存到redis中
	2. 把对应实体的id存到布隆过滤器中
	3. 查数据
		3.1 查询是否在过滤器中
			是 - > 从redis中查，如果过期，从mysql查
			否 - > 返回空值，不查询

	视频 - > videoID

	用户 - > userID

	评论 - > commentID

	点赞 - > likeID

	关注 - > followID
*/

func Addelement2BL(element string) {
	BloomFilter.Add([]byte(element))
}

func BlExist(element string) bool {
	return BloomFilter.Test([]byte(element))
}

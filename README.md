# SingleNumberGo

實作一個函式 SingleNumberGo 找出給定整數陣列中只有出現過一次的數字, 假設給定的整數陣列只有元素只有一個是出現一次其他都是成對出現

## 參考解法

透過 XOR 的特性

知道 a XOR b XOR b = a, 若且為若 a, b 都是整數

```golang
func SingleNumber(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length%2 == 0 {
		return -1
	}
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
```
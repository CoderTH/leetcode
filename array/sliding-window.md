## Sliding window
### 简介
利用两个指针i, j，始终保持区间[i, j]时符合题目条件的区间

### 难点
* 主要是对符合题目条件的判断
* 边界问题
### 滑动窗口模版
```
//根据题意写判断函数
func check(n Type) {/*statement*/}
//形参列表要根据题意变换
func slidingWindow(nums []int) {
    n := len(nums)
    //使用i指针遍历整个数组
    for i, j := 0, 0; i < n; i++ {
        //调整j指针使[i, j]符合题意
        for j <= i && check() {
            /*statement*/
            j++
        }
    }
}
```
>对于任何一个滑动窗口我们基本都可以套用上面模版
### leetcode相关题目
#### 209. 长度最小的子数组
![](https://i.loli.net/2021/11/11/piKgLEBATIdHbt1.png)

> 其实我们当我们看到连续子数组时我们就应该知道可以使用滑动窗口来解决问题，其中的条件为数组和>=target这一点就很重要，其他的我们套用模版，直接看代码

```
func minSubArrayLen(target int, nums []int) int {
    start,end := 0,0
    ans := math.MaxInt32
    sum := 0
    for end<len(nums){
        sum+=nums[end]
        for sum>=target{
            ans = min(ans,end-start+1)
            sum -= nums[start]
            start++
        }
        end++
    }
    if ans == math.MaxInt32{
        return 0
    }
    return ans
}

func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}
```

### leetcode相关题目
#### 904. 水果成篮
![](https://i.loli.net/2021/11/11/UvKsajQhkiyu34e.png)

> 有一说一，leetcode不讲武德，这题目乍一看真的不容易看懂，在这里简单的介绍一下：给我们一个数组，这个数组中的每一个元素都代表了一颗树，（例如：[1,2,1,2,3,2],1代表苹果树，2代表香蕉树，3代表....），我们有两个篮子，每一个篮子可以装同一种果实，并且不限大小，要我们求出连续相邻的树中我们最多能取多少果实。题目意思大概了解以后其实就简单了，很明显其实就是连续子数组问题，而条件就是我们篮子大小不能超过2，看代码：

```
func totalFruit(fruits []int) int {
    hashmap := make(map[int]int, 3)
    start,end,res := 0,0,0
    for end <len(fruits){
        hashmap[fruits[end]]++
        for start<=end&&len(hashmap)==3{
            hashmap[fruits[start]]--
            if hashmap[fruits[start]] == 0{
                delete(hashmap, fruits[start])
            }
            start++
        }
        res = max(res, end-start+1)
        end++
    }
    return res
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
```

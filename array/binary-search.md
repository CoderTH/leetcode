## Binary Search
### 简介
二分查找其实比较简单，每次查找时比对数组中间值与target，如果相等返回当前index，如果不相同，在根据比对大小不同去左右区间再次查找
### 难点
* 首先得注意数组必须是有序的
* 开闭合区间问题-通常由两种做法
  * 前闭后闭
  * 前闭后开
### leetcode相关题目
#### 704. Binary Search
![](https://i.loli.net/2021/11/07/ihkgcaNoC4UKHtP.png)
-----
**前闭后闭：**
```
func search(nums []int, target int) int {
    l,r := 0,len(nums)-1
    for l<=r{
        mid := l+(r-l)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid]<target{
            l = mid + 1
        }else {
            r = mid - 1
        }
    }
    return -1
}
```

**前闭后开:**
```
func search(nums []int, target int) int {
    l,r := 0,len(nums)
    for l<r{
        mid := l+(r-l)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid]<target{
            l = mid+1
        }else {
            r = mid
        }
    }
    return -1
}
```

#### 35. Search Insert Position
![](https://i.loli.net/2021/11/07/s2qCmBEPdkIwT9A.png)
**代码：**
```
func searchInsert(nums []int, target int) int {
    l,r := 0,len(nums)-1
    for l<=r{
        mid := l+(r-l)/2
        if nums[mid]==target{
            return mid
        }else if nums[mid]<target{
            l = mid+1
        }else{
            r = mid-1
        }
    }
    return r
}
```
#### 34. Find First and Last Position of Element in Sorted Array

![](https://i.loli.net/2021/11/07/LACZx14X9s3OgrP.png)

**解析：**

>简单看下题目其实很简单，给一个有序的数组，再给定一个target，这个target在数组中可能多次出现，需要算出第一次出现与最后一次出现的下标位置，**有序的数组**可以让我们基本确定使用二分法来解决问题
* 首先使用二分法找到其中的一个target值
* 内嵌一个循环向两边扩散查找第一个位置与第二个位置

**代码如下：**

```
func searchRange(nums []int, target int) []int {
    l,r := 0,len(nums)-1
    start,end := -1,-1
    for l<=r{
        mid := l+(r-l)/2
        if nums[mid] == target{
            start,end = mid,mid
            for{
                if start-1 >= 0 && nums[start-1] == target {
                    start = start - 1
                } else if end+1 <= len(nums)-1 && nums[end+1] == target {
                    end = end + 1
                } else {
                    break
                }
            }
            break
        }else if nums[mid]<target{
            l = mid + 1
        }else{
            r = mid -1
        }
    }
    return []int{start,end}
}
```
#### 367. Valid Perfect Square

![](https://i.loli.net/2021/11/07/7mOhDkEjxTNCsQJ.png)

**解析：**
>比较简单，只需要判断0-num这个区间有没有 n*n==num的情况，我们一样可以使用二分法

**代码：**
```
func isPerfectSquare(num int) bool {
    l,r := 0,num
    for l<=r {
        mid := l+(r-l)/2
        if mid * mid == num{
            return true
        }else if  mid * mid < num{
            l = mid + 1
        }else{
            r = mid - 1
        }
    }
    return false
}
```

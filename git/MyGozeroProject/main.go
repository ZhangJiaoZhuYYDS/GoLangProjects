package main

func main() {
	nums := []int{1,2,3}
	subsets(nums)
}

func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
func s(nums []int) (ans [][]int)  {
	set := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums){
			ans = append(ans,append([]int(nil),set...))
			return
		}
		set = append(set,nums[i])
		dfs(i+1)
		set = set[:len(set)-1]
		dfs(i+1)
	}
}
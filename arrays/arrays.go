package arrays

func Sum(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func SumAll(numsList ...[]int) []int {
    sums := []int{}

    for _, nums := range numsList {
        sums = append(sums, Sum(nums...))
    }

    return sums
}

func SumAllTails(numsList ...[]int) []int {
    sums := []int{}

    for _, nums := range numsList {
        if len(nums) == 0 {
            nums = []int{0}
        }

        tail := nums[1:]
        sums = append(sums, Sum(tail...))
    }

    return sums
}

// TODO: Read https://go.dev/blog/slices-intro

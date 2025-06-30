import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
        return 1
    }
	fibonacciNums := make([]int, n+1)
	fibonacciNums[1] = 1
	if n < 2 {
		return fibonacciNums[n]
	}
	for i := 2; i <= n; i++ {
		fmt.Println(i, fibonacciNums)
		fibonacciNums[i] = fibonacciNums[i-1] + fibonacciNums[i-2]
	}
	fmt.Println(fibonacciNums)
	return fibonacciNums[n]
}
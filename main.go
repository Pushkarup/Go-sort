package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"
)

// Helper function to swap elements in an array.
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// findMin returns the minimum value in the array.
func findMin(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

// findMax returns the maximum value in the array.
func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

// Sorting algorithms

// Selection Sort
func selectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		swap(arr, i, minIdx)
	}
}

// Bubble Sort
func bubbleSort(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

// Insertion Sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

// Quick Sort
func quickSort(arr []int) {
	quickSortRecursive(arr, 0, len(arr)-1)
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSortRecursive(arr, low, pivot-1)
		quickSortRecursive(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

// Heap Sort
func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
}

// Counting Sort
func countingSort(arr []int) {
	max := findMax(arr)
	count := make([]int, max+1)
	output := make([]int, len(arr))

	for _, num := range arr {
		count[num]++
	}
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	copy(arr, output)
}

// Radix Sort
func radixSort(arr []int) {

	max := findMax(arr)
	for exp := 1; max/exp > 0; exp *= 10 {
		countSortByDigit(arr, exp)
	}
}

func countSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for _, num := range arr {
		index := (num / exp) % 10
		count[index]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}
	copy(arr, output)
}

// Comb Sort
func combSort(arr []int) {
	n := len(arr)
	gap := n
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}
		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				swap(arr, i, i+gap)
				sorted = false
			}
		}
	}
}

// Pigeonhole Sort
func pigeonholeSort(arr []int) {
	min := findMin(arr)
	max := findMax(arr)
	size := max - min + 1
	hole := make([][]int, size)

	for _, num := range arr {
		hole[num-min] = append(hole[num-min], num)
	}

	idx := 0
	for _, list := range hole {
		for _, num := range list {
			arr[idx] = num
			idx++
		}
	}
}

// Cycle Sort
func cycleSort(arr []int) {
	n := len(arr)
	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := arr[cycleStart]
		pos := cycleStart

		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}
		if pos == cycleStart {
			continue
		}

		for item == arr[pos] {
			pos++
		}
		arr[pos], item = item, arr[pos]

		for pos != cycleStart {
			pos = cycleStart
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
			}
			for item == arr[pos] {
				pos++
			}
			arr[pos], item = item, arr[pos]
		}
	}
}

// Cocktail Sort
func cocktailSort(arr []int) {
	n := len(arr)
	swapped := true
	start := 0
	end := n - 1

	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		end--
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				swapped = true
			}
		}
		start++
	}
}

// Strand Sort
func strandSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	out := []int{}
	for len(arr) > 0 {
		sublist := []int{arr[0]}
		arr = arr[1:]

		for i := 0; i < len(arr); {
			if arr[i] > sublist[len(sublist)-1] {
				sublist = append(sublist, arr[i])
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				i++
			}
		}
		out = merge(out, sublist)
	}
	return out
}

// Bitonic Sort
func bitonicSort(arr []int) {
	bitonicSortRecursive(arr, 0, len(arr), 1)
}

func bitonicSortRecursive(arr []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSortRecursive(arr, low, k, 1)
		bitonicSortRecursive(arr, low+k, k, 0)
		bitonicMerge(arr, low, cnt, dir)
	}
}

func bitonicMerge(arr []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			if (arr[i] > arr[i+k]) == (dir == 1) {
				swap(arr, i, i+k)
			}
		}
		bitonicMerge(arr, low, k, dir)
		bitonicMerge(arr, low+k, k, dir)
	}
}

// Pancake Sort
func pancakeSort(arr []int) {
	for curr_size := len(arr); curr_size > 1; curr_size-- {
		maxIndex := findMaxIndex(arr, curr_size)
		if maxIndex != curr_size-1 {
			flip(arr, maxIndex)
			flip(arr, curr_size-1)
		}
	}
}

func findMaxIndex(arr []int, n int) int {
	maxIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func flip(arr []int, i int) {
	start := 0
	for start < i {
		swap(arr, start, i)
		start++
		i--
	}
}

// BogoSort or Permutation Sort
func bogoSort(arr []int) {
	for !isSorted(arr) {
		rand.Shuffle(len(arr), func(i, j int) {
			swap(arr, i, j)
		})
	}
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// Gnome Sort
func gnomeSort(arr []int) {
	index := 0
	for index < len(arr) {
		if index == 0 {
			index++
		}
		if arr[index] >= arr[index-1] {
			index++
		} else {
			swap(arr, index, index-1)
			index--
		}
	}
}

// Sleep Sort
func sleepSort(arr []int) {
	done := make(chan bool, len(arr))
	for _, num := range arr {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Println(n)
			done <- true
		}(num)
	}
	for range arr {
		<-done
	}
}

// Structure Sorting (custom sort using sort.Slice)
func structureSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

// Stooge Sort
func stoogeSort(arr []int) {
	stoogeSortRecursive(arr, 0, len(arr)-1)
}

func stoogeSortRecursive(arr []int, l, h int) {
	if l >= h {
		return
	}
	if arr[l] > arr[h] {
		swap(arr, l, h)
	}
	if h-l+1 > 2 {
		t := (h - l + 1) / 3
		stoogeSortRecursive(arr, l, h-t)
		stoogeSortRecursive(arr, l+t, h)
		stoogeSortRecursive(arr, l, h-t)
	}
}

// Tag Sort
func tagSort(arr []int) ([]int, []int) {
	original := make([]int, len(arr))
	copy(original, arr)
	sort.Ints(arr)
	return original, arr
}

// Tree Sort
type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func treeSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	root := &treeNode{arr[0], nil, nil}
	for _, num := range arr[1:] {
		insert(root, num)
	}

	sortedArr := inorder(root, nil)
	copy(arr, sortedArr)
}

func insert(root *treeNode, value int) {
	if value < root.value {
		if root.left == nil {
			root.left = &treeNode{value, nil, nil}
		} else {
			insert(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &treeNode{value, nil, nil}
		} else {
			insert(root.right, value)
		}
	}
}

func inorder(root *treeNode, arr []int) []int {
	if root != nil {
		arr = inorder(root.left, arr)
		arr = append(arr, root.value)
		arr = inorder(root.right, arr)
	}
	return arr
}

// Odd-Even Sort / Brick Sort
func oddEvenSort(arr []int) {
	n := len(arr)
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 1; i <= n-2; i += 2 {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				isSorted = false
			}
		}
		for i := 0; i <= n-2; i += 2 {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				isSorted = false
			}
		}
	}
}

// 3-way Merge Sort
func mergeSort3Way(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	third := len(arr) / 3
	left := mergeSort3Way(arr[:third])
	middle := mergeSort3Way(arr[third : 2*third])
	right := mergeSort3Way(arr[2*third:])

	return merge3(left, middle, right)
}

func merge3(left, middle, right []int) []int {
	l, m, r := 0, 0, 0
	result := make([]int, len(left)+len(middle)+len(right))
	for i := 0; i < len(result); i++ {
		if l < len(left) && (m >= len(middle) || left[l] <= middle[m]) && (r >= len(right) || left[l] <= right[r]) {
			result[i] = left[l]
			l++
		} else if m < len(middle) && (r >= len(right) || middle[m] <= right[r]) {
			result[i] = middle[m]
			m++
		} else {
			result[i] = right[r]
			r++
		}
	}
	return result
}
func getArrayInput() []int {
	var n int
	fmt.Print("[INPUT] Total number of elements in Array: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("[INPUT] Enter the Elements: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func logo() {
	fmt.Printf("▒█▀▀█ ▒█░▒█ ▒█▀▀▀█ ▒█░▒█ ▒█░▄▀ ░█▀▀█ ▒█▀▀█   ▒█░▒█ ▒█▀▀█ ░█▀▀█ ▒█▀▀▄ ▒█░▒█ ▒█░░▒█ ░█▀▀█ ▒█░░▒█ \n")
	fmt.Printf("▒█▄▄█ ▒█░▒█ ░▀▀▀▄▄ ▒█▀▀█ ▒█▀▄░ ▒█▄▄█ ▒█▄▄▀   ▒█░▒█ ▒█▄▄█ ▒█▄▄█ ▒█░▒█ ▒█▀▀█ ▒█▄▄▄█ ▒█▄▄█ ▒█▄▄▄█ \n")
	fmt.Printf("▒█░░░ ░▀▄▄▀ ▒█▄▄▄█ ▒█░▒█ ▒█░▒█ ▒█░▒█ ▒█░▒█   ░▀▄▄▀ ▒█░░░ ▒█░▒█ ▒█▄▄▀ ▒█░▒█ ░░▒█░░ ▒█░▒█ ░░▒█░░ \n")
	fmt.Printf("                                                                       SORTING ALGORITHMS      \n")
}

func main() {
	clearConsole()
	logo()
	arr := getArrayInput()
	fmt.Printf("\n\n")
	fmt.Println("\t[Original array] : ", arr)

	var choice int
	fmt.Println("Choose a sorting algorithm:")
	fmt.Println("[01] Selection Sort    [13] Strand Sort")
	fmt.Println("[02] Bubble Sort       [14] Bitonic Sort")
	fmt.Println("[03] Insertion Sort    [15] Pancake Sort")
	fmt.Println("[04] Merge Sort        [16] BogoSort or Permutation Sort")
	fmt.Println("[05] Quick Sort        [17] Gnome Sort")
	fmt.Println("[06] Heap Sort         [18] Sleep Sort")
	fmt.Println("[07] Counting Sort     [19] Structure Sorting")
	fmt.Println("[08] Radix Sort        [20] Stooge Sort")
	fmt.Println("[09] Comb Sort         [21] Tag Sort (Original & Sorted)")
	fmt.Println("[10] Pigeonhole Sort   [22] Tree Sort")
	fmt.Println("[11] Cycle Sort        [23] Odd-Even Sort / Brick Sortv")
	fmt.Println("[12] Cocktail Sort     [24] 3-way Merge Sort")
	fmt.Printf("\n[INPUT] Your Choice : ")
	fmt.Scan(&choice)
	fmt.Printf("\n\n")

	switch choice {
	case 1:
		fmt.Printf("█▀ █▀▀ █░░ █▀▀ █▀▀ ▀█▀ █ █▀█ █▄░█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▄█ ██▄ █▄▄ ██▄ █▄▄ ░█░ █ █▄█ █░▀█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n²)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		selectionSort(arr)
	case 2:
		fmt.Printf("█▄▄ █░█ █▄▄ █▄▄ █░░ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄█ █▄█ █▄█ █▄█ █▄▄ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		bubbleSort(arr)
	case 3:
		fmt.Printf("█ █▄░█ █▀ █▀▀ █▀█ ▀█▀ █ █▀█ █▄░█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█ █░▀█ ▄█ ██▄ █▀▄ ░█░ █ █▄█ █░▀█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		insertionSort(arr)
	case 4:
		fmt.Printf("█▀▄▀█ █▀▀ █▀█ █▀▀ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█░▀░█ ██▄ █▀▄ █▄█ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n log n)\tO(n)")

		arr = mergeSort(arr)
	case 5:
		fmt.Printf("█▀█ █░█ █ █▀▀ █▄▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▀▀█ █▄█ █ █▄▄ █░█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n²)\t\tO(log n)")
		quickSort(arr)
	case 6:
		fmt.Printf("█░█ █▀▀ ▄▀█ █▀█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▀█ ██▄ █▀█ █▀▀   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n log n)\tO(1)")
		heapSort(arr)
	case 7:
		fmt.Printf("█▀▀ █▀█ █░█ █▄░█ ▀█▀ █ █▄░█ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄▄ █▄█ █▄█ █░▀█ ░█░ █ █░▀█ █▄█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n + k)\tO(n + k)\tO(n + k)\tO(k)")
		countingSort(arr)
	case 8:
		fmt.Printf("█▀█ ▄▀█ █▀▄ █ ▀▄▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▀▄ █▀█ █▄▀ █ █░█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(nk)\t\tO(nk)\t\tO(nk)\t\tO(n + k)")
		radixSort(arr)
	case 9:
		fmt.Printf("█▀▀ █▀█ █▀▄▀█ █▄▄   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄▄ █▄█ █░▀░█ █▄█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n² / 2^p)\tO(n²)\t\tO(1)")
		combSort(arr)
	case 10:
		fmt.Printf("█▀█ █ █▀▀ █▀▀ █▀█ █▄░█ █░█ █▀█ █░░ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▀▀ █ █▄█ ██▄ █▄█ █░▀█ █▀█ █▄█ █▄▄ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n + N)\tO(n + N)\tO(n + N)\tO(n + N)")
		pigeonholeSort(arr)
	case 11:
		fmt.Printf("█▀▀ █▄█ █▀▀ █░░ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄▄ ░█░ █▄▄ █▄▄ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n²)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		cycleSort(arr)
	case 12:
		fmt.Printf("█▀▀ █▀█ █▀▀ █▄▀ ▀█▀ ▄▀█ █ █░░   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄▄ █▄█ █▄▄ █░█ ░█░ █▀█ █ █▄▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		cocktailSort(arr)
	case 13:
		fmt.Printf("█▀ ▀█▀ █▀█ ▄▀█ █▄░█ █▀▄   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▄█ ░█░ █▀▄ █▀█ █░▀█ █▄▀   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n²)\t\tO(n²)\t\tO(n²)\t\tO(n)")
		arr = strandSort(arr)
	case 14:
		fmt.Printf("█▄▄ █ ▀█▀ █▀█ █▄░█ █ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄█ █ ░█░ █▄█ █░▀█ █ █▄▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log² n)\tO(n log² n)\tO(n log² n)\tO(n log² n)")
		bitonicSort(arr)
	case 15:
		fmt.Printf("█▀█ ▄▀█ █▄░█ █▀▀ ▄▀█ █▄▀ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▀▀ █▀█ █░▀█ █▄▄ █▀█ █░█ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(n)")
		pancakeSort(arr)
	case 16:
		fmt.Printf("█▄▄ █▀█ █▀▀ █▀█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄█ █▄█ █▄█ █▄█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO((n+1)!)\tO(∞)\t\tO(1)")
		bogoSort(arr)
	case 17:
		fmt.Printf("█▀▀ █▄░█ █▀█ █▀▄▀█ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄█ █░▀█ █▄█ █░▀░█ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		gnomeSort(arr)
	case 18:
		fmt.Printf("█▀ █░░ █▀▀ █▀▀ █▀█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▄█ █▄▄ ██▄ ██▄ █▀▀   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n log n)\tO(n log n)\tO(n)")
		sleepSort(arr)
	case 19:
		fmt.Printf("█▀ ▀█▀ █▀█ █░█ █▀▀ ▀█▀ █░█ █▀█ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▄█ ░█░ █▀▄ █▄█ █▄▄ ░█░ █▄█ █▀▄ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n log n)\tO(log n)")
		structureSort(arr)
	case 20:
		fmt.Printf("█▀ ▀█▀ █▀█ █▀█ █▀▀ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("▄█ ░█░ █▄█ █▄█ █▄█ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n^(log 3 / log 1.5))\tO(n^(log 3 / log 1.5))\tO(n^(log 3 / log 1.5))\tO(n)")
		stoogeSort(arr)
	case 21:
		fmt.Printf("▀█▀ ▄▀█ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("░█░ █▀█ █▄█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n log n)\tO(n)")
		_, sorted := tagSort(arr)
		fmt.Println("Sorted array:", sorted)
	case 22:
		fmt.Printf("▀█▀ █▀█ █▀▀ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("░█░ █▀▄ ██▄ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n²)\t\tO(n)")
		treeSort(arr)
	case 23:
		fmt.Printf("█▀█ █▀▄ █▀▄ ▄▄ █▀▀ █░█ █▀▀ █▄░█   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█▄█ █▄▀ █▄▀ ░░ ██▄ ▀▄▀ ██▄ █░▀█   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n)\t\tO(n²)\t\tO(n²)\t\tO(1)")
		oddEvenSort(arr)
	case 24:
		fmt.Printf("█ █ █ ▄▄ █░█░█ ▄▀█ █▄█   █▀▄▀█ █▀▀ █▀█ █▀▀ █▀▀   █▀ █▀█ █▀█ ▀█▀\n")
		fmt.Printf("█ █ █ ░░ ▀▄▀▄▀ █▀█ ░█░   █░▀░█ ██▄ █▀▄ █▄█ ██▄   ▄█ █▄█ █▀▄ ░█░\n")
		fmt.Println("\nBest Time\tAverage Time\tWorst Time\tSpace")
		fmt.Println("--------------------------------------------------------")
		fmt.Println("O(n log n)\tO(n log n)\tO(n log n)\tO(n)")
		arr = mergeSort3Way(arr)
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Printf("\n")
	fmt.Println("\t[Sorted array] : ", arr)
}

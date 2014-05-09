package main
import(
  "github.com/diebels727/readline"
  "container/heap"
  "fmt"
  "time"
  "strconv"
  )

type HighHeap []int
type LowHeap []int

func (h HighHeap) Len() int           { return len(h) }
func (h HighHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h HighHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HighHeap) Push(x interface{}) {
  *h = append(*h, x.(int))
}

func (h *HighHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}


func (h *HighHeap) ExtractMin() int {
  old := *h
  x := old[0]
  *h = old[1:]
  heap.Init(h)
  return x
}


func (h LowHeap) Len() int           { return len(h) }
func (h LowHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h LowHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LowHeap) Push(x interface{}) {
  *h = append(*h, x.(int))
}

func (h *LowHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

func (h *LowHeap) ExtractMax() int {
  old := *h
  x := old[0]
  *h = old[1:]
  heap.Init(h)
  return x
}

func Rebalance() {
  if len(*l) > len(*h) {
    max := (*l).ExtractMax()
    heap.Push(h,max)
  }
  if len(*l) < len(*h) {  //len(*l) >= len(*h)
    min := (*h).ExtractMin()
    heap.Push(l,min)
  }
}

func ComputeMedian() {
  var median int
  total_len := len(*l)+len(*h)
  if ( total_len % 2 == 0) {
    median = (*l)[0]
  } else {
    if (len(*l) > len(*h)) {
      median = (*l)[0]
    } else {
      median = (*h)[0]
    }
  }
  fmt.Println("h: ",*h)
  fmt.Println("l: ",*l)
  fmt.Println("median: ",median)
  total += median
}

func runtime(fn func()) {
  start := time.Now()
  fn()
  elapsed := time.Since(start)
  fmt.Println("runtime: ",elapsed)
}

func Mapper(str string) {
  a,_ := strconv.ParseInt(str,10,0)
  v := int(a)

  if line == 0 {
    heap.Push(l,v)
    line++
  } else {
    if line == 1 {
      max := (*l)[0]
      if v > max {
        heap.Push(h,v)
        line++
      } else {
        heap.Push(l,v)
      }
    } else {
      max := (*l)[0]
      // min := (*h)[0]
      if v > max {
        heap.Push(h,v)
      } else {
        heap.Push(l,v)
      }
    }
  }
  // fmt.Println("---")
  // fmt.Println("h: ",*h)
  // fmt.Println("l: ",*l)
  Rebalance()
  ComputeMedian()

  //if minimum or maximum undef, handle
  //get minimum from high heap and maximum from low heap
  //compare minimum and maximum to a
  //if a > maximum in low heap
    //insert in to high heap
  //else if a < minimum in high heap
    //insert in to low heap

  //rebalance heaps

  //compute median
}

var h *HighHeap
var l *LowHeap
var line int
var total int

func main() {
  line = 0
  total = 0
  // h = &HighHeap{}
  // l = &LowHeap{}
  // heap.Init(h)
  // heap.Init(l)
  // fmt.Println(len(*h))
  // fmt.Println(len(*l))

  //1 2 3 6 7

  //6
  h = &HighHeap{}
  l = &LowHeap{}
  // readline.Map("Median.txt",Mapper)
  readline.Map("tc2.txt",Mapper)
  fmt.Println(total)
  modded := int64(total) % 10000
  fmt.Println(modded)
  // fmt.Println(len(*h))
  // fmt.Println(len(*l))
  // heap.Push(h,7)
  // heap.Push(h,8)
  // heap.Push(l,6)
  // Rebalance()
  // fmt.Println("low heap:",*l)
  // fmt.Println("high heap:",*h)
  // fmt.Println((*l)[0])
  // fmt.Println((*h)[0])
  // fmt.Println(h)
  // fmt.Println(*l)
}
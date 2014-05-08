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

func Balance(low_heap *LowHeap,high_heap *HighHeap) {
  l := *low_heap
  h := *high_heap
  l_len := len(l)
  h_len := len(l)
  if l_len < h_len {
    max := l.ExtractMax()
    heap.Push(&h,max)
  } else {  //l_len > high_len
    min := h.ExtractMin()
    heap.Push(&l,min)
  }
}

func runtime(fn func()) {
  start := time.Now()
  fn()
  elapsed := time.Since(start)
  fmt.Println("runtime: ",elapsed)
}

func Mapper(str string) {
  a,_ := strconv.ParseInt(str,10,0)
  //get heaps

  //if minimum or maximum undef, handle
  //get minimum from high heap and maximum from low heap
  //compare minimum and maximum to a
  //if a > maximum in low heap
    //insert in to high heap
  //else if a < minimum in high heap
    //insert in to low heap

  //rebalance heaps

  //compute median

  heap.Push(h,int(a))
}

var h *HighHeap
var l *LowHeap

func main() {
  h = &HighHeap{}
  l = &LowHeap{}
  heap.Init(h)
  heap.Init(l)
  readline.Map("Median.txt",Mapper)
}
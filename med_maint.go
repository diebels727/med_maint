package main
import(
  // "github.com/diebels727/readline"
  "container/heap"
  "fmt"
  "time")

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
  if l_len < high_len {
    max := l.ExtractMax()
    heap.Push(h,max)
  } else {  //l_len > high_len
    min := h.ExtractMin()
    heap.Push(l,min)
  }
}

func runtime(fn func()) {
  start := time.Now()
  fn()
  elapsed := time.Since(start)
  fmt.Println("runtime: ",elapsed)
}

// func Mapper(str string) {
//
// }


func main() {
  h := &HighHeap{4,5,6,1,3,4}
  heap.Init(h)
  heap.Push(h,7)
  heap.Push(h,-7)
  fmt.Println(h.ExtractMin())
  fmt.Println(h)
  // readline.Map("Median.txt",Mapper)

}
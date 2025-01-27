class NumberBinaryHeap {
  heap = new Array()
  heapSize = 0
  heapCap = 0

  constructor(heap) {
    if (heap) {
      this.heap = heap
      this.heapSize = heap.length
      this.heapCap = heap.length
    }
  }

  IsValueValid(value) {
    return !(value === undefined || value === null || Number.isNaN(value))
  }

  IsLess(idx1, idx2) {
    if (!this.heap[idx2]) return true

    return this.heap[idx1] <= this.heap[idx2]
  }

  Swap(idx1, idx2) {
    const bkp = this.heap[idx1]
    this.heap[idx1] = this.heap[idx2]
    this.heap[idx2] = bkp
  }

  // Bottom up
  Swim(idx) {
    let idxParent = parseInt((idx - 1) / 2)

    while (idx > 0 && this.IsLess(idx, idxParent)) {
      this.Swap(idxParent, idx)

      idx = idxParent
      idxParent = parseInt((idx - 1) / 2)
    }
  }

  // Top Down
  Sink(idx) {
    console.log(this.heap)
    const leftIdx = 2*idx + 1
    const rightIdx = 2*idx + 2
    let smallest = leftIdx

    if (rightIdx < this.heapSize && this.IsLess(rightIdx, leftIdx)) {
      smallest = rightIdx
    }

    if (smallest >= this.heapSize || !this.IsLess(smallest, idx)) return

    this.Swap(smallest, idx)
    this.Sink(smallest)
  }

  Add(item) {
    if (!this.IsValueValid(item)) throw new Error('Invalid Value')

    if (this.heapSize < this.heapCap) {
      this.heap[this.heapSize] = item
    } else {
      this.heap.push(item)
    }
    
    this.Swim(this.heapSize)
    this.heapSize++
  }

  Poll() {
    const first = this.heap[0]
    this.heap[0] = this.heap.at(-1)
    this.heapSize--
    this.heap[this.heapSize] = undefined

    this.Sink(0)

    return first
  }

  IsMinHeap(idxRoot = 0) {
    if (idxRoot >= this.heapSize) return true

    const leftIdx = 2 * idxRoot + 1
    const rightIdx = 2 * idxRoot + 2

    if (leftIdx < this.heapSize && !this.IsLess(idxRoot, leftIdx)) return false
    if (rightIdx < this.heapSize && !this.IsLess(idxRoot, rightIdx)) return false

    return this.IsMinHeap(leftIdx) && this.IsMinHeap(rightIdx)
  }
}

const heap = new NumberBinaryHeap([1, 2, 20, 4, 10])
console.log(heap.IsMinHeap(0))
heap.Add(0)
console.log(heap.heap);
console.log(heap.IsMinHeap(0))
heap.Poll()
console.log(heap.heap);
console.log(heap.IsMinHeap(0))

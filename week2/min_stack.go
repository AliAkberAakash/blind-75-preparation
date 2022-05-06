type MinStack struct {
    List []int
    TopIndex int
    Min int
}


func Constructor() MinStack {
    return MinStack{
        List: make([]int, 0),
        TopIndex: -1,
        Min: int(^uint(0) >> 1),
    }
}


func (this *MinStack) Push(val int)  {
  
    this.List = append(this.List,val)
    this.TopIndex++
    if val<this.Min {
        this.Min = val
    }
   
}


func (this *MinStack) Pop()  {
   
    this.List = this.List[:this.TopIndex]
    this.TopIndex--
    
    this.Min = int(^uint(0) >> 1)
    if(this.TopIndex>=0){
        for _,c := range this.List {
            if c<this.Min {
                this.Min = c
            }
        }
    }
   
}


func (this *MinStack) Top() int {
    
    return this.List[this.TopIndex]
}


func (this *MinStack) GetMin() int {
   
    return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

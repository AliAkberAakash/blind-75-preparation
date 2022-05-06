class MyQueue() {

    private var myStack = ArrayDeque<Int>()
    private var buffer = ArrayDeque<Int>()   
    
    var pos : Int = -1
    
    fun push(x: Int) {
        myStack.add(x)
        pos++
    }

    fun pop(): Int {
        var s : Int = 0
        while(!myStack.isEmpty()){
            buffer.add(myStack.removeLast())
        }
        s = buffer.removeLast()
        while(!buffer.isEmpty()){
            myStack.add(buffer.removeLast())
        }
        pos--
        return s
    }

    fun peek(): Int {
        var s : Int = 0
        while(!myStack.isEmpty()){
            buffer.add(myStack.removeLast())
        }
        s = buffer.last()
        while(!buffer.isEmpty()){
            myStack.add(buffer.removeLast())
        }
        return s
    }

    fun empty(): Boolean {
        return pos == -1
    }

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * var obj = MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */

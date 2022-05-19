bool isValid(String s)  {
    Map<String,String> lookupTable = {
      ")":"(",
      "}":"{",
      "]":"["
    };
 
    
    List<String> stack = [];
    
    for(int i=0; i<s.length; i++){
      var l = stack.length;
      
      if(lookupTable.containsKey(s[i])){
        if(l>0 && stack.last == lookupTable[s[i]]){
          stack.removeLast();
        }else{
          return false;
        }
      }else{
        stack.add(s[i]);
      }
      
    }
    
    return stack.isEmpty;
}

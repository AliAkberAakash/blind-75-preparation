class Solution {
public:
    bool isValid(string s) {
        int len = s.size();
        
        if(len%2 !=0) return false;
        
        
        stack<int> buffer;
        for(char c : s){
            if(isLeft(c)){
                buffer.push(c);
            }else{
                if(buffer.empty()) return false;
                
                char top = buffer.top();
                buffer.pop();
                
                if(!isMatch(top, c)) return false;   
            }
        }
        
        if(!buffer.empty()) return false;
        
        return true;
    }
    
    bool isLeft(char c){
        return c=='(' || c=='{' || c=='[';
    }
    
    bool isMatch(char left, char right){
        return (left=='(' && right==')') ||
            (left=='{' && right=='}') ||
            (left=='[' && right==']');
    }
    
};

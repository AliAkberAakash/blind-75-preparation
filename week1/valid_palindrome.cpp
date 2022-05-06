class Solution {
public:
    bool isPalindrome(string s) {
        string s2 = "";
        
        for(char c : s){
            char lowerC = tolower(c);
            if(isLetter(lowerC) || isNumber(lowerC))
                s2 += lowerC;
        }
        
        string rev = s2;
        reverse(rev.begin(),rev.end());
        
        return s2==rev;
    }
    
    bool isLetter(char lowerC){
        return lowerC >= 'a' && lowerC <= 'z';
    }
    
    bool isNumber(char lowerC){
        return lowerC >= '0' && lowerC <= '9';
    }
    
};

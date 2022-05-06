class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int mx = INT_MIN;
        int currIdx = 0;
        
        for(int i=0; i<prices.size(); i++){
            int localMax = max(0,prices[i]-prices[currIdx]);
            if(localMax>mx){
                mx = localMax;
            }
            if(localMax==0)
                    currIdx = i;
        }
        
        return mx;
        
    }
    
    int max(int a, int b){
        return a>=b ? a : b;
    }
    
};

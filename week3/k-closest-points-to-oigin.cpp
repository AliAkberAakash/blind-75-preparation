class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        vector<vector<int>> answer;
        map<double,vector<vector<int>>> lookupTable;
        priority_queue<double, vector<double>, greater<double>> pq;
        
        for(int i=0; i<points.size(); i++){
            double d = distance(points[i]);
            pq.push(d);
            
            lookupTable[d].push_back(points[i]);
        }
        
        for(int i=0; i<k && !pq.empty(); i++){
            double top = pq.top();
            vector<vector<int>> x = lookupTable[top];
            
            if(x.size()>0){
                for (vector<int> v : x){
                    answer.push_back(v);
                    vector<vector<int>> m;
                    lookupTable[top] = m;
                }
            }
            
            pq.pop();
        }
        
        return answer;
    }
    
    double distance(vector<int> point){
        int x = point[0]*point[0];
        int y = point[1]*point[1];
        
        return sqrt(x+y);
    }
    
};

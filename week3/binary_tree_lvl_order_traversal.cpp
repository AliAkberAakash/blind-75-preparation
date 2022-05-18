class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> ans;
        vector<queue<TreeNode*>> q1;
        
        if(root == NULL){
            return ans;
        }
        
        ans.push_back(vector<int>());
        ans[0].push_back(root->val);
        
        q1.push_back(queue<TreeNode*>());
        q1[0].push(root);
        
        for(int i=0; i<q1.size(); i++){
            queue<TreeNode*> q2;
            vector<int> lvl;
            
            while(!q1[i].empty()) {
                
                TreeNode* front = q1[i].front();
                q1[i].pop();
                
                if(front->left != NULL){
                    q2.push(front->left);
                    lvl.push_back(front->left->val);
                }
                if(front->right != NULL){
                    q2.push(front->right);
                    lvl.push_back(front->right->val);
                }
            }
            
            if(!q2.empty()){
                q1.push_back(q2);
            }
            
            if(lvl.size() > 0){
                ans.push_back(lvl);
            }
        }
        
        return ans;
    }
};

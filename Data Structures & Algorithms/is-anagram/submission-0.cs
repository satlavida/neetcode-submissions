public class Solution {
    public bool IsAnagram(string s, string t) {
        if(s.Length != t.Length) return false; 
        Dictionary<char, int> ss = new Dictionary<char, int>();
        Dictionary<char, int> ts = new Dictionary<char, int>(); 

        for(int i = 0; i < s.Length; i++) 
        {
            if(ss.ContainsKey(s[i]))
                ss[s[i]]++;
            else
                ss.Add(s[i], 1);
            if(ts.ContainsKey(t[i]))
                ts[t[i]]++;
            else
                ts.Add(t[i], 1);
        }
        foreach (char x in ss.Keys) 
        {
            if(!ts.ContainsKey(x)) return false; 
            if(ss[x] != ts[x]) return false;
        }
        return true; 
    }
}

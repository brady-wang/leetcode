public class Method01 {
    public static void main(String[] args) {
        int x = 1;
        int y = 2;
        int z;
        Solution solution = new Solution();
        z = solution.hammingDistance(x,y);
        System.out.println(z);
    }

}

class Solution {
    public int hammingDistance(int x, int y) {
        return Integer.bitCount(x ^ y);
    }
}


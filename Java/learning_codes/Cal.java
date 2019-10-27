import java.util.Arrays;

public class Cal {
    public static void main(String[] args) {
        int[] ns = { 1, 4, 9, 16, 25 };
        int sum = 0;
        System.out.println(Arrays.toString(ns));
        for (int n : ns) {
            sum += n;
        }
        System.out.println(sum); // 55
    }
}
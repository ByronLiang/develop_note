public class Test {
    public static void main(String[] args) {
        // String sa = "123";
        // String sk = "123";
        // if (sa.equals(sk)) {
        //     System.out.printf("a=%d, b=%d", a, b);
        // }
        String s1 = "hello";
        String s2 = "HELLO".toLowerCase();
        System.out.println(s1);
        System.out.println(s2);
        if (s1 == s2) {
            System.out.println("s1 == s2");
        } else {
            System.out.println("s1 != s2");
        }
    }
}
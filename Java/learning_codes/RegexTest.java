import java.util.regex.Pattern;

public class RegexTest {
    public static void main(String[] args) {
        boolean isMatch = Pattern.matches(
                ".*already in group.*",
                "can not join this group, reason:user: abc already in group: 10288");
        if (isMatch) {
            System.out.println("匹配");
        } else {
            System.out.println("不匹配");
        }
    }
}
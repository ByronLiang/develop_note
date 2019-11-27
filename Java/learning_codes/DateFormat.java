import java.util.Calendar;
import java.util.Date;
import java.text.SimpleDateFormat;

public class DateFormat {
    public static void main(String[] args) {
        String time = "2019-11-30 19:00:00";
        String[] list = time.split(" ");
        System.out.println("获得的日期"+list[0]);
        String[] date_list = list[0].split("-");
        Calendar calendar=Calendar.getInstance();
        //获得此刻的时间
		System.out.println("Calendar获得的时间 "+calendar.getTime());
        System.out.println("Calendar获得的时间 "+calendar.getTimeInMillis());
        long week = calendar.get(Calendar.DAY_OF_WEEK);
        System.out.println("Calendar获得的时间 星期数据: "+week);
        long today = calendar.getTimeInMillis();
        calendar.set(
            Integer.parseInt(date_list[0]), 
            (Integer.parseInt(date_list[1]) - 1), 
            Integer.parseInt(date_list[2])
        );
        long target = calendar.getTimeInMillis();
        System.out.println("Calendar获得的时间目标: "+target);
        long days=(target-today) / (1000 * 60 * 60 * 24);
        System.out.println("相距 "+days);
    }
}
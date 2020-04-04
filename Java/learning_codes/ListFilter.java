import java.util.List;
import java.util.ArrayList;

public class ListFilter {
    public static void main(String[] args) {
    	List<Long> list = new ArrayList<Long>();
    	list.add((long) 1);
    	list.add((long) 5);
    	System.out.println("origin list: "+ list.toString());
    	if (list.contains((long) 10)) {
    		System.out.println("match");
    	} else {
    		System.out.println("unmatch");
    	}
    	Long target = (long) 12;
    	Boolean status = list.removeIf(item -> item.equals(target));
    	if (status) {
    		System.out.println("remove target; filter list: "+ list.toString());
    	} else {
    		list.add(target);
    		System.out.println("add target; origin list: "+ list.toString());
    	}
    }
}
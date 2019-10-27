import java.util.Arrays;

public class Basic {
    public static void main(String[] args) {
       Cat cat = new Cat();
       String set_name = "BiLK";
       cat.setName(set_name);
       cat.setHobbies("eat fish", "sleep", "playing");
       String name = cat.getName();
       System.out.println(name);
       set_name = "akb123";
       System.out.println(cat.getName());
       System.out.println(cat.getHobbies());
       for (String hobby : cat.getHobbyArray()) {
            System.out.println(hobby);
        }
    }
}

class Cat {
    private String name;
    private String[] hobbies;

    public String exchange(String name) {
        return name + "_007";
    }

    public void setHobbies(String... hobbies) {
        this.hobbies = hobbies;
    }

    public String getHobbies() {
        return Arrays.toString(hobbies);
    }

    public String[] getHobbyArray()
    {
        return this.hobbies;
    }

    public void setName(String name) {
        String temp = this.exchange(name);
        this.name = temp;
    }

    public String getName() {
        return this.name != null ? this.name : "None Name";
    }
}
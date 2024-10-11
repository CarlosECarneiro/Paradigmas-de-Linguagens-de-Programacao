public class Main {
    public static void main(String[] args) {

        Escola escola1 = new Escola("Escolinha");
        Escola escola2 = new Escola("Escolinhazinha");

        Professor professor1 = new Professor("Zé",escola1,escola2);
        System.out.println(professor1);

    }
}
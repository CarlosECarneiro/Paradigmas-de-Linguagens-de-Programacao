public class Main {
    public static void main(String[] args) {
        Configuracao config = Configuracao.getInstance();
        Configuracao config2 = Configuracao.getInstance();

        if(config == config2){
            System.out.println(true);
        }
    }
}
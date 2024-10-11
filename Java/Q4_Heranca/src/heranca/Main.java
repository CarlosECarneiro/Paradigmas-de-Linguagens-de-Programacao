package heranca;

public class Main {
    public static void main(String[] args) {
        Animal bicho = new Cachorro("cachorro","Bicho");
        Animal bicho2 = new Gato("gato","Bichano");
        System.out.println(bicho.emitir_som());
        System.out.println(bicho2.emitir_som());
    }
}
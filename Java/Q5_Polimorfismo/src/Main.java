public class Main {
    public static void main(String[] args) {
        Animal bicho = new Cachorro("cachorro","Bicho");
        Animal bicho2 = new Gato("gato","Bichano");
        Animal.emitir_sons(bicho,bicho2);
    }

}
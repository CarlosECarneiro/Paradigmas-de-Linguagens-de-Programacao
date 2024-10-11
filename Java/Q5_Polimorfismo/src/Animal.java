public class Animal {
    String especie;
    String nome;
    public Animal(String especie, String nome){
        this.especie=especie;
        this.nome=nome;
    }

    public String emitir_som(){
        return "";
    }

    public static void emitir_sons(Animal... animais) {
        for (Animal animai : animais) {
            System.out.println(animai.emitir_som());
        }

    }

}

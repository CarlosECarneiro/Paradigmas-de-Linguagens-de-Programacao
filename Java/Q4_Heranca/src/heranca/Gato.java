package heranca;

public class Gato extends Animal {
    public Gato(String especie, String nome) {
        super(especie, nome);
    }

    public String emitir_som(){
        return "Miau";
    }
}




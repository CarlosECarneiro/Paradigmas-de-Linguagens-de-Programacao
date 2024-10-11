import java.util.ArrayList;

public class Escola {
    String nome;
    ArrayList<Professor> professores = new ArrayList<Professor>();

    public Escola(String nome){
        this.nome = nome;
    }

    @Override
    public String toString() {
        return "Escola{" +
                "nome='" + nome + '\'' +
                '}';
    }
}

import java.util.ArrayList;

public class Professor {
    String nome;
    ArrayList<Escola> escolas = new ArrayList<Escola>();
    public Professor(String nome, Escola... escolas){
        this.nome=nome;
        for (int i = 0; i < escolas.length; i++) {
            this.escolas.add(escolas[i]);
            escolas[i].professores.add(this);
        }
    }

    @Override
    public String toString() {
        return "Professor{" +
                "nome='" + nome + '\'' +
                ", escolas=" + escolas +
                '}';
    }
}

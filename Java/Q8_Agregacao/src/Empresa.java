import java.util.ArrayList;

public class Empresa {
    String nome;
    ArrayList<Empregado> empregados = new ArrayList<Empregado>();

    public Empresa(String nome){
        this.nome=nome;
    }

    public void addEmpregados(Empregado... empregados){
        for (int i = 0; i < empregados.length; i++) {
            this.empregados.add(empregados[i]);
        }

    }
}

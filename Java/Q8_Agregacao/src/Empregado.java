public class Empregado {
    String nome;
    double salario;

    public Empregado(String nome, double salario){
        this.nome=nome;
        this.salario=salario;
    }

    @Override
    public String toString() {
        return "Empregado{" +
                "nome='" + nome + '\'' +
                '}';
    }
}

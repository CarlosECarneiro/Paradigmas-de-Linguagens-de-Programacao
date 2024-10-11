public class FuncionarioHorista extends Funcionario{
    double valorHora;

    public FuncionarioHorista(String nome, double valorHora){
        this.nome = nome;
        this.valorHora = valorHora;
    }

    public double calcularSalario(double horas) {
        return valorHora*horas;
    }

}

public class FuncionarioAssalariado extends Funcionario{
    double salario;

    public FuncionarioAssalariado(String nome, double salario){
        this.nome = nome;
        this.salario = salario;
    }

    public double calcularSalario(double horas) {
        return salario;
    }

}

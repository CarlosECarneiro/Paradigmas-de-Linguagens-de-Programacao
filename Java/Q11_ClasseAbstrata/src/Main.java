public class Main {
    public static void main(String[] args) {

        Funcionario funcionario1 = new FuncionarioAssalariado("Zé",2000);
        Funcionario funcionario2 = new FuncionarioHorista("João",20);
        System.out.println("Funcionario Assalariado: "+funcionario1.calcularSalario(200));
        System.out.println("Funcionario Horista: "+funcionario2.calcularSalario(200));

    }
}
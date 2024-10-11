public class Main {
    public static void main(String[] args) {

        Empresa empresa1 = new Empresa("Zé's empresa");
        Empregado empregado1 = new Empregado("João",1600);
        Empregado empregado2 = new Empregado("Jorge",5000);
        empresa1.addEmpregados(empregado1,empregado2);
        System.out.println(empresa1.empregados);
    }
}
public class Main {
    public static void main(String[] args) {

        ContaBancaria conta = new ContaBancaria("Zé",100);

        System.out.println("Saldo atual: "+conta.getSaldo());

        conta.depositar(200);
        System.out.println("Saldo atual: "+conta.getSaldo());

        conta.sacar(50);
        System.out.println("Saldo atual: "+conta.getSaldo());
    }
}
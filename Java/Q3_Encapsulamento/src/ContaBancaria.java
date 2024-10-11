public class ContaBancaria {
    private double saldo = 0;
    String titular;
    public ContaBancaria(String titular, double saldo){
        this.titular = titular;
        this.saldo = saldo;
    }
    public double getSaldo() {
        return saldo;
    }

    public void setSaldo(double saldo) {
        this.saldo = saldo;
    }

    public void depositar(double dinheiro){
        this.saldo += dinheiro;
    }

    public void sacar(double dinheiro){
        this.saldo-=dinheiro;
    }
}

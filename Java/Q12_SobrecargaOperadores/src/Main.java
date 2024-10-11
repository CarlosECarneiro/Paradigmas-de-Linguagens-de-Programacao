public class Main {
    public static void main(String[] args) {

        Produto produto1 = new Produto("Treco",10);
        Produto produto2 = new Produto("Coiso",35);
        Produto produto3 = new Produto("Negocio",7);

        System.out.println("Soma dos produtos: "+Produto.add(produto1,produto2,produto3));
    }
}
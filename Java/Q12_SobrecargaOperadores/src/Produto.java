public class Produto {
    String nome;
    double preco;

    public Produto(String nome, double preco){
        this.nome = nome;
        this.preco = preco;
    }

    public static double add(Produto... produtos){
        double soma = 0;
        for (int i = 0; i < produtos.length; i++) {
            soma += produtos[i].preco;
        }
        return soma;
    }
}

public class Carro {
    String marca;
    String modelo;
    int ano;
    int velocidade = 0;

    public Carro(String marca, String modelo, int ano){
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    public void acelerar(){
        this.velocidade+=10;
    }

    public void frear(){
        this.velocidade-=10;
    }

    public void exibirVelocidade(){
        System.out.println("Velocidade atual: "+this.velocidade);
    }
}

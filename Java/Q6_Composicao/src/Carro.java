public class Carro {
    String marca;
    String modelo;
    Motor motor;

    public Carro(String marca, String modelo){
        this.marca = marca;
        this.modelo = modelo;
        this.motor = new Motor("Gasolina",150);
    }
}

public class Main {
    public static void main(String[] args) {
        Carro carro1 = new Carro("Fiat","Uno",1998);
        Carro carro2 = new Carro("Volkswagen","Golf",2014);
        Carro carro3 = new Carro("Honda","Civic",2023);
        System.out.println(carro1.marca+" "+carro1.modelo+" "+carro1.ano);
        System.out.println(carro2.marca+" "+carro2.modelo+" "+carro2.ano);
        System.out.println(carro3.marca+" "+carro3.modelo+" "+carro3.ano);
        carro1.exibirVelocidade();
        carro1.acelerar();
        carro1.exibirVelocidade();
        carro1.frear();
        carro1.exibirVelocidade();
    }
}
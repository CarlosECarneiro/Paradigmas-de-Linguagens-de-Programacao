public class Contrato implements Imprimivel{
    String conteudo;

    public Contrato(String conteudo){
        this.conteudo = conteudo;
    }

    public void imprimir(){
        System.out.println("Imprimindo contrato: "+this.conteudo);
    }


}

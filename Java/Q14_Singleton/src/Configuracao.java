public class Configuracao {
    private static Configuracao instancia;

    private Configuracao(){
    }

    public static Configuracao getInstance(){
        if(instancia == null){
            instancia = new Configuracao();
        }
        return instancia;
    }
}

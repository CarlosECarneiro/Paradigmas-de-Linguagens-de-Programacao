public class Matematica {
    public static int fatorial(int n){
        int resultado = 1;
        for (int i = n; i >0; i--) {
            resultado*=i;
        }
        return resultado;
    }
}

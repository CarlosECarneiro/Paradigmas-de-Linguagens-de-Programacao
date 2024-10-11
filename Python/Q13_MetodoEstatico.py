class Matematica:
    @staticmethod
    def fatorial(n):
        resultado = 1
        for i in range(1,n+1):
            resultado*= i
        return resultado
mat = Matematica()
print(mat.fatorial(5))
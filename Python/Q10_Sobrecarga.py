class Calculadora:
    def somar2(self,x,y):
        return x+y
    def somar3(self,x,y,z):
        return x+y+z

calculadora = Calculadora()
print(calculadora.somar2(5,3))
print(calculadora.somar3(5,3,7))

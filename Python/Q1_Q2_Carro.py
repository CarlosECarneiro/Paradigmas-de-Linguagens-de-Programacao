class Carro:
    def __init__(self,marca,modelo,ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade=0
    def acelerar(self):
        self.velocidade+=10
    def frear(self):
        self.velocidade-=10
    def exibirVelocidade(self):
        print("Velocidade atual:",self.velocidade)
carro1 = Carro("Fiat","Uno",1998)
carro2 = Carro("Volkswagen","Golf",2014)
carro3 = Carro("Honda","Civic",2023)

print(carro1.marca,carro1.modelo,carro1.ano)
print(carro2.marca,carro2.modelo,carro2.ano)
print(carro3.marca,carro3.modelo,carro3.ano)

carro1.exibirVelocidade()
carro1.acelerar()
carro1.exibirVelocidade()
carro1.frear()
carro1.exibirVelocidade()
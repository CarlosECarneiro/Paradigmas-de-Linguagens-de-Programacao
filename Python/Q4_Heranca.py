class Animal:
    def __init__(self,especie,nome):
        self.especie = especie
        self.nome = nome

    def emitir_som(self):
        pass

class Cachorro(Animal):
    def emitir_som(self):
        return "Au Au"

class Gato(Animal):
    def emitir_som(self):
        return "Miau"

bicho = Cachorro(Cachorro, "Bicho")
bicho2 = Gato(Gato, "Bichano")
print(bicho2.emitir_som())
print(bicho.emitir_som())
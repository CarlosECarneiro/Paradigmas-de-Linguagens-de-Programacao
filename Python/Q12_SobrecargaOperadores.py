class Produto:
    def __init__(self,nome,preco) -> None:
        self.nome = nome
        self.preco = preco
    def __add__(self,outro):
        return self.preco + outro.preco

produto1 = Produto("Pão",5)
produto2 = Produto("Pedra",50)
print("Soma do preço dos produtos: ",produto1+produto2)
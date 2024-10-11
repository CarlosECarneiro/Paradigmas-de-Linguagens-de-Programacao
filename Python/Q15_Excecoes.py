from typing import Any

class SaldoInsuficienteException(Exception):
    def __init__(self):
        self.mensagem = "Saldo insuficiente"

    def __str__(self):
        return f"Erro: {self.mensagem}"
class ContaBancaria:
    def __init__(self,saldo,titular):
        self.__saldo=saldo
        self.titular=titular
    def depositar(self,dinheiro):
        self.__saldo+=dinheiro
    def sacar(self,dinheiro):
        if dinheiro > self.__saldo:
            raise SaldoInsuficienteException()
        self.__saldo-=dinheiro
    def get_Saldo(self):
        return self.__saldo
conta = ContaBancaria(100,"João")
print(conta.get_Saldo())
conta.depositar(200)
print(conta.get_Saldo())
conta.sacar(301)
print(conta.get_Saldo())
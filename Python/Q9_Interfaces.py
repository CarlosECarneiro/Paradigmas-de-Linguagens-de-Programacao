from abc import ABC, abstractmethod

class Imprimível(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimível):
    def __init__(self, conteudo):
        self.conteudo = conteudo
    def imprimir(self):
        print(f"Imprimindo Relatório: {self.conteudo}")

class Contrato(Imprimível):
    def __init__(self, conteudo):
        self.conteudo = conteudo
    def imprimir(self):
        print(f"Imprimindo Contrato: {self.conteudo}")

relatorio = Relatorio("Relatorio")
contrato = Contrato("Contrato")
relatorio.imprimir()
contrato.imprimir()
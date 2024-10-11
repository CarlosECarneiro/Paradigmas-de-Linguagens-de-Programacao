from abc import ABC, abstractmethod

class Funcionario(ABC):
    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, nome, valorHora):
        self.nome = nome
        self.valorHora = valorHora
    def calcularSalario(self, horas):
        return horas * self.valorHora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, nome, salario):
        self.nome = nome
        self.salario = salario
    def calcularSalario(self):
        return self.salario

funcionario1 = FuncionarioAssalariado("Zé",2000)
funcionario2 = FuncionarioHorista("João",20)

print("Salário Horista: ",funcionario2.calcularSalario(150))
print("Salário Assalariado: ",funcionario1.calcularSalario())
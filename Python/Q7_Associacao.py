class Escola:
    def __init__(self,nome):
        self.nome = nome
        self.professores = []
    def __str__(self):
        return f"{self.nome}{self.professores}"
class Professor:
    def __init__(self,nome,*escolas):
        self.nome = nome
        self.escolas = escolas
        for i in range(len(escolas)):
            escolas[i].professores.append(self)
    
    def __str__(self) -> str:
        return f"{self.nome}({self.escolas})"

escola = Escola("Ideal")
escola2 = Escola("Iouta")
professor = Professor("Zé",escola,escola2)
professor2 = Professor("João",escola,escola2)
print(escola2.professores)
class Configuracao(object):
  _instances = {}
  def __new__(class_, *args, **kwargs):
    if class_ not in class_._instances:
        class_._instances[class_] = super(Configuracao, class_).__new__(class_, *args, **kwargs)
    return class_._instances[class_]

config = Configuracao()
config2 = Configuracao()
print(config == config2)
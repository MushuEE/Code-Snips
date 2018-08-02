docs = {}

def register_docs(cls):
    docs[cls.__name__] = cls.__doc__

class Meta(type):
    def __new__(meta, name, bases, class_dict):
        cls = type.__new__(meta, name, bases, class_dict)
        register_docs(cls)
        return cls

class RegDocs(metaclass=Meta):
    '''
    My Docs are here
    '''
    def __init__(self):
        pass

print(docs)
from lexer import *
import sys

class Parser:
    def __init__self(self,lexer):
        self.lexer = lexer
        self.curToken = None
        self.peekToken = None
        self.nextToken()
        self.nextToken()  # calling it twice to initiate the current and peek
    
    def checkToken(self,kind):
        return kind == self.curToken.kind
    def checkPeek(self,kind):
        return kind==self.peekToken.kind
    # Try to match current token. If not, error. Advances the current token.
    def match(self,kind):
        if not self.checkToken(kind):
            self.abort("Expected " + kind.name + ", got " + self.curToken.kind.name)
        self.nextToken()

    def nextToken(self):
        self.curToken = self.peekToken # initially it will be None
        self.peekToken = self.lexer.getToken() # this is will be the first token , when called second time it will move to the next token and the current token will be the first token
    
    def abort(self,message):
        sys.exit("Error: "+ message)



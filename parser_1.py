from lexer import *
import sys

class Parser:
    def __init__(self,lexer):
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
            self.abort("Expected " + kind.name + ", got "+ self.curToken.kind.name)
        self.nextToken()

    def nextToken(self):
        self.curToken = self.peekToken # initially it will be None
        self.peekToken = self.lexer.getToken() # this is will be the first token , when called second time it will move to the next token and the current token will be the first token
    
    def abort(self,message):
        sys.exit("Error: "+ message)

    #PRODUCTION RULES
    # program : statement* [ 0 or more statements]
    def program(self):
        print("PROGRAM")
        while not self.checkToken(TokenType.EOF):
            self.statement()
    
    def statement(self):
        if self.checkToken(TokenType.PRINT):
            print("STATEMENT-PRINT")
            self.nextToken()
            if self.checkToken(TokenType.STRING):
                self.nextToken()
            else:
                self.expression()
        elif self.checkToken(TokenType.IF):
            print("STATEMENT-IF")
            self.nextToken()
            self.comparison()
            self.match(TokenType.THEN)
            self.nl()
            while not self.checkToken(TokenType.ENDIF):
                self.statement()

            self.match(TokenType.ENDIF)
            
        elif self.checkToken(TokenType.WHILE):
            print("STATEMENT-WHILE")
            self.nextToken()
            self.comparison()
            self.match(TokenType.REPEAT)
            self.nl()
            while not self.checkToken(TokenType.ENDWHILE):
                self.statement()

            self.match(TokenType.ENDWHILE)
        self.nl()

    def nl(self):
        print("NEWLINE")
        self.match(TokenType.NEWLINE)
        while self.checkToken(TokenType.NEWLINE):
            self.nextToken()

    #def expression(self):




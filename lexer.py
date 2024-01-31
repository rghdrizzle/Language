# -*- coding: utf-8 -*-
"""
Created on Sat Jan 27 13:45:24 2024

@author: admin
"""
import enum
import sys

class TokenType(enum.Enum):
    EOF = -1
    NEWLINE = 0
    NUMBER = 1
    IDENT = 2
    STRING = 3
    # Keywords.
    LABEL = 101
    GOTO = 102
    PRINT = 103
    INPUT = 104
    LET = 105
    IF = 106
    THEN = 107
    ENDIF = 108
    WHILE = 109
    REPEAT = 110
    ENDWHILE = 111
    # Operators.
    EQ = 201  
    PLUS = 202
    MINUS = 203
    ASTERISK = 204
    SLASH = 205
    EQEQ = 206
    NOTEQ = 207
    LT = 208
    LTEQ = 209
    GT = 210
    GTEQ = 211
    INC = 212
    DEC = 213
    


class Token:   
    def __init__(self, tokenText, tokenKind):
        self.text = tokenText   
        self.kind = tokenKind   
        
    @staticmethod
    def checkIfKeyword(tokText):
        for kind in TokenType:
            if kind.name == tokText and kind.value>=100 and kind.value<=200: #checks if the tokentext given is the same as the enum's name and checks if the value is in the range
                return kind
        return None 

class Lexer:
    def __init__(self, source):
        self.source = source + "\n"  # Source code to lex as a string. Append a newline to simplify lexing/parsing the last token/statement.
        self.curChar =''
        self.curPos = -1
        pass

    # Process the next character.
    def nextChar(self):
        self.curPos +=1
        if self.curPos >= len(self.source):
            self.curChar = "\0" #eof
        else:
            self.curChar = self.source[self.curPos] 
        pass

    # Return the lookahead character.
    def peek(self): # to look at the next character without updating the curPos counter
        lookaheadCount =  self.curPos + 1    
        if lookaheadCount >= len(self.source):
            return "\0"
        return self.source[lookaheadCount]
        pass

    # Invalid token found, print error message and exit.
    def abort(self, message):
        sys.exit("Lexing error. " + message)
		
    # Skip whitespace except newlines, which we will use to indicate the end of a statement.
    def skipWhitespace(self):
        while self.curChar == ' ' or self.curChar == '\t' or self.curChar == '\r':
            self.nextChar()
		
    # Skip comments in the code.
    def skipComment(self):
        if self.curChar == '#':
            while self.curChar != '\n':
                self.nextChar()

    # Return the next token.
    def getToken(self):
        self.skipWhitespace()
        self.skipComment()
        token = Token(self.curChar,None)
        if self.curChar == '+':
            if self.peek()== '+':
                lastChar = self.curChar
                self.nextChar()
                token = Token(lastChar+self.curChar,TokenType.INC)
            else:
                token = Token(self.curChar, TokenType.PLUS)
        elif self.curChar == '-':
            if self.peek()== '-':
                lastChar = self.curChar
                self.nextChar()
                token = Token(lastChar+self.curChar,TokenType.DEC)
            else:
                token = Token(self.curChar, TokenType.MINUS)
        elif self.curChar == '*':
            token = Token(self.curChar, TokenType.ASTERISK)
        elif self.curChar == '/':
            token = Token(self.curChar, TokenType.SLASH)
        elif self.curChar == '\n':
            token = Token(self.curChar, TokenType.NEWLINE)
        elif self.curChar == '\0':
            token = Token('', TokenType.EOF)
        elif self.curChar == '=':
            if self.peek() == '=':
                lastChar = self.curChar
                self.nextChar()
                token = Token(lastChar + self.curChar, TokenType.EQEQ)
            else:
                token = Token(self.curChar, TokenType.EQ)
        elif self.curChar == '>':
            # Checking if this token is > or >=
            if self.peek() == '=':
                lastChar = self.curChar
                self.nextChar()
                token = Token(lastChar + self.curChar, TokenType.GTEQ)
            else:
                token = Token(self.curChar, TokenType.GT)
        elif self.curChar == '<':
                # Checking if this token is < or <=
                if self.peek() == '=':
                    lastChar = self.curChar
                    self.nextChar()
                    token = Token(lastChar + self.curChar, TokenType.LTEQ)
                else:
                    token = Token(self.curChar, TokenType.LT)
        elif self.curChar == '!':
            if self.peek() == '=':
                lastChar = self.curChar
                self.nextChar()
                token = Token(lastChar + self.curChar, TokenType.NOTEQ)
            else:
                self.abort("Expected !=, got !" + self.peek())
        elif self.curChar == '\"':
            self.nextChar()
            startPos = self.curPos
            while self.curChar != '\"':
                if self.curChar == '\r' or self.curChar == '\n' or self.curChar == '\t' or self.curChar == '\\' or self.curChar == '%':
                    self.abort("Illegal character in string.")
                self.nextChar()
            tokenText = self.source[startPos : self.curPos]
            token = Token(tokenText,TokenType.STRING)
        elif self.curChar.isdigit():
            startPos = self.curPos
            while self.peek().isdigit():
                self.nextChar()
            if self.peek() == '.': # Decimal!
                self.nextChar()
                # Must have at least one digit after decimal.
                if not self.peek().isdigit(): 
                    self.abort("Illegal character in number.")
                while self.peek().isdigit():
                    self.nextChar()
            tokText = self.source[startPos : self.curPos + 1] 
            token = Token(tokText, TokenType.NUMBER)
        elif self.curChar.isalpha():
            startPos = self.curPos
            while self.peek().isalnum():
                self.nextChar()
            # Check if the token is in the list of keywords.
            tokText = self.source[startPos : self.curPos + 1] # Get the substring.
            keyword = Token.checkIfKeyword(tokText)
            if keyword == None: # Identifier
                token = Token(tokText, TokenType.IDENT)
            else:   # Keyword
                token = Token(tokText, keyword)
        else:
            # Unknown token!
            #self.abort("Unknown token: " + self.curChar)
            pass
        self.nextChar()
        return token
    
    
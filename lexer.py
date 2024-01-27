# -*- coding: utf-8 -*-
"""
Created on Sat Jan 27 13:45:24 2024

@author: admin
"""
import enum
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
        pass
		
    # Skip whitespace except newlines, which we will use to indicate the end of a statement.
    def skipWhitespace(self):
        pass
		
    # Skip comments in the code.
    def skipComment(self):
        pass

    # Return the next token.
    def getToken(self):
        if self.curChar == '+':
            pass	
        elif self.curChar == '-':
            pass	
        elif self.curChar == '*':
            pass	
        elif self.curChar == '/':
            pass	# Slash token.
        elif self.curChar == '\n':
            pass	# Newline token.
        elif self.curChar == '\0':
            pass	# EOF token.
        else:
            # Unknown token!
            pass
        self.nextChar()
class Token:
    def __init__(self,tokentext,tokenkind):
        self.text = tokentext
        self.kind = tokenkind
class TokenType(enum.Enum):  # enums in python have a property to hold values so for example LABEL has the value 101 ( enum in c# is different)
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
        
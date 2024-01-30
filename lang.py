# -*- coding: utf-8 -*-
"""
Created on Sat Jan 27 14:08:58 2024

@author: admin
"""

from lexer import *

def main():
    source = "+- \"This is a string\" # This is a comment!\n */ 23 45.45" 
    lexer = Lexer(source)
    print(lexer.source)

    token = lexer.getToken()
    while token.kind != TokenType.EOF:
        print(token.kind)
        token = lexer.getToken()
    
main()

        

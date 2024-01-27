# -*- coding: utf-8 -*-
"""
Created on Sat Jan 27 14:08:58 2024

@author: admin
"""

from lexer import *

def main():
    source = "LOL IM THE BEST"
    lexer = Lexer(source)
    
    while lexer.peek()!="\0":
        print(lexer.curChar)
        lexer.nextChar()
        
main()
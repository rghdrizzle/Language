# -*- coding: utf-8 -*-
"""
Created on Sat Jan 27 14:08:58 2024

@author: admin
"""

from lexer import *
from parser_1 import *
import sys

def main():
    print("Langauge Compiler")
    if len(sys.argv)!= 2:
        sys.exit("Error: Need a source file as argument.")
    with open(sys.argv[1],'r') as input:
        source = input.read()

    lexer = Lexer(source)
    parser = Parser(source)
    parser.program()
    
main()

        

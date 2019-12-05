import random

ls = ["!","@","\\","\"","#","№","$",";","%",":","&","?","*","(",")","-","_","+","=","[","]","{","}","'","/",".","<",">",",","|"]

for i in range(5):
    random.shuffle(ls)
    for item in ls:
        print(item, end='')
    print("\n")

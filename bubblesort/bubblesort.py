#!/usr/bin/python

array = [2,2,1,3,4,5,6,300,24,7,11]

def bubbleSort(list):
    for i in range(0, len(list)):
        for j in range(1, len(list)-i):
            if list[j-1] > list[j]:
                tmp = list[j-1]
                list[j-1] = list[j]
                list[j] = tmp
    return list

def main():
    if __name__ == "__main__":
        print "before bubble sort:", array
        print "after bubble sort:", bubbleSort(array)

main()
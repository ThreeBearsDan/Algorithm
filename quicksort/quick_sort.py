#!/usr/bin/python

array = [2,2,1,3,4,5,6,300,24,7,11]

def quickSort(list):
    if len(list) > 1:
        llist = []
        rlist = []
        key = [list[0]]
        for i in range(1, len(list)):
            if list[i] > key[0]:
	        rlist.append(list[i])
	    elif list[i] < key[0]:
                llist.append(list[i])
	    else:
	        key.append(list[i])
        return quickSort(llist) + key + quickSort(rlist)
    else:
        return list

if __name__ == "__main__":
    print "before quick sort:", array
    print "after quick sort:", quickSort(array)   
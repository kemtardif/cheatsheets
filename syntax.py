# Python Syntax /////////////////////////

# Swap variables x and y

x = 2
y = 3

z = x
x = y
y = z

# IF-Else statements

if something1:
    if something2:
        Do something1
    else:
        Do something 2
elif something2:
    Do something3
else:
    Do something4

# Functions

def function(x, y):
    Do something with x and y
    return result 

a = function(b,c)

# Lists

emptylist = []
a = [3, "Hello", [1, 2]]

a.append(1) # Add 1 to the list
a.pop() # Revove last item in list
a.pop(0) # Remove item with index 0 in list

a[0] = x 
a[0], a[1] = a[1], a[0] # Swap two elements in list

list(range(1,5)) # Create list [1, 2, 3, 4]

# For loops

a = [x, y, z]
total = 0

for e in a:
    total = total + e

total2 = 0

for i in range(1,5):
    total2 += i

a = [1,2,3,4,-5]
total3 = 0
for element in a
    if element <=0
        break       # Stop loop if condition is True
    total3 =+ element

# While loops

total2 = 0
j = 1
while j < 5:
    total2 += j
    j += 1

a = [1,2,3,4,5]
total3 = 0
i = 0
while i <len(a) and a[i] > 0
    total3 += i
    i += 1

while True:
    Do something
    if condition:   #infinite loop until condition is true
        break



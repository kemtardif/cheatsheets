/*CONSOLE AND USER INPUT
//////////////////*/

Console.Write("Hello, my name is " + name);
Console.WriteLine("Hello, my name is " + name); //Add new line
Console.Readline(); //User input
/n //Add new line to text


//VARIABLES AND DATA TYPES//////////////

string x = "Kem";
int y;
y = 32;
char z = 'Z';
float, double, decimal // rational numbers with increasing order of precision
bool x1 = true //or false

//Strings////////////

string x = "What's good, homie";
x.Length //Return total # of characters
x.ToUpper();
x.ToLower();
x.Contains("homie"); //Return true or false
x[i]; //Return character at index i
x.IndexOf("Homie"); //Return index of first character in parentheses
x.Substring(5;3); //Return the characters starting at index 5
int y = Convert.ToInt32(x); //String to integer
double z = Convert.ToDouble(x); //Convert to rational number
//NUMBERS//////////

Math.Abs();
Math.Pow(x,a); //x to power of a
Math.Sqrt();
Math.Max(x, y); 
Math.Min(x, y);
Math.Round();

//ARRAYS///////

int[] array = {1, 2, 3};
int[,] array2D = { {1, 2}, {3, 4}}; //Generalize to N-demensional arrays
string[] array1 = new string[n] //Empty array that allows n elements
string[,] array12D = new sting[2,3] //Empty arrays with two indices of arrays with 3 indices
array[i];
array2D[0,1]
object[] arrayObj = new Object[5]; //Array with different data types that allows 5 elements

//LISTS//////////////

List<object> listObj = new List<object> (); //List of various data types to hold undefined # of elements
listObj.Add(n); //Add n to list
listObj.RemoveAt(n); //Remove element of index n
 
//METHODS//////

static void Main(string[] args) {...} //All code inside will be ran

static void NewMethod(string a, double b, ...) {...} //Method with arguments a, b, ... that return nothing (void)

static int ReturnMethod(int a, int b, ....) { ///Method that return an integer result
    Do something;
    return result;
} 

//IF STATEMENTS

if (condition1 && condition2) {
    Do something
} else if (conditon3 || !condition4) {
    Do something
} else {
    Do something
}

//SWITCH STATEMENTS///////

static string GetDay(int numDay){ //Case statement with defautl value in no case is true
    string days;
    switch(numDay){
        case 0:
            days = "Sunday";
            break
        case 1:
            days = "Monday";
            break
            ...
        default : 
            days = "Invalid day";
            break
    }

    return days;
}

//WHILE LOOPS////////////

while(condition) {...}

do {...} while (condition); //Do something once before checking condition

//FOR LOOPS///////

for(int i = 0; i < array.Length; i++) {...}
for (int num in array) {...}

// EXCEPTIONS///////////

try {
    Code that might result in an exception
}
catch(Exception e){ //Catch all types of exceptions
    Do something in case the exception arise
    Console.Write(e.Message); //Display the exception on console
}
catch (FormatException e) { //Catch on Format exception
    Do something
    Console.Write(e.Message);
}
finally { //Will do this no matter if we have exception or not
    Do something
}

//OOP/////////////

//Classes are created as new class.cs file

class NewClass 
{
    public string attr1;
    public int attr2;
    public double attr3;
    private int attr4; //Only methods of this class can access this value
    public static int attr5; //Attribute shared by all instances
    public int array = new int[3];

    public NewClass(string a, int b)    //Constructor
    {
        attr1 = a;
        attr2 = b;
        attr5++;

    }

    public int Method1() //Method that returns integer result
    {
        Do something
        attr1 = value;
        return result;
    }

    public int Attr4 
    {
        get {return attr4;} //Return value of private attribute
        set {               //Set value of private attribute if condition is true
            if(condition) {
                attr4 = value;
            }
        }
    }

    public static void staticMethod (int value)
    {
        Do something
    }

}

//Instances are created in public void Main

NewClass instance = new NewClass("Hi", 3);
instance.attr3 = 3.45;
NewClass instance1 = new NewClass();
instance.Method1();
instance.Attr4; //Call value of private attribute
instance.Attr4 = value //Set value of attr4 if it satisfies condition
NewClass.attr5; //Access static attribute
NewClass.staticMethod(n); //Access static Method

//INHERITANCE/////////////////////

class SuperClass {
    public string attr1;
    public int attr2;

    static void Method1()
    {
        Do something
    }

    public int Method2()
    {
        Do something
    }

    public virtual void Method3() //Method that can be overwritten by sub-classes
    {
        Do something
    }
}

class SubClass : SuperClass {

    public int attr3;

    public override void Method3()
    {
        Do something else
    }

    public int Method4()
    {
        Do something
    }
}

SuperClass class1 = new SuperClass(string value);
SubClass class2 = new SubClass(string value1, int value 2, int value 3);

class1.Method1();
SuperClass.Method2();
class2.Method1();
class2.Method4();

class1.Method3(); //Do something
class2.Method3(); //Do something else
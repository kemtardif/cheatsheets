# this is a comment

class Student:
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade 

    def get_grade(self):
        return self.grade

class Course:

    def __init__(self, name, max_students):
        self.name = name
        self.max_students = max_students
        self.students = []
        self.is_active = True 

    def add_students(self, student):
        if len(self.students) < self.max_students:
            self.students.append(student)
            return True
        elif len(self.students) == self.max_students:
            print('Wait in line')
            return False
        else: 
            return False

    def get_average_grade(self):
        value = 0
        for student in self.students:
            value += student.get_grade() 

        return value / len(self.students)

s1 = Student("Tim", 20, 95)
s2 = Student("Barbara", 18, 86)
s3 = Student("Carole", 26, 100)

course = Course("Science", 2)

course.add_students(s1)
course.add_students(s2)

print(course.students[0].name)

print(course.get_average_grade())


# ModelBasedSW_Project
___

## Exercise 1 (Apply both translation methodes):
- We translated the functions, working in the solution folder

## Exercise 2 (Posibility):
- It is not possible to simply translate sum with a generic translation, becous interface does not implement the "+" function
- It is not possible to monomorphize a funtion in go to multiple types without changing its name and thereby decresing the maintainabilty of the code. To keep the behavior of the funtion beeing monomorphized, one would need to change the function in question to a dictionary, taking an interface as argument.
- Translating the swap function to the generic argument type "interface{}" does work but does not allow to be used with a primitive variable, because it "does not implement \*interface{}"
  - e.g.: "\*int does not implement \*interface{}"

## Exercise 3 (Lanuage Features):
- As stated before a language that allowes for function overloading based on argument types would simplify the monomorphization prozess.
- To keep the code with generic translation safe, the language needs to support type assertions or the ability to iterate/branch between different implementations beeing handet to the method
- Structural subtyping is one way to keep a generic translation relatively generic, because every struct that fullfils the needs of the interface automaticaly is a subtype.

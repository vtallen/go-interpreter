### 1.1 Lexical Analysis
This interpreter will function by taking the source code, converting it into a list of tokens, then taking those tokens and building a syntax tree

* The first step is creating a lexer, which turns the source code into the tokens

* How literals are dealt with also varies by implementation, some are not converted to tokens till they get parsed
* The meaning of white space varies by interpreter implementation. In this interpreter, white space acts only as a separator between tokens 

* This interpreter will create a fictitious language called monkey


**Example lexer output:** ![[lexer_output.png]]

### 1.2 Defining Our Tokens
* The tokens the interpreter spits out need to be defined first

**Example monkey code:** ![[monkey_syntax_example.png]]

**What tokens are in this code?**
* Numbers: 5, 10
* Variable names: five, ten, add, result
* Words that are not numbers: let, fn
* Special characters: (, ), {, }, etc

**How each will get handled:**
* Numbers will be treated as just numbers and be given a separate type
* Variable names (identifiers) will be treated the same regardless of what they contain. All we care about is that they are variables
* Non-variable words (keywords) will get treated differently depending on what they are as behavior should be different between a *let* or an *fn*
* Symbols will get handled differently depending on the symbol as well

### token/token.go
**Token Type as a string**
* Allows us to differentiate between types of tokens
* Easy to debug as a string can be printed
* Downside is it is less performent than using an int or byte would be 

All possible token types will be represented by constants

**Special tokens:**
* ILLEGAL - represents a token not defined by the parser
* EOF - represents the end of a file and tells the parser that it can stop
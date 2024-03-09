class ASTNode:
    def __init__(self, value, children=None):
        self.value = value
        self.children = children if children else []

def generate_three_address_code(node):
    global temporary_counter
    if node.value == 'assign':
        target = node.children[0]
        source = node.children[1]
        print(f"{target} = {source}")
    elif node.value == 'add':
        op1 = node.children[0]
        op2 = node.children[1]
        result = node.children[2]
        print(f"{result} = {op1} + {op2}")
    elif node.value == 'sub':
        op1 = node.children[0]
        op2 = node.children[1]
        result = node.children[2]
        print(f"{result} = {op1} - {op2}")
    # Handle other language constructs similarly

def parse(code):
    tokens = code.split(';')
    ast = ASTNode('program')
    for token in tokens:
        token = token.strip()
        if token:
            parts = token.split('=')
            if len(parts) >= 2:  # Ensure there is at least one '=' character
                target = parts[0].strip()
                expression = parts[1].strip()
                ast.children.append(ASTNode('assign', [target, expression]))
            else:
                print(f"Ignoring invalid token: {token}")
    return ast


temporary_counter = 0

def generate_temporary():
    global temporary_counter
    temporary_counter += 1
    return f"temp{temporary_counter}"

# Example input code
input_code = """
a=10;
for(i=0;i<10;i++)
{
  if(i==0)
  {
    a=b+c;
  }
  else
  {
    d=a*1;
  }
}
"""

# Parse the input code and generate the AST
ast = parse(input_code)

# Generate three-address code from the AST
generate_three_address_code(ast)

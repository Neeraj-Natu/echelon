package parser

import (
	"testing"
	"github.com/Neeraj-Natu/shifu/ast"
	"github.com/Neeraj-Natu/shifu/lexer"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input 			 string
		expectedVariable string
		expectedValue 	 interface{}
	} {
		{"let x = 5;","x",5},
		{"let foobar = y;","foobar","y"},
	  }

	for _, tt := range tests {
		
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doesnot contain 1 statements. got=%d",len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t,stmt, tt.expectedVariable){
			return
		}
		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}  
}

func TestReturnStatement(t *testing.T) {
	tests := []struct{
		input		  string
		expectedValue interface{}
	}{
		{"return 5;",5},
		{"return foobar;","foobar"},
	}

	for _,tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t,p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doesnot contain 1 statements. got=%d", len(program.Statements))
		}
		
		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
					returnStmt.TokenLiteral())
		}
		if testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue){
			return
		}

	}

}

func TestVariableExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t,p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements doesnot contain a statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	variable, ok := stmt.Expression.(*ast.Variable)
	if !ok {
		t.Fatalf("exp not *ast.Variable. got=%T", stmt.Expression)
	}
	if variable.Value != "foobar" {
		t.Errorf("variable.Value not %s. got=%s","foobar",variable.Value)
	}
	if variable.TokenLiteral() != "foobar" {
		t.Errorf("variable.TokenLiteral() not %s. got=%s","foobar",variable.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := `5`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t,p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements doesnot contain a statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d",5,literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral() not %s. got=%s","5",literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		Value		 interface{}
	}{
		{"!5","!",5},
		{"-15","-",15},
		{"!foobar","!","foobar"},
	}
	for _,tt := range prefixTests {
		l:= lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t,p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doesnot contain a statement. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp,ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",tt.operator,exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right,tt.Value){
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input        string
		leftValue	 int64
		operator	 string
		rightValue 	 int64
	}{
		{"5 + 10;",5,"+",10},
		{"15 - 5",15,"-",5},
		{"15 * 12",15,"*",12},
		{"18 / 6",18,"/",6},
		{"25 > 23",25,">",23},
		{"25 > 23",25,">",23},
		{"25 < 23",25,"<",23},
		{"334 == 3",334,"==",3},
		{"231 != 23",231,"!=",23}, 
	}
	for _,tt := range infixTests{
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doesnot contain %d statement. got=%d\n",1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		
		if !testInfixExpression(t,stmt.Expression, tt.leftValue,tt.operator,tt.rightValue){
			return
		}
	}

}

func TestOperatorPrecedenceParsing(t *testing.T){
	tests := []struct {
		input 	 string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected,actual)
		}
	}
}




// All helper functions
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value() not '%s'. got = %s",name,letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got = %s",name,letStmt.Name.TokenLiteral())
		return false
	}
	return true
}


func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _,msg := range errors {
		t.Errorf("parser error : %q", msg)
	}
	t.FailNow()
}


func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	literal, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if literal.Value != value {
		t.Errorf("literal.Value not %d. got=%d",value,literal.Value)
		return false
	}
	return true
}

func testVariable(t *testing.T, exp ast.Expression, value string) bool {
	variable, ok := exp.(*ast.Variable)
	if !ok {
		t.Errorf("exp not *ast.Variable. got=%T", exp)
		return false
	}
	if variable.Value != value {
		t.Errorf("variable.Value not %s. got=%s",value,variable.Value)
		return false
	}
	if variable.TokenLiteral() != value {
		t.Errorf("variable.TokenLiteral not %s. got=%s",value,variable.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type){
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)	
	case string:
		return testVariable(t,exp,v)	
	}
	t.Errorf("type of exp not handled. got=%T.",exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%t(%s)",exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left,left){
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not %s. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right,right) {
		return false
	}

	return true
}
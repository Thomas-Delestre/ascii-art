!/bin/bash
echo "--reverse example00.txt"
echo "expected output : Usage"
go run . --reverse example00.txt
echo "*******************************************************************"
echo "--output=test.txt Hello World! varsity"
echo "expected output : a file test.txt create in /fonctions/options/output/"
go run . --output=test.txt "Hello World!" varsity
echo "*******************************************************************"
echo "--align=right "Hello World!" varsity"
echo "expected output : Hello World! on right to the terminal"
go run . --align=right "Hello World!" varsity
echo "*******************************************************************"
echo "--color=blue "Hello World!" varsity"
echo "expected output : Hello World! in blue"
go run . --color=blue "Hello World!" varsity
echo "*******************************************************************"
echo "--reverse example00.txt"
echo "expected output : Hello World"
go run . --reverse=example00.txt
echo "*******************************************************************"
echo "--reverse=testShadow.txt"
echo "expected output : 123 ABC abc?"
go run . --reverse=testShadow.txt
echo "*******************************************************************"
echo "--reverse=testVarsity.txt"
echo "expected output: 123 ABC abc?"
go run . --reverse=testVarsity.txt
echo "*******************************************************************"
echo "--reverse=testThinkertoy.txt"
echo "expected output : 123 ABC abc?"
go run . --reverse=testThinkertoy.txt
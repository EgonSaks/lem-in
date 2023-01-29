echo --------------------------------------
echo
echo example00 6 turns
echo
go run main.go example00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example01 8 turns
echo
go run main.go example01.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example02 11 turn
echo
go run main.go example02.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example03 6 turns
echo
go run main.go example03.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example04 6 turns
echo
go run main.go example04.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example05 8 turns
echo
go run main.go example05.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo badexample00
echo
go run main.go badexample00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo badexample01
echo
go run main.go badexample01.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example06 less than 1.5 min
echo
go run main.go example06.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo example07 less than 2.5 min
echo
go run main.go example07.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo exampleMillionAnts less than 2.5 min
echo
go run main.go exampleMillionAnts.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo test_negativeants 
echo
go run main.go test_negativeants.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo only_ants
echo
go run main.go only_ants.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo test_toomanyants
echo
go run main.go test_toomanyants.txt
echo
echo --------------------------------------
echo empty file
echo
go run main.go example0.txt
echo
echo --------------------------------------
echo
echo "Thank you :)"
exit 0




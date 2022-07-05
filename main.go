package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var users map[int]User
var id int

func crateUser() {
	clearConsole()
	fmt.Print("Ingresa Username: ")
	username := readLine()
	fmt.Print("Ingresa Email: ")
	email := readLine()
	fmt.Print("Ingresa Edad: ")
	age := checkParseInt(readLine())
	id++
	user := User{id, username, email, age}
	users[id] = user
	fmt.Print("\n>>>Usuario Creado Exitosamente\n")
}

func readUser() {
	clearConsole()
	if _, ok := users[1]; !ok {
		fmt.Println("No hay usuarios creados")
	} else {
		fmt.Print("Listado de Usuarios\n\n")
		for id, user := range users {
			fmt.Println("ID:", id, "/ Username:", user.username, "/ Email:", user.email, "/ Age:", user.age)
		}
	}
}

func updateUser() {
	clearConsole()
	fmt.Print("Ingresa el id del Usuario a Eliminar: ")
	id := checkParseInt(readLine())
	if user, ok := users[id]; ok {
		clearConsole()
		fmt.Println("Seleccione el campo que desea Actualizar del usuario:", user.username)
		fmt.Println("1) Username")
		fmt.Println("2) Email")
		fmt.Println("3) Age")
		fmt.Println("4) Todos los campos")
		fmt.Println("5) Cancelar")
		opcion := readLine()

		if opcion == "5" {
			clearConsole()
			fmt.Print("Actualizacion Cancelada\n\n")
			main()
		}
		switch opcion {
		case "1":
			clearConsole()
			fmt.Print("Ingrese el nuevo Username: ")
			user.username = readLine()
			users[id] = user
		case "2":
			clearConsole()
			fmt.Print("Ingrese el nuevo Email: ")
			user.email = readLine()
			users[id] = user
		case "3":
			clearConsole()
			fmt.Print("Ingrese la nueva Age: ")
			age := checkParseInt(readLine())
			user.age = age
			users[id] = user
		case "4":
			clearConsole()
			fmt.Print("Ingrese el nuevo Username: ")
			user.username = readLine()
			fmt.Print("Ingrese el nuevo Email: ")
			user.email = readLine()
			fmt.Print("Ingrese la nueva Age: ")
			age := checkParseInt(readLine())
			user.age = age
			users[id] = user
		}
		fmt.Print(">>>Usuario Actualizado Extitosamente\n\n")
	} else {
		clearConsole()
		fmt.Print("El usuario con ID: '", id, "' no existe\n\n")
	}

}

func deleteUser() {
	clearConsole()
	fmt.Print("Ingresa el id del Usuario a Eliminar: ")

	id := checkParseInt(readLine())

	if _, ok := users[id]; ok {
		delete(users, id)
		fmt.Print(">>>Usuario Eliminado Extitosamente\n\n")
	} else {
		fmt.Print("El usuario con id:", id, "no existe\n\n")
	}

}

func readLine() string {

	if opcion, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(opcion, "\n")
	}

}

func checkParseInt(readline string) int {
	if data, err := strconv.Atoi(readline); err != nil {
		panic("No se puede convertir un String a un Entero")
	} else {
		return data
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)
	for {

		fmt.Print("\nCRUD CLI\n\n")
		fmt.Println("1) Crear")
		fmt.Println("2) Listar")
		fmt.Println("3) Actualizar")
		fmt.Println("4) Eliminar")
		fmt.Print("5) Salir del Sistema\n\n")

		fmt.Print("Seleccione una Opcion ")
		opcion := readLine()

		if opcion == "5" || opcion == "exit" {
			break
		}
		switch opcion {
		case "1", "Crear", "crear":
			crateUser()
		case "2", "Listar", "listar":
			readUser()
		case "3", "Actualizar", "actualizar":
			updateUser()
		case "4", "Eliminar", "eliminar":
			deleteUser()
		default:
			clearConsole()
			fmt.Println("Funcion no valida")

		}
	}
	clearConsole()
	fmt.Print("Saliendo del Sitema ")
	for i := 3; i > 0; i-- {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 800)

	}
	fmt.Println("\nSalida Exitosa, Vuelva Pronto!")
}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
)

func MsgBoxFile() (filename string, err error) {
	filename, err = dialog.File().Filter("Archivo", "All Files", "*").Title("Escoge tu archivo").Load()
	if err != nil {
		return
	}
	return
}

func main() {
	fmt.Print("Selecciona tu archivo a spoofear la extension (Presiona enter para continuar)")
	fmt.Scanln()
	file, err := MsgBoxFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Consigue la extension el archivo
	Ext := filepath.Ext(file)

	// pregunta la extension
	fmt.Print("\nEscribe la extension que quieres (ej .png - .pdf - .mp3)> ")
	var extension string
	_, err = fmt.Scanln(&extension)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if string(extension[0]) == "." {
		extension = extension[1:]
	}
	// Reversea el string
	Ext_Spoof := Reverse(extension)

	// Consegui solo el nombre del archivo (sin ruta)
	Name := filepath.Base(file)

	// Quita la extension del nombre
	Name = Name[:len(Name)-4]

	// Da formato para usar RTLO
	Change := fmt.Sprintf("%s\u202e\ufeff%s%s", Name, Ext_Spoof, Ext)

	//Cambia el nombre del archivo aplicando el cambio con RTLO
	os.Rename(file, Change)
}

func Reverse(s string) (f string) {
	for _, i := range s {
		f = string(i) + f
	}
	return
}

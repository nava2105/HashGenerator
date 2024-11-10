package main

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

var uploadTemplate = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <title>Subir Archivo PDF</title>
</head>
<body>
    <div class="container mt-5">
        <h1>Subir Archivo PDF</h1>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <div class="form-group">
                <label for="file">Selecciona un archivo PDF:</label>
                <input type="file" name="file" id="file" class="form-control" accept=".pdf" required>
            </div>
            <button type="submit" class="btn btn-primary">Subir</button>
        </form>
        {{if .}}
        <div class="mt-3 alert alert-info">
            <strong>Hash SHA256 del archivo:</strong> {{.}}
        </div>
        {{end}}
    </div>
    <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.3/dist/umd/popper.min.js"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
</body>
</html>
`

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Limitar el tamaño del archivo a subir
		r.ParseMultipartForm(10 << 20) // 10 MB

		// Obtener el archivo del formulario
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Verificar que el archivo sea un PDF
		if filepath.Ext(header.Filename) != ".pdf" {
			http.Error(w, "Solo se permiten archivos PDF", http.StatusBadRequest)
			return
		}

		// Crear un hash SHA256
		hash := sha256.New()

		// Copiar el contenido del archivo al hash
		if _, err := io.Copy(hash, file); err != nil {
			http.Error(w, "Error al procesar el archivo", http.StatusInternalServerError)
			return
		}

		// Obtener el hash en formato hexadecimal
		hashSum := hash.Sum(nil)
		hashHex := fmt.Sprintf("%x", hashSum)

		// Renderizar la plantilla con el hash
		tmpl := template.Must(template.New("upload").Parse(uploadTemplate))
		tmpl.Execute(w, hashHex)
		return
	}

	// Si no es un POST, simplemente muestra el formulario
	tmpl := template.Must(template.New("upload").Parse(uploadTemplate))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/upload", uploadHandler) // Ruta para subir archivos
	fmt.Println("Servidor escuchando en http://localhost:8080/upload")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func nuevoproducto(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("productos.html"))
	tmpl.Execute(w, nil)
}

func procesar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
		productoForm := r.Form.Get("producto")
		cantidadForm := r.Form.Get("cantidad")
		precioForm := r.Form.Get("precio")
		produ := Producto{ID: id_next, Producto: productoForm, Cantidad: cantidadForm, Precio: precioForm}
		producto = append(producto, produ)
		fmt.Println(producto)
		fmt.Println(produ)
		id_next = id_next + 1
		http.Redirect(w, r, "/mostrar", http.StatusFound)
	} else {
		errorHandler(w, r, http.StatusMethodNotAllowed)
	}
}

func mostrar(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/mostrar" {
		tmpl := template.Must(template.ParseFiles("ver.html"))
		tmpl.Execute(w, producto)
	} else {
		errorHandler(w, r, http.StatusNotFound)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Chamo no se encontró un coño")
	}
}
func eliminar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
		idForm := r.Form.Get("id")
		id, err := strconv.Atoi(idForm)
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
		for i, p := range producto {
			if p.ID == id {
				producto = append(producto[:i], producto[i+1:]...)
				break
			}
		}
		http.Redirect(w, r, "/mostrar", http.StatusFound)
	} else {
		errorHandler(w, r, http.StatusMethodNotAllowed)
	}
}
func actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
		idForm := r.Form.Get("id")
		id, err := strconv.Atoi(idForm)
		if err != nil {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
		productoForm := r.Form.Get("producto")
		cantidadForm := r.Form.Get("cantidad")
		precioForm := r.Form.Get("precio")

		for i, p := range producto {
			if p.ID == id {
				producto[i].Producto = productoForm
				producto[i].Cantidad = cantidadForm
				producto[i].Precio = precioForm
				break
			}
		}
		http.Redirect(w, r, "/mostrar", http.StatusFound)
	} else {
		errorHandler(w, r, http.StatusMethodNotAllowed)
	}
}
func main() {
	//127.0.0.1 localhost ::1
	http.HandleFunc("/nuevoproducto", nuevoproducto)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/mostrar", mostrar)
	http.HandleFunc("/crear", procesar)
	http.HandleFunc("/eliminar", eliminar)
	http.HandleFunc("/actualizar", actualizar)

	fmt.Println("Servidor iniciando en el puerto 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}

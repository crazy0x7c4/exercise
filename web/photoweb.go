package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoArray, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArray {
		templateName = fileInfo.Name()
		if stuffix := path.Ext(templateName); stuffix != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading templates:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func main() {
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/form", safeHandler(formHandler))
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates["upload.html"].Execute(w, nil)
		check(err)
		// io.WriteString(w,
		// 	"<html><body><form method=\"POST\" action=\"/upload\" "+
		// 		"enctype=\"multipart/form-data\">"+
		// 		"Choose an image to upload:"+
		// 		"<input type=\"file\" name=\"image\" />"+
		// 		"<input type=\"submit\" value=\"upload\"/>"+
		// 		"</form></body></html>")
	} else if r.Method == "POST" {
		file, fileHeader, err := r.FormFile("image")
		defer file.Close()
		check(err)
		fileName := fileHeader.Filename

		saveFile, err := os.Create(UPLOAD_DIR + "/" + fileName)
		defer saveFile.Close()
		check(err)

		_, errCopy := io.Copy(saveFile, file)
		check(errCopy)

		http.Redirect(w, r, "/view?id="+fileName, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArray, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	images := []string{}
	for _, fileInfo := range fileInfoArray {
		images = append(images, fileInfo.Name())
	}
	locals := make(map[string]interface{})
	locals["images"] = images

	templates["list.html"].Execute(w, locals)

	// var listHtml string
	// for _, fileInfo := range fileInfoArray {
	// 	imageId := fileInfo.Name()
	// 	listHtml += "<html><body><li><a href=\"/view?id=" +
	// 		imageId + "\">" + imageId +
	// 		"</a></li></body></html>"

	// }
	// listHtml += "<br><a href=\"/upload\">upload an image</a>"
	// io.WriteString(w, listHtml)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template := templates["form.html"]
		template.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		for k, v := range r.Form {
			log.Println(k, v)
		}
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); err != nil && ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v. - %v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

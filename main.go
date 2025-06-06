package main

import (
    "embed"
    "html/template"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

//go:embed templates/index.html
var templates embed.FS

//go:embed templates/favicon.ico
var favicon []byte

type PageData struct {
    ServerName string
    IP         string
    Time       string
}

func getClientIP(r *http.Request) string {
    xff := r.Header.Get("X-Forwarded-For")
    if xff != "" {
        ips := strings.Split(xff, ",")
        return strings.TrimSpace(ips[0])
    }
    xri := r.Header.Get("X-Real-IP")
    if xri != "" {
        return xri
    }
    return r.RemoteAddr
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("Iniciando servidor na porta 8080...")

    tmpl, err := template.ParseFS(templates, "templates/index.html")
    if err != nil {
        log.Fatal("Erro ao carregar template embedado: ", err)
    }

    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/x-icon")
        w.WriteHeader(http.StatusOK)
        _, err := w.Write(favicon)
        if err != nil {
            log.Println("Erro ao enviar favicon:", err)
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Requisição recebida de %s", getClientIP(r))

        serverName, err := os.Hostname()
        if err != nil {
            serverName = "Desconhecido"
        }

        data := PageData{
            ServerName: serverName,
            IP:         getClientIP(r),
            Time:       time.Now().Format("2006-01-02 15:04:05"),
        }

        err = tmpl.Execute(w, data)
        if err != nil {
            log.Printf("Erro ao executar template: %v", err)
            http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
            return
        }

        log.Println("Resposta enviada com sucesso")
    })

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Erro ao iniciar o servidor: ", err)
    }
}

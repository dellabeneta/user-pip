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

// PageData define os dados a serem exibidos no template
type PageData struct {
    ServerName string
    IP         string
    Time       string
}

// getClientIP tenta obter o IP real do cliente, considerando proxies
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

    // Carrega o template uma vez
    tmpl, err := template.ParseFS(templates, "templates/index.html")
    if err != nil {
        log.Fatal("Erro ao carregar template embedado: ", err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Requisição recebida de %s", getClientIP(r))

        serverName, err := os.Hostname()
        if err != nil {
            serverName = "Desconhecido"
        }

        ip := getClientIP(r)
        currentTime := time.Now().Format("2006-01-02 15:04:05")

        data := PageData{
            ServerName: serverName,
            IP:         ip,
            Time:       currentTime,
        }

        // Usa o template carregado com embed
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

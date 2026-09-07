/*
 * example.go - arquivo de exemplo com vulnerabilidades intencionais
 * utilizado para validar as regras customizadas do semgrep para Go
 * NÃO UTILIZAR EM PRODUÇÃO - contém código inseguro proposital
 */

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

// ==========================================================
// Vulnerabilidade 1: Command Injection (CWE-78)
// Regra Semgrep: go-command-injection
// exec.Command com input do usuário não sanitizado
// ==========================================================
func commandInjectionExample(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("cmd")

	// exec.Command com input externo - risco de injeção de comandos
	cmd := exec.Command("sh", "-c", userInput)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "erro ao executar comando", 500)
		return
	}

	fmt.Fprintf(w, "Resultado: %s", output)
}

// ==========================================================
// Vulnerabilidade 2: Erro Ignorado (CWE-754)
// Regra Semgrep: go-ignored-error
// _ = err ou erro não tratado em operações críticas
// ==========================================================
func ignoredErrorExample() {
	// erro ignorado com _ - pode ocultar falhas críticas
	result, _ := http.Get("https://api.exemplo.com/dados-sensiveis")
	if result != nil {
		defer result.Body.Close()
	}

	// outro exemplo de erro ignorado
	data, _ := exec.Command("cat", "/etc/passwd").Output()
	fmt.Println(string(data))
}

// ==========================================================
// Vulnerabilidade 3: Hash Fraco (CWE-327)
// Regra Semgrep: go-weak-hash
// crypto/md5 e crypto/sha1 para dados sensíveis
// ==========================================================
func weakHashExample(password string) {
	// md5 é considerado criptograficamente quebrado
	hash := md5.Sum([]byte(password))
	fmt.Printf("MD5 da senha: %x\n", hash)

	// sha1 também não é recomendado para dados sensíveis
	sha1Hash := sha1.Sum([]byte(password))
	fmt.Printf("SHA1 da senha: %x\n", sha1Hash)

	// usando md5.New() diretamente
	hasher := md5.New()
	hasher.Write([]byte(password))
	fmt.Printf("MD5 (New): %x\n", hasher.Sum(nil))
}

// ==========================================================
// Vulnerabilidade 4: SQL Injection (CWE-89)
// Regra Semgrep: go-sql-string-concat
// fmt.Sprintf dentro de db.Query - concatenação de strings em SQL
// ==========================================================
func sqlInjectionExample(db *sql.DB, userInput string) {
	// sql injection via fmt.Sprintf em Query
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", userInput)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("erro na query:", err)
		return
	}
	defer rows.Close()

	// sql injection via fmt.Sprintf diretamente no Exec
	db.Exec(fmt.Sprintf("DELETE FROM sessions WHERE user_id = '%s'", userInput))

	// sql injection via QueryRow
	db.QueryRow(fmt.Sprintf("SELECT password FROM users WHERE email = '%s'", userInput))
}

func main() {
	fmt.Println("=== Exemplo de código Go vulnerável para validação Semgrep ===")

	// demonstrar hash fraco
	weakHashExample("senha_super_secreta")

	// demonstrar erro ignorado
	ignoredErrorExample()

	// iniciar servidor com rota vulnerável a command injection
	http.HandleFunc("/exec", commandInjectionExample)
	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}

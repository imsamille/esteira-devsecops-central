/*
 * example.c - arquivo de exemplo com vulnerabilidades intencionais
 * utilizado para validar as regras customizadas do semgrep para C/C++
 * NÃO UTILIZAR EM PRODUÇÃO - contém código inseguro proposital
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* ==========================================================
 * Vulnerabilidade 1: Buffer Overflow (CWE-120/CWE-121)
 * Regra Semgrep: cpp-unsafe-string-functions
 * Funções inseguras: strcpy, strcat, sprintf, gets
 * ========================================================== */
void buffer_overflow_example() {
    char small_buffer[10];
    char *user_input = "esta string é muito maior que o buffer de destino";

    // strcpy não verifica o tamanho do destino
    strcpy(small_buffer, user_input);

    // strcat concatena sem verificar overflow
    char dest[20] = "hello";
    strcat(dest, user_input);

    // sprintf não limita o tamanho da saída
    char formatted[10];
    sprintf(formatted, "user: %s", user_input);

    // gets é inerentemente inseguro (removido do C11)
    char line[50];
    gets(line);
}

/* ==========================================================
 * Vulnerabilidade 2: Null Pointer Dereference (CWE-476)
 * Regra Semgrep: cpp-malloc-without-null-check
 * malloc sem verificação de retorno NULL
 * ========================================================== */
void malloc_without_check_example() {
    // alocação sem verificar se retornou NULL
    int *data = malloc(sizeof(int) * 1000000);
    *data = 42;  // se malloc falhou, isso causa segfault

    free(data);
}

/* ==========================================================
 * Vulnerabilidade 3: Command Injection (CWE-78)
 * Regra Semgrep: cpp-command-injection
 * system(), popen(), exec*() com input não sanitizado
 * ========================================================== */
void command_injection_example(const char *user_input) {
    char command[256];
    sprintf(command, "ls -la %s", user_input);

    // system() executa comandos do shell - risco de injeção
    system(command);

    // popen() também permite injeção de comandos
    FILE *fp = popen(command, "r");
    if (fp) {
        pclose(fp);
    }
}

/* ==========================================================
 * Vulnerabilidade 4: Integer Overflow (CWE-190)
 * Regra Semgrep: cpp-integer-overflow-risk
 * Casts entre tipos inteiros de tamanhos diferentes
 * ========================================================== */
void integer_overflow_example() {
    long big_number = 2147483648L;  // maior que INT_MAX

    // cast de long para int pode causar truncamento/overflow
    int truncated = (int)big_number;

    // cast para short é ainda mais perigoso
    short small = (short)big_number;

    printf("Original: %ld, Truncado int: %d, Truncado short: %d\n",
           big_number, truncated, small);
}

int main(int argc, char *argv[]) {
    printf("=== Exemplo de código C vulnerável para validação Semgrep ===\n");

    buffer_overflow_example();
    malloc_without_check_example();

    if (argc > 1) {
        command_injection_example(argv[1]);
    }

    integer_overflow_example();

    return 0;
}

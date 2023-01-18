# Gotutui
<p align = "center">
    <img src="https://media.giphy.com/media/Csn0eDUKV8I1L5UGgJ/giphy.gif"
         alt="gif tui grep script"/>
</p>

### TUI Elements

Built with charmbracelet/lipgloss,bubbles.

Usando los widgets que tiene para presentar un programa más interactivo por la terminal.

La carpeta seqpi contiene programas de c++ que hice para representar la secuencia de números de la cifra trascendental pi. 
Ahora tiene color, y para escoger ver $e\; ,\;\pi\; ,\;\phi$

Tienes que tener go instalado, si tu cpu no es amd64 entonces con g++ compila *dcypher_nums.cpp*, este programa es llamado desde go y espera que se llame epiphi.

> compilar en /seqpi/:
```bash
    g++ -std=c++20 dcypher_nums.cpp -o epiphi
```

> Scripts use `grep` command so it won't work on windows without WSL.
I have yet to translate it to powershell if replicable



### Uso:

```bash
    go run TUIe1.go
```

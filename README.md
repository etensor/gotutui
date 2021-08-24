# Gotutui
## TUI Elements

Built with charmbracelet/lipgloss,bubbles.

Usando los widgets que tiene para presentar un programa más interactivo por la terminal.

La carpeta seqpi contiene programas de c++ que hice para representar la secuencia de números de la cifra trascendental pi. 
Ahora tiene color, y para escoger ver $e \or \pi \or \phi$

Tienes que tener go instalado, si tu cpu no es amd64 entonces con g++ compila *dcypher_nums.cpp*, este programa es llamado desde go y espera que se llame epiphi.

#### if arch != amd64 and execs -> 1:
> compilar en /seqpi/:
```bash
    g++ dcypher_nums.cpp -o epiphi
```
Hay un script en zsh que no se si falle, entonces arreglar el principio del archivo .sh con bash o sh...
 #!/usr/bin/zsh  $ \; \rightarrow \;$ #!/usr/bin/bash

Ese programa tambien representa las cifras usando **grep**, pero no las cuenta, solo las separa de acuerdo al modo escogido. 

### Uso:

```bash
    go run TUIe1.go
```

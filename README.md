# File explorer CLI

Utilidad de línea de comandos para navegar carpetas a través de una interfaz interactiva.

![](README_media/file_explorer_cli.gif "demo")

Librerías utilizadas: 

- https://github.com/eiannone/keyboard Para capturar eventos de Keypress desde la terminal.

Probado en WSL2 (Ubuntu 24.04.1)

## Instalación (Linux o WSL2)

Clonar repositorio y entrar al directorio

```bash
git clone https://github.com/byga12/file_explorer_cli ; cd file_explorer_cli
```

Instalar la aplicación

```bash
go install
```

Verificar la carpeta de instalación

```bash
go list -f '{{.Target}}'
```

Por ejemplo, si el comando arroja /home/Juan123/go/bin/file_explorer_cli la carpeta de instalación es /home/Juan123/go/bin . Este último path lo utilizaremos en el siguiente paso.

Configurar .bashrc

Normalmente este archivo se encuentra en la carpeta del usuario:

```bash
nano /home/Juan123/.bashrc
```

Al final del archivo, añadir estas líneas (para actualizar el PATH y asignar un alias)

```bash
export PATH=$PATH:/home/Juan123/go/bin
alias fe="file_explorer_cli; cd $(cat /tmp/file_explorer_cli_path.txt)"
```

Reiniciar la terminal y probar el alias

```bash
fe
```

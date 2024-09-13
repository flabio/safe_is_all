#!/bin/bash

# Lista de subdirectorios donde se encuentran los archivos Go
directories=("gatewey" "safe_msvc_rol" "safe_msvc_user" "safe_msvc_school")

# Recorre cada subdirectorio y ejecuta el archivo Go
for dir in "${directories[@]}"; do
  if [ -d "$dir" ]; then
    echo "Navegando al subdirectorio $dir"
    cd "$dir" || exit
    
    if [ -f "main.go" ]; then
      echo "Ejecutando go run main.go en $dir"
      go run main.go
    else
      echo "main.go no existe en el subdirectorio $dir"
    fi

    # Vuelve al directorio principal
    cd - > /dev/null
  else
    echo "El subdirectorio $dir no existe."
  fi
done
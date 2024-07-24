#!/bin/bash

# Определение переменных
PROJECT_NAME="ArchUpdateTimer"
PROJECT_DIR="$(pwd)"
BUILD_DIR="$PROJECT_DIR/build"
AUTOSTART_DIR="$HOME/.config/autostart"
DESKTOP_FILE="$AUTOSTART_DIR/$PROJECT_NAME.desktop"
EXECUTABLE_FILE="$BUILD_DIR/$PROJECT_NAME"

# Функция для удаления файла автозапуска
remove_autostart_file() {
    if [ -f "$DESKTOP_FILE" ]; then
        echo "Удаление файла автозапуска..."
        rm "$DESKTOP_FILE"
        echo "Файл автозапуска удален."
    else
        echo "Файл автозапуска не найден."
    fi
}

# Функция для удаления собранного исполняемого файла и конфига
remove_executable_folder() {
    if [ -d "$BUILD_DIR" ]; then
        echo "Удаление папки сборки..."
        # удаление папки сборки
        rm -rf "$BUILD_DIR"
        echo "Папка сборки удалена."
    else
        echo "Папка сборки не найдена."
    fi
}

# Основная логика скрипта
main() {
    remove_autostart_file
    remove_executable_folder

    echo "Удаление завершено."
}

# Запуск основной функции
main

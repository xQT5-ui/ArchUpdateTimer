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

# Функция для удаления собранного исполняемого файла
remove_executable_file() {
    if [ -f "$EXECUTABLE_FILE" ]; then
        echo "Удаление собранного исполняемого файла..."
        rm "$EXECUTABLE_FILE"
        echo "Собранный исполняемый файл удален."
    else
        echo "Собранный исполняемый файл не найден."
    fi
}

# Основная логика скрипта
main() {
    remove_autostart_file
    remove_executable_file

    echo "Удаление завершено."
}

# Запуск основной функции
main

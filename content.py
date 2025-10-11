### Создаём оглавление для Markdown


import re

def create_anchor(header_line: str) -> str:
    """
    Создает якорь в стиле GitHub/Markdown, правильно обрабатывая спецсимволы.
    """
    title = header_line.lstrip('# ').strip()
    anchor = title.lower()

    # Заменяем неразрывные пробелы (U+A0) на обычные
    anchor = anchor.replace('\xa0', ' ')

    # Удаляем все, что не является буквой, цифрой, пробелом или дефисом
    anchor = re.sub(r'[^a-zа-я0-9\s-]', '', anchor)

    # Заменяем пробелы на дефисы
    anchor = re.sub(r'\s+', '-', anchor)

    # Убираем дефисы в начале или в конце
    anchor = anchor.strip('-')

    return f"#{anchor}"

def generate_toc(filepath: str):
    """
    Генерирует оглавление для Markdown файла.
    """
    toc = []
    first_h1_skipped = False

    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            for line in f:
                if not line.startswith('#'):
                    continue

                header_line = line.strip()

                if "оглавление" in header_line.lower():
                    continue

                if header_line.startswith('# '):
                    if not first_h1_skipped:
                        first_h1_skipped = True
                        continue

                # Определяем уровень заголовка (H1=1, H2=2 и т.д.)
                level = len(header_line) - len(header_line.lstrip('#'))

                # Рассчитываем отступ. H1 - нет отступа, H2 - один отступ и т.д.
                # Формула (level - 1) восстанавливает правильную вложенность.
                indent_level = max(0, level - 1)
                indent = '  ' * indent_level

                title = header_line.lstrip('# ').strip()
                anchor = create_anchor(header_line)
                toc_line = f"{indent}- [{title}]({anchor})"
                toc.append(toc_line)

    except FileNotFoundError:
        return f"Ошибка: Файл '{filepath}' не найден."

    return "\n".join(toc)

# --- Основная часть ---
if __name__ == "__main__":
    readme_file = "README.md"
    output_file = "TOC.md" # Имя файла для сохранения результата

    # Генерируем оглавление
    table_of_contents = generate_toc(readme_file)

    # Сохраняем оглавление в файл
    try:
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(table_of_contents)
        print(f"✅ Оглавление успешно сохранено в файл: {output_file}")
    except IOError as e:
        print(f"❌ Не удалось записать в файл {output_file}: {e}")

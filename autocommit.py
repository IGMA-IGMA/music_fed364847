import git
import os

def find_repo_root():
    """Находит корень Git-репозитория, начиная с текущей папки."""
    try:
        # Ищем .git в текущей папке или выше
        repo = git.Repo(os.getcwd(), search_parent_directories=True)
        return repo
    except git.exc.InvalidGitRepositoryError:
        return None

def generate_commit_message(repo):
    """Анализирует изменения и создает краткое описание."""
    # Получаем список измененных файлов
    changed_files = [item.a_path for item in repo.index.diff(None)]
    # Получаем список новых (untracked) файлов
    untracked_files = repo.untracked_files
    
    summary = []
    if changed_files:
        summary.append(f"Updated: {', '.join(changed_files[:2])}")
    if untracked_files:
        summary.append(f"Added: {', '.join(untracked_files[:2])}")
    
    if not summary:
        return "Minor code cleanup"
        
    message = " | ".join(summary)
    return message[:72] # Ограничение длины для Git

def sync_and_push():
    repo = find_repo_root()
    
    if not repo:
        print("❌ Ошибка: Не удалось найти Git-репозиторий в этой папке или выше.")
        return

    print(f"📂 Работаем в репозитории: {repo.working_tree_dir}")

    try:
        # 1. Добавляем всё в индекс
        print("🔍 Индексируем изменения...")
        repo.git.add(all=True)

        # 2. Pull (чтобы не было конфликтов с коллегами)
        print("📥 Подтягиваем изменения коллег (Pull)...")
        repo.remotes.origin.pull()

        # 3. Проверяем, нужно ли делать коммит
        if not repo.is_dirty(untracked_files=True):
            print("✨ Новых изменений для коммита нет.")
            return

        # 4. Формируем сообщение и коммитим
        commit_msg = f"Auto: {generate_commit_message(repo)}"
        print(f"📝 Коммит: {commit_msg}")
        repo.index.commit(commit_msg)

        # 5. Пушим
        print("🚀 Пушим в репозиторий...")
        repo.remotes.origin.push()
        print("✅ Все изменения успешно отправлены!")

    except git.exc.GitCommandError as e:
        print(f"❌ Ошибка Git: {e}")
    except Exception as e:
        print(f"❌ Системная ошибка: {e}")

if __name__ == "__main__":
    sync_and_push()
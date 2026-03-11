import os
import re
import yaml
from dotenv import load_dotenv
from pathlib import Path

class YAMLEnvParser:
    def __init__(self, yaml_path='config/db_config.yaml', env_path='config/.env'):
        self.yaml_path = Path(yaml_path)
        self.env_path = Path(env_path)
        self._load_env()
        self.raw_config = self._load_yaml()
        self.parsed_config = self._parse_variables(self.raw_config)
    
    def _load_env(self):
        load_dotenv(self.env_path)

    def _load_yaml(self):
        with open(self.yaml_path, 'r', encoding='utf-8') as f:
            config = yaml.safe_load(f)
        return config
    
    def _parse_variables(self, data):
        if isinstance(data, dict):
            return {key: self._parse_variables(value) for key, value in data.items()}
        elif isinstance(data, list):
            return [self._parse_variables(item) for item in data]
        elif isinstance(data, str):
            pattern = r'\${([^}]+)}'
            
            def replace_var(match):
                var_name = match.group(1)
                var_value = os.getenv(var_name)
                if var_value is None:
                    return match.group(0)
                return var_value
            
            return re.sub(pattern, replace_var, data)
        else:
            return data
    
    def get_config(self):
        return self.parsed_config
    
    def get_db_config(self):
        return self.parsed_config
    
    def get_connection_params(self):
        config = self.parsed_config
        return {
            'host': config.get('host'),
            'port': config.get('port'),
            'database': config.get('name'),
            'user': config.get('user'),
            'password': config.get('password')
        }
    
    def get_connection_string(self):
        config = self.parsed_config
        return f"postgresql://{config['user']}:{config['password']}@{config['host']}:{config['port']}/{config['name']}"
    
    def __str__(self):
        config = self.parsed_config
        safe_config = config.copy()
        if 'password' in safe_config:
            safe_config['password'] = '$#$#$#$#'
        return f'DatabaseConfig({safe_config})'

if __name__ == '__main__':
    parser = YAMLEnvParser()
    config = parser.get_config()
    for key, value in config.items():
        if key != 'password':
            print(f'  {key}: {value}')
        else:
            print(f'  {key}: {'$#' * len(value)}')
    params = parser.get_connection_params()
    print(f'\n Параметры подключения: {params}')
    conn_string = parser.get_connection_string()
    print(f'\n Строка подключения: {conn_string}')
    import psycopg2
    connection = psycopg2.connect(**params)
    if connection:
        cursor = connection.cursor()
        cursor.execute('SELECT version();')
        version = cursor.fetchone()
        cursor.close()
        connection.close()
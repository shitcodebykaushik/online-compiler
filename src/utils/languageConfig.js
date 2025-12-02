export const languageConfig = {
  c: {
    id: 50,
    name: 'C (GCC 9.2.0)',
    label: 'C',
    icon: '🔷',
    monacoLang: 'c',
    extension: '.c'
  },
  cpp: {
    id: 54,
    name: 'C++ (GCC 9.2.0)',
    label: 'C++',
    icon: '🔶',
    monacoLang: 'cpp',
    extension: '.cpp'
  },
  python: {
    id: 71,
    name: 'Python (3.8.1)',
    label: 'Python',
    icon: '🐍',
    monacoLang: 'python',
    extension: '.py'
  },
  java: {
    id: 62,
    name: 'Java (OpenJDK 13.0.1)',
    label: 'Java',
    icon: '☕',
    monacoLang: 'java',
    extension: '.java'
  },
  javascript: {
    id: 63,
    name: 'JavaScript (Node.js 12.14.0)',
    label: 'JavaScript',
    icon: '🟨',
    monacoLang: 'javascript',
    extension: '.js'
  },
  go: {
    id: 60,
    name: 'Go (1.13.5)',
    label: 'Go',
    icon: '🔵',
    monacoLang: 'go',
    extension: '.go'
  },
  rust: {
    id: 73,
    name: 'Rust (1.40.0)',
    label: 'Rust',
    icon: '🦀',
    monacoLang: 'rust',
    extension: '.rs'
  },
  php: {
    id: 68,
    name: 'PHP (7.4.1)',
    label: 'PHP',
    icon: '🐘',
    monacoLang: 'php',
    extension: '.php'
  }
};

export const getLanguageById = (langKey) => {
  return languageConfig[langKey] || languageConfig.python;
};

export const getAllLanguages = () => {
  return Object.entries(languageConfig).map(([key, value]) => ({
    key,
    ...value
  }));
};

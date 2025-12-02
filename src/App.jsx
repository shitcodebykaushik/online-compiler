import { useState, useEffect } from 'react';
import Header from './components/Header';
import CodeEditor from './components/CodeEditor';
import Console from './components/Console';
import LanguageSelector from './components/LanguageSelector';
import ControlPanel from './components/ControlPanel';
import useCodeExecution from './hooks/useCodeExecution';
import { getLanguageById } from './utils/languageConfig';
import { getTemplate } from './utils/codeTemplates';
import './App.css';

function App() {
  const [theme, setTheme] = useState('dark');
  const [language, setLanguage] = useState('python');
  const [code, setCode] = useState('');

  const {
    output,
    error,
    loading,
    executionTime,
    memory,
    executeCode,
    clearOutput
  } = useCodeExecution();

  // Initialize code with template
  useEffect(() => {
    setCode(getTemplate(language));
  }, []);

  // Update code when language changes
  const handleLanguageChange = (newLang) => {
    setLanguage(newLang);
    setCode(getTemplate(newLang));
    clearOutput();
  };

  const handleThemeToggle = () => {
    const newTheme = theme === 'dark' ? 'light' : 'dark';
    setTheme(newTheme);
    document.documentElement.setAttribute('data-theme', newTheme);
  };

  const handleRun = () => {
    const langConfig = getLanguageById(language);
    executeCode(code, langConfig.id);
  };

  const handleReset = () => {
    setCode(getTemplate(language));
    clearOutput();
  };

  const handleShare = () => {
    // TODO: Implement share functionality
    alert('Share functionality will be implemented with backend!');
  };

  const handleSave = () => {
    // TODO: Implement save functionality
    alert('Save functionality will be implemented with backend!');
  };

  const langConfig = getLanguageById(language);

  return (
    <div className="app" data-theme={theme}>
      <Header theme={theme} onThemeToggle={handleThemeToggle} />

      <main className="main-container">
        <div className="content-wrapper">
          <div className="top-bar">
            <LanguageSelector
              selected={language}
              onSelect={handleLanguageChange}
            />

            <ControlPanel
              onRun={handleRun}
              onReset={handleReset}
              onShare={handleShare}
              onSave={handleSave}
              isRunning={loading}
            />
          </div>

          <div className="editor-console-grid">
            <div className="editor-section">
              <CodeEditor
                code={code}
                onChange={setCode}
                language={langConfig.monacoLang}
                theme={theme}
              />
            </div>

            <div className="console-section">
              <Console
                output={output}
                error={error}
                loading={loading}
                executionTime={executionTime}
                memory={memory}
                onClear={clearOutput}
              />
            </div>
          </div>
        </div>
      </main>

      <footer className="footer">
        <p>
          Built with React, Monaco Editor, Go, Judge0, Docker, Redis & SQLite
        </p>
        <p className="footer-note">
          Frontend ready • Backend integration pending
        </p>
      </footer>
    </div>
  );
}

export default App;

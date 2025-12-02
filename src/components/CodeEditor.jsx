import React, { useState } from 'react';
import Editor from '@monaco-editor/react';
import { FileCode } from 'lucide-react';
import './CodeEditor.css';

const CodeEditor = ({ code, onChange, language, theme }) => {
    const [isEditorReady, setIsEditorReady] = useState(false);

    const handleEditorDidMount = (editor, monaco) => {
        setIsEditorReady(true);

        // Configure editor options
        editor.updateOptions({
            fontSize: 14,
            fontFamily: 'Fira Code, Consolas, Monaco, monospace',
            lineHeight: 21,
            minimap: { enabled: false },
            scrollBeyondLastLine: false,
            smoothScrolling: true,
            cursorBlinking: 'smooth',
            cursorSmoothCaretAnimation: true,
            padding: { top: 16, bottom: 16 },
            renderLineHighlight: 'all',
            roundedSelection: true,
            wordWrap: 'on',
        });
    };

    return (
        <div className="code-editor-container">
            <div className="editor-header">
                <div className="editor-title">
                    <FileCode size={18} />
                    <span>Code Editor</span>
                </div>
                <div className="editor-status">
                    {isEditorReady && <span className="status-dot"></span>}
                    <span className="status-text">
                        {isEditorReady ? 'Ready' : 'Loading...'}
                    </span>
                </div>
            </div>

            <div className="editor-wrapper">
                <Editor
                    height="100%"
                    language={language}
                    value={code}
                    onChange={onChange}
                    theme={theme === 'dark' ? 'vs-dark' : 'light'}
                    onMount={handleEditorDidMount}
                    options={{
                        selectOnLineNumbers: true,
                        automaticLayout: true,
                    }}
                    loading={
                        <div className="editor-loading">
                            <div className="loading-spinner"></div>
                            <p>Loading editor...</p>
                        </div>
                    }
                />
            </div>
        </div>
    );
};

export default CodeEditor;

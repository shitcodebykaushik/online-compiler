import React from 'react';
import { Terminal, Trash2, Clock, MemoryStick } from 'lucide-react';
import './Console.css';

const Console = ({ output, error, loading, executionTime, memory, onClear }) => {
    const hasOutput = output || error;

    return (
        <div className="console-container">
            <div className="console-header">
                <div className="console-title">
                    <Terminal size={18} />
                    <span>Output</span>
                </div>

                <div className="console-actions">
                    {executionTime && (
                        <div className="stat">
                            <Clock size={14} />
                            <span>{executionTime}ms</span>
                        </div>
                    )}
                    {memory && (
                        <div className="stat">
                            <MemoryStick size={14} />
                            <span>{memory}KB</span>
                        </div>
                    )}
                    <button
                        className="btn-icon"
                        onClick={onClear}
                        aria-label="Clear output"
                    >
                        <Trash2 size={16} />
                    </button>
                </div>
            </div>

            <div className="console-content">
                {loading ? (
                    <div className="console-loading">
                        <div className="loading-dots">
                            <span></span>
                            <span></span>
                            <span></span>
                        </div>
                        <p>Executing code...</p>
                    </div>
                ) : hasOutput ? (
                    <div className="console-output">
                        {error ? (
                            <pre className="error-text">{error}</pre>
                        ) : (
                            <pre className="output-text">{output}</pre>
                        )}
                    </div>
                ) : (
                    <div className="console-empty">
                        <Terminal size={48} strokeWidth={1} />
                        <p>Run your code to see the output here</p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default Console;

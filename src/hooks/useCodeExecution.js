import { useState } from 'react';

// Get API URL from environment variable or use default
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// Custom hook for code execution
// This will be connected to the backend API later
const useCodeExecution = () => {
    const [output, setOutput] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const [executionTime, setExecutionTime] = useState(null);
    const [memory, setMemory] = useState(null);

    const executeCode = async (code, languageId) => {
        setLoading(true);
        setOutput('');
        setError('');
        setExecutionTime(null);
        setMemory(null);

        try {
            const response = await fetch(`${API_URL}/api/v1/execute`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    language_id: languageId,
                    code: code,
                    stdin: ''
                })
            });

            const result = await response.json();

            if (!result.success) {
                setError(result.error || 'Execution failed');
            } else {
                if (result.output) {
                    setOutput(result.output);
                }
                if (result.error) {
                    setError(result.error);
                }
                if (result.execution_time) {
                    setExecutionTime(result.execution_time.toFixed(2));
                }
                if (result.memory_kb) {
                    setMemory(result.memory_kb);
                }
            }
        } catch (err) {
            setError(`Network Error: ${err.message}\n\nMake sure the backend server is running on http://localhost:8080`);
        } finally {
            setLoading(false);
        }
    };

    const clearOutput = () => {
        setOutput('');
        setError('');
        setExecutionTime(null);
        setMemory(null);
    };

    return {
        output,
        error,
        loading,
        executionTime,
        memory,
        executeCode,
        clearOutput
    };
};

export default useCodeExecution;

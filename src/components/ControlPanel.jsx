import React from 'react';
import { Play, RotateCcw, Share2, Save } from 'lucide-react';
import './ControlPanel.css';

const ControlPanel = ({
    onRun,
    onReset,
    onShare,
    onSave,
    isRunning,
    disabled
}) => {
    return (
        <div className="control-panel">
            <button
                className="btn btn-primary btn-run"
                onClick={onRun}
                disabled={isRunning || disabled}
            >
                <Play size={18} fill="currentColor" />
                {isRunning ? 'Running...' : 'Run Code'}
            </button>

            <div className="control-divider"></div>

            <button
                className="btn btn-secondary"
                onClick={onReset}
                disabled={isRunning}
                title="Reset to template"
            >
                <RotateCcw size={18} />
                Reset
            </button>

            <button
                className="btn btn-secondary"
                onClick={onShare}
                disabled={isRunning}
                title="Share code"
            >
                <Share2 size={18} />
                Share
            </button>

            <button
                className="btn btn-secondary"
                onClick={onSave}
                disabled={isRunning}
                title="Save code"
            >
                <Save size={18} />
                Save
            </button>
        </div>
    );
};

export default ControlPanel;

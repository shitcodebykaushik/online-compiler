import React from 'react';
import { Code2, Moon, Sun, Github } from 'lucide-react';
import './Header.css';

const Header = ({ theme, onThemeToggle }) => {
    return (
        <header className="header">
            <div className="header-container">
                <div className="header-left">
                    <div className="logo">
                        <Code2 size={28} strokeWidth={2.5} />
                        <span className="logo-text gradient-text">CodeRunner</span>
                    </div>
                </div>

                <nav className="header-nav">
                    <a href="#compiler" className="nav-link">Compiler</a>
                    <a href="#about" className="nav-link">About</a>
                    <a href="#docs" className="nav-link">Docs</a>
                </nav>

                <div className="header-right">
                    <button
                        className="btn-icon theme-toggle"
                        onClick={onThemeToggle}
                        aria-label="Toggle theme"
                    >
                        {theme === 'dark' ? <Sun size={20} /> : <Moon size={20} />}
                    </button>

                    <a
                        href="https://github.com"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="btn-icon"
                        aria-label="GitHub"
                    >
                        <Github size={20} />
                    </a>
                </div>
            </div>
        </header>
    );
};

export default Header;

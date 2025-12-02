import React, { useState, useRef, useEffect } from 'react';
import { ChevronDown, Check } from 'lucide-react';
import { getAllLanguages } from '../utils/languageConfig';
import './LanguageSelector.css';

const LanguageSelector = ({ selected, onSelect }) => {
    const [isOpen, setIsOpen] = useState(false);
    const dropdownRef = useRef(null);
    const languages = getAllLanguages();

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
                setIsOpen(false);
            }
        };

        document.addEventListener('mousedown', handleClickOutside);
        return () => document.removeEventListener('mousedown', handleClickOutside);
    }, []);

    const selectedLang = languages.find(lang => lang.key === selected) || languages[0];

    const handleSelect = (langKey) => {
        onSelect(langKey);
        setIsOpen(false);
    };

    return (
        <div className="language-selector" ref={dropdownRef}>
            <button
                className="selector-button"
                onClick={() => setIsOpen(!isOpen)}
                aria-expanded={isOpen}
            >
                <span className="lang-icon">{selectedLang.icon}</span>
                <span className="lang-name">{selectedLang.label}</span>
                <ChevronDown
                    size={16}
                    className={`chevron ${isOpen ? 'open' : ''}`}
                />
            </button>

            {isOpen && (
                <div className="dropdown-menu">
                    <div className="dropdown-header">
                        Select Language
                    </div>
                    <div className="dropdown-list">
                        {languages.map((lang) => (
                            <button
                                key={lang.key}
                                className={`dropdown-item ${selected === lang.key ? 'active' : ''}`}
                                onClick={() => handleSelect(lang.key)}
                            >
                                <span className="lang-icon">{lang.icon}</span>
                                <div className="lang-info">
                                    <span className="lang-label">{lang.label}</span>
                                    <span className="lang-version">{lang.name}</span>
                                </div>
                                {selected === lang.key && (
                                    <Check size={16} className="check-icon" />
                                )}
                            </button>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default LanguageSelector;

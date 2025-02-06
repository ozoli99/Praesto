import React from "react";

const Button = ({ children, onClick, className = "" }) => (
    <button
        onClick={onClick}
        className={`py-2 px-4 bg-blue-500 text-white rounded ${className}`}>
        {children}
    </button>
);

export default Button;
